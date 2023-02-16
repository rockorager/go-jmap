package jmap

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

// A JMAP Client
type Client struct {
	sync.Mutex
	// The HttpClient.Client to use for requests. The HttpClient.Client should handle
	// authentication. Calling WithBasicAuth or WithAccessToken on the
	// Client will set the HttpClient to one which uses authentication
	HttpClient *http.Client

	// The JMAP Session Resource Endpoint. If the client detects the Session
	// object needs refetching, it will automatically do so.
	SessionEndpoint string

	// the JMAP Session object
	Session *Session
}

// Set the HttpClient to a client which authenticates using the provided
// username and password
func (c *Client) WithBasicAuth(username string, password string) *Client {
	ctx := context.Background()
	auth := username + ":" + password
	t := &oauth2.Token{
		AccessToken: base64.StdEncoding.EncodeToString([]byte(auth)),
		TokenType:   "basic",
	}
	cfg := &oauth2.Config{}
	c.HttpClient = oauth2.NewClient(ctx, cfg.TokenSource(ctx, t))
	return c
}

// Set the HttpClient to a client which authenticates using the provided Access
// Token
func (c *Client) WithAccessToken(token string) *Client {
	ctx := context.Background()
	t := &oauth2.Token{
		AccessToken: token,
		TokenType:   "bearer",
	}
	cfg := &oauth2.Config{}
	c.HttpClient = oauth2.NewClient(ctx, cfg.TokenSource(ctx, t))
	return c
}

// Authenticate authenticates the client and retrieves the Session object.
// Authenticate will be called automatically when Do is called if the Session
// object hasn't already been initialized. Call Authenticate before any requests
// if you need to access information from the Session object prior to the first
// request
func (c *Client) Authenticate() error {
	c.Lock()
	if c.SessionEndpoint == "" {
		c.Unlock()
		return fmt.Errorf("no session url is set")
	}
	c.Unlock()

	req, err := http.NewRequest("GET", c.SessionEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("couldn't authenticate")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	s := &Session{}
	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}
	c.Session = s
	return nil
}

// Do performs a JMAP request and returns the response
func (c *Client) Do(req *Request) (*Response, error) {
	c.Lock()
	if c.Session == nil {
		c.Unlock()
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	c.Unlock()
	// Check the required capabilities before making the request
	for _, uri := range req.Using {
		c.Lock()
		_, ok := c.Session.Capabilities[uri]
		c.Unlock()
		if !ok {
			return nil, fmt.Errorf("server doesn't support required capability '%s'", uri)
		}
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if req.Context == nil {
		req.Context = context.Background()
	}
	httpReq, err := http.NewRequestWithContext(req.Context, "POST", c.Session.APIURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		return nil, decodeHttpError(httpResp)
	}

	data, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, fmt.Errorf("error? %v", err)
	}

	return resp, nil
}

// Upload sends binary data to the server and returns blob ID and some
// associated meta-data.
//
// There are some caveats to keep in mind:
// - Server may return the same blob ID for multiple uploads of the same blob.
// - Blob ID may become invalid after some time if it is unused.
// - Blob ID is usable only by the uploader until it is used, even for shared accounts.
func (c *Client) Upload(accountID ID, blob io.Reader) (*UploadResponse, error) {
	c.Lock()
	if c.SessionEndpoint == "" {
		c.Unlock()
		return nil, fmt.Errorf("jmap/client: SessionEndpoint is empty")
	}
	if c.Session == nil {
		c.Unlock()
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}

	url := strings.ReplaceAll(c.Session.UploadURL, "{accountId}", string(accountID))
	c.Unlock()
	req, err := http.NewRequest("POST", url, blob)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, decodeHttpError(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	info := &UploadResponse{}
	err = json.Unmarshal(data, info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Download downloads binary data by its Blob ID from the server.
func (c *Client) Download(accountID ID, blobID ID) (io.ReadCloser, error) {
	c.Lock()
	if c.SessionEndpoint == "" {
		c.Unlock()
		return nil, fmt.Errorf("jmap/client: SessionEndpoint is empty")
	}
	if c.Session == nil {
		c.Unlock()
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}

	urlRepl := strings.NewReplacer(
		"{accountId}", string(accountID),
		"{blobId}", string(blobID),
		"{type}", "application/octet-stream",
		"{name}", "filename",
	)
	tgtUrl := urlRepl.Replace(c.Session.DownloadURL)
	c.Unlock()
	req, err := http.NewRequest("GET", tgtUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		return nil, decodeHttpError(resp)
	}

	return resp.Body, nil
}

func decodeHttpError(resp *http.Response) error {
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		return fmt.Errorf("HTTP %d %s", resp.StatusCode, resp.Status)
	}

	reqErr := &RequestError{}
	if err := json.NewDecoder(resp.Body).Decode(reqErr); err != nil {
		return fmt.Errorf("HTTP %d %s (failed to decode JSON body: %v)", resp.StatusCode, resp.Status, err)
	}

	return reqErr
}

// UploadResponse is the object returned in response to blob upload.
type UploadResponse struct {
	// The id of the account used for the call.
	Account ID `json:"accountId"`

	// The id representing the binary data uploaded. The data for this id is
	// immutable. The id only refers to the binary data, not any metadata.
	ID ID `json:"blobId"`

	// The media type of the file (as specified in RFC 6838, section 4.2) as
	// set in the Content-Type header of the upload HTTP request.
	Type string `json:"type"`

	// The size of the file in octets.
	Size uint64 `json:"size"`
}

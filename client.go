package jmap

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

// A JMAP Client
type Client struct {
	// The HttpClient.Client to use for requests. The HttpClient.Client should handle
	// authentication.
	HttpClient *http.Client

	// The JMAP Session Resource Endpoint
	Endpoint string

	// the JMAP session object
	session *Session
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
func (c *Client) Authenticate() (*Session, error) {
	if c.Endpoint == "" {
		return nil, fmt.Errorf("no session url is set")
	}

	req, err := http.NewRequest("GET", c.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("couldn't authenticate")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s := &Session{}
	err = json.Unmarshal(data, s)
	if err != nil {
		return nil, err
	}
	c.session = s
	return s, nil
}

// Do performs a JMAP request and returns the response
func (c *Client) Do(req *Request) (*Response, error) {
	if c.session == nil {
		_, err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", c.session.APIURL, bytes.NewReader(body))
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
		return nil, fmt.Errorf("error: %d", httpResp.StatusCode)
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

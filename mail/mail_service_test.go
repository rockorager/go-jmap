package mail

import (
	"io"
	"net/http"
	"os"
	"testing"

	"git.sr.ht/~rockorager/go-jmap/client"
)

type TestHTTPClient struct {
	// resp is the json response
	resp       io.ReadCloser
	statusCode int
	req        *http.Request
}

func (t *TestHTTPClient) Do(req *http.Request) (*http.Response, error) {
	t.req = req
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	resp := &http.Response{
		StatusCode: t.statusCode,
		Body:       t.resp,
		Header:     header,
	}
	return resp, nil
}

func newTestClientWithSession(t *testing.T) client.HTTPClient {
	return &TestHTTPClient{
		statusCode: 200,
		resp:       openFile(t, "./test-data/session-response.json"),
	}
}

func openFile(t *testing.T, path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("error opening filepath %s: %v", path, err)
	}
	return file
}

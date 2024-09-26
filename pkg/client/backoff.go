package client

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cenkalti/backoff/v4"
)

// NewBackoffHTTPClient returns a new instance of an http client with a retry policy
func NewBackoffHTTPClient(policy backoff.BackOff, opts ...Option) *BackoffHTTPClient {
	innerClient := NewHTTPClient(opts...)
	return &BackoffHTTPClient{innerClient, policy}
}

// BackoffHTTPClient represents a client that can retry its operations following a specified policy
type BackoffHTTPClient struct {
	client Client
	policy backoff.BackOff
}

// Do runs client's Do operation wrapped in a retry
func (c *BackoffHTTPClient) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	// get the body of the request to copy it in all retries
	var body []byte
	if req.Body != nil {
		defer req.Body.Close()
		body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}
	err = backoff.Retry(func() error {
		var e error
		// cloning the initial request because the body is consumed at each attempt
		request := req.Clone(req.Context())
		request.Body = io.NopCloser(bytes.NewBuffer(body))
		//nolint:bodyclose
		resp, e = c.client.Do(request)
		return e
	}, c.policy)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DoJSON runs the client's DoJSON operation wrapped in a retry
func (c *BackoffHTTPClient) DoJSON(req *http.Request, respObj interface{}) error {
	return backoff.Retry(func() error {
		return c.client.DoJSON(req, respObj)
	}, c.policy)
}

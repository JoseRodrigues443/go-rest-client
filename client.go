// Package client that provides HTTP rest client capabilities.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// HTTPClient interface for http requests.
type HTTPClient interface {
	buildRequest(method Method, path string, body interface{}) (*http.Request, error)
	call(req *http.Request, v interface{}) (*http.Response, error)
}

// Client stores the needed information for a client.
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

// NewClient creates a base instance of a basic client.
func NewClient(baseURL string, userAgent string) *Client {
	base, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		BaseURL:   base,
		UserAgent: userAgent,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// Method is the HTTP verb to use.
type Method string

// buildRequest creates a request.
func (c *Client) buildRequest(method Method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, string(method), u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

// request makes the final http request and returns a http response.
func (c *Client) request(req *http.Request, value interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer dclose(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(value)
	return resp, err
}

// Call makes the final http request and returns a http response.
func (c *Client) Call(method Method, path string, body interface{}, value interface{}) (*http.Response, error) {
	req, err := c.buildRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return c.request(req, value)
}

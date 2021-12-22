// Package client that provides HTTP rest client capabilities.
package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

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

const (
	// Get verb.
	Get Method = "GET"
	// Head verb.
	Head Method = "HEAD"
	// Post verb.
	Post Method = "POST"
	// Put verb.
	Put Method = "PUT"
	// Patch verb.
	Patch Method = "PATCH" // RFC 5789
	// Delete verb.
	Delete Method = "DELETE"
	// Connect verb.
	Connect Method = "CONNECT"
	// Options verb.
	Options Method = "OPTIONS"
	// Trace verb.
	Trace Method = "TRACE"
)

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
	req, err := http.NewRequest(string(method), u.String(), buf)
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

// call makes the final http request and returns a http response.
func (c *Client) call(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

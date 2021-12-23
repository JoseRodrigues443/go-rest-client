// Package client that provides HTTP rest client capabilities.
package client

import (
	"io"
	"log"
	"time"
)

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

const duration time.Duration = 80

// dclose handles close on a safe manner.
func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Address   = "https://www.test.com"
	UserAgent = "Test/1.0"
)

func Test_Main(t *testing.T) {
	main() // for golint to ignore dead code warning
}

func Test_NewClient(t *testing.T) {
	client := NewClient(Address, UserAgent)

	assert.Equal(t, Address, client.BaseURL.String(), "Should have the same base URL")
	assert.Equal(t, UserAgent, client.UserAgent, "Should have the same user agent")
}

func Test_buildRequest(t *testing.T) {
	client := NewClient(Address, UserAgent)

	req, err := client.buildRequest(Get, "test", nil)

	assert.Nil(t, err, "Should not give error")
	assert.NotNil(t, req, "Response should exist")
}

type Message struct {
	Name, Text string
}

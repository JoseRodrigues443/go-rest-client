package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Address   = "https://www.test.com"
	UserAgent = "Test/1.0"
)

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

func Test_Call(t *testing.T) {
	client := NewClient(Address, UserAgent)

	ctx, _ := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", Address, nil)

	var m Message
	res, err := client.call(req, &m)

	assert.Nil(t, err, "Should not give error")
	assert.NotNil(t, res, "Response should exist")
}

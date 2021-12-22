package client

import (
	"net/http"
)

type Client struct {
	cl *http.Client
}

package client

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_NewClient(t *testing.T) {
	faker := faker.New()
	address := faker.Person().Name()
	userAgent := faker.Internet().URL()

	client := NewClient(address, userAgent)

	assert.Equal(t, address, client.BaseURL.Path, "Should have the same base URL")
	assert.Equal(t, userAgent, client.UserAgent, "Should have the same user agent")
}

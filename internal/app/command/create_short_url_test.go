package command_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlValid(t *testing.T) {

	URI := "https://google.com"

	u, err := url.ParseRequestURI(URI)

	assert.NoError(t, err)

	assert.NotNil(t, u)

	assert.Equal(t, "https", u.Scheme)
}

func TestUrlNoValid(t *testing.T) {
	URI := "google.com"

	_, err := url.ParseRequestURI(URI)

	assert.Error(t, err)
}

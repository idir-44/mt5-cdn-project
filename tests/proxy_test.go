package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxyHealth(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

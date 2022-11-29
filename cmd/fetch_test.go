package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestFetch tests making a request to https://rest.coinapi.io/v1/exchangerate
// As no API key is provided the request should fail with 401 unauthorized.
func TestFetch(t *testing.T) {
	val := fetch(&cobra.Command{}, []string{})
	assert.Equal(
		t,
		val,
		fmt.Errorf("request to %s/%s/%s failed with status code %d", coinApiUrl, crypto, in, 401),
	)
}

// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package breakingbad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	// Initialize the client using the environment variables
	options := ClientOptions{
		ApiVersion: "v1",
	}

	client, err := NewClient(options)

	// Assertions
	assert.NoError(t, err, "Expected no error when initializing client")
	assert.NotNil(t, client, "Client should not be nil")
	assert.Equal(t, "https://api.breakingbadquotes.xyz", client.ApiUrl, "Client API URL should match the environment variable")
	assert.Equal(t, "v1", client.ApiVersion, "Client API version should match the environment variable")
}

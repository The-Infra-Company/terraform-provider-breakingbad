// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package breakingbad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGettingQuote(t *testing.T) {
	// Initialize the client
	options := ClientOptions{
		ApiVersion: "v1",
	}
	fmt.Printf("Initializing client with API URL: %s\n", options.ApiUrl)

	client, err := NewClient(options)
	if err != nil {
		t.Fatalf("Error creating client: %s", err)
	}

	// Attempt to get quote
	response, err := client.GetQuote()
	if err != nil {
		t.Errorf("Error getting quote: %s", err)
	} else {
		// Log the response
		responseJSON, err := json.Marshal(response)
		if err != nil {
			t.Errorf("Error marshaling server response: %s", err)
		} else {
			// Attempt to pretty-print the JSON
			var prettyResponse bytes.Buffer
			if err := json.Indent(&prettyResponse, responseJSON, "", "  "); err == nil {
				fmt.Printf("\n=== Parsed Response Body ===\n%s\n", prettyResponse.String())
			} else {
				fmt.Println("Failed to pretty-print JSON.")
			}
		}
	}
}

// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package breakingbad

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetQuote() (QuoteApiResponse, error) {
	// Create a new request
	req, err := c.newRequest("GET", fmt.Sprintf("/%s/quotes/", c.ApiVersion), nil)
	if err != nil {
		return QuoteApiResponse{}, fmt.Errorf("failed to create request for GetQuote: %w", err)
	}

	// Perform the request
	res, err := c.doRequest(req)
	if err != nil {
		return QuoteApiResponse{}, fmt.Errorf("failed to fetch quote: %w", err)
	}

	// Parse the response (expecting an array of quotes)
	var quotes []QuoteApiResponse
	if err := json.Unmarshal(res, &quotes); err != nil {
		return QuoteApiResponse{}, fmt.Errorf("failed to parse quotes response: %w", err)
	}

	// Check if the array is empty
	if len(quotes) == 0 {
		return QuoteApiResponse{}, fmt.Errorf("no quotes found in the response")
	}

	// Return the first quote
	return quotes[0], nil
}

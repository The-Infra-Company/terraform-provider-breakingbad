// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package breakingbad

import (
	"fmt"
	"io"
	"net/http"
)

const (
	ApiEndpoint = "https://api.breakingbadquotes.xyz"
	ApiVersion  = "v1"
)

type ClientOptions struct {
	ApiUrl     string
	ApiVersion string
}

// Client definition.
type Client struct {
	ApiUrl     string
	ApiVersion string

	HTTPClient *http.Client
}

// NewClient returns a new client and authenticates.
func NewClient(options ClientOptions) (*Client, error) {
	// Default to the constant API endpoint if not provided
	apiUrl := ApiEndpoint
	if options.ApiUrl != "" {
		apiUrl = options.ApiUrl
	}

	// Default to the constant API version if not provided
	apiVersion := ApiVersion
	if options.ApiVersion != "" {
		apiVersion = options.ApiVersion
	}

	client := &Client{
		ApiUrl:     apiUrl,
		ApiVersion: apiVersion,
		HTTPClient: http.DefaultClient,
	}

	return client, nil
}

// NewRequest creates a new HTTP request
func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.ApiUrl+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// doRequest sends an HTTP request and handles the response.
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return body, nil
	}

	return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
}

// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package breakingbad

// Struct to match the entire API response.
type QuoteApiResponse struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

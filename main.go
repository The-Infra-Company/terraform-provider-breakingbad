// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-breakingbad/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Dynamically injected via Goreleaser.
var version string = "0.1.0"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/The-Infra-Company/breakingbad",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/The-Infra-Company/terraform-provider-breakingbad/internal/breakingbad"
)

// Ensure the provider satisfies the provider.Provider interface.
var _ provider.Provider = &BreakingBadProvider{}

// BreakingBadProvider defines the provider implementation.
type BreakingBadProvider struct {
	// version is set to the provider version on release.
	version string
}

// BreakingBadProviderModel describes the provider data model.
type BreakingBadProviderModel struct {
	URL types.String `tfsdk:"api_url"`
}

func (p *BreakingBadProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "breakingbad"
	resp.Version = p.version
}

func (p *BreakingBadProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_url": schema.StringAttribute{
				Optional:    true,
				Description: "The URL endpoint for the Breaking Bad Quotes API. Defaults to 'https://api.breakingbadquotes.xyz'.",
			},
		},
	}
}

func (p *BreakingBadProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config BreakingBadProviderModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set default endpoint if not provided
	url := "https://api.breakingbadquotes.xyz"
	if !config.URL.IsNull() {
		url = config.URL.ValueString()
	}

	// Initialize the Breaking Bad API client
	client, err := breakingbad.NewClient(breakingbad.ClientOptions{
		ApiUrl: url,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Breaking Bad API client",
			fmt.Sprintf("Error: %s", err),
		)
		return
	}

	// Log the client initialization
	tflog.Info(ctx, "Configured Breaking Bad API client", map[string]interface{}{
		"api_url": url,
	})

	// Make the client available during data source and resource operations
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *BreakingBadProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *BreakingBadProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		QuotesDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BreakingBadProvider{
			version: version,
		}
	}
}

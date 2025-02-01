// Copyright (c) The Infra Company
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/The-Infra-Company/terraform-provider-breakingbad/internal/breakingbad"
)

// Ensure the data source implements the DataSource interface.
var _ datasource.DataSource = &QuoteDataSource{}

// QuoteDataSource defines the data source implementation.
type QuoteDataSource struct {
	client *breakingbad.Client
}

// QuoteDataSourceModel describes the data source data model.
type QuoteDataSourceModel struct {
	Quote  types.String `tfsdk:"quote"`
	Author types.String `tfsdk:"author"`
}

func QuotesDataSource() datasource.DataSource {
	return &QuoteDataSource{}
}

func (d *QuoteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quote"
}

func (d *QuoteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"quote": schema.StringAttribute{
				Computed:    true,
				Description: "The quote from the Breaking Bad API.",
			},
			"author": schema.StringAttribute{
				Computed:    true,
				Description: "The author of the quote.",
			},
		},
	}
}

func (d *QuoteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*breakingbad.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *breakingbad.Client, got: %T", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *QuoteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QuoteDataSourceModel

	// Fetch the quote from the API
	quote, err := d.client.GetQuote()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch quote",
			fmt.Sprintf("Error: %s", err),
		)
		return
	}

	// Map the API response to the Terraform data model
	data.Quote = types.StringValue(quote.Quote)
	data.Author = types.StringValue(quote.Author)

	// Log the data
	tflog.Info(ctx, "Fetched quote from Breaking Bad API", map[string]interface{}{
		"quote":  quote.Quote,
		"author": quote.Author,
	})

	// Save the data into the Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

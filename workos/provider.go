package workos

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/workos/workos-go/pkg/organizations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &workosProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &workosProvider{}
}

type workosClient struct {
	Organizations *organizations.Client
}

// workosProvider is the provider implementation.
type workosProvider struct{}

// workosProviderModel maps provider schema data to a Go type.
type workosProviderModel struct {
	Host   types.String `tfsdk:"host"`
	ApiKey types.String `tfsdk:"api_key"`
}

// Metadata returns the provider type name.
func (p *workosProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "workos"
}

// Schema defines the provider-level schema for configuration data.
func (p *workosProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a WorkOS API client for data sources and resources.
func (p *workosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring WorkOS client")
	var config workosProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown WorkOS API Key",
			"The provider cannot create the WorkOS API client as there is an unknown configuration value for the WorkOS API key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the WORKOS_API_KEY environment variable.")
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown WorkOS API Host",
			"The provider cannot create the WorkOS API client as there is an unknown configuration value for the WorkOS API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the WORKOS_API_HOST environment variable.")
	}

	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("WORKOS_API_HOST")
	apiKey := os.Getenv("WORKOS_API_KEY")

	if !config.ApiKey.IsNull() {
		apiKey = config.ApiKey.ValueString()
	}

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing WorkOS API Secret Key",
			"The provider cannot create the WorkOS API client as there is a missing or empty value for the WorkOS API key. "+
				"Set the password value in the configuration or use the WORKOS_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "workos_host", host)
	ctx = tflog.SetField(ctx, "workos_api_key", apiKey)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "workos_api_key")
	organizations.SetAPIKey(apiKey)
	if host != "" {
		organizations.DefaultClient.Endpoint = host
	}

	client := &workosClient{
		Organizations: organizations.DefaultClient,
	}
	resp.DataSourceData = client
	resp.ResourceData = client
	tflog.Info(ctx, "Configured WorkOS client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *workosProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOrganizationsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *workosProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrganizationResource,
	}
}

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &ContainerregistryProvider{}

// ContainerregistryProvider defines the provider implementation.
type ContainerregistryProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ContainerregistryProviderModel describes the provider data model.
type ContainerregistryProviderModel struct {
}

func (p *ContainerregistryProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "containerregistry"
	resp.Version = p.version
}

func (p *ContainerregistryProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *ContainerregistryProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ContainerregistryProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *ContainerregistryProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewContainerregistryResource,
	}
}

func (p *ContainerregistryProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ContainerregistryProvider{
			version: version,
		}
	}
}

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.Provider = (*structureDeployProvider)(nil)

// var _ provider. = (*structureDeployProvider)(nil)

type structureDeployProvider struct {
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &structureDeployProvider{}
	}
}

func (p *structureDeployProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Debug(ctx, "Debug")
}

func (p *structureDeployProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "structure-deploy"
}

func (p *structureDeployProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		SubscriptionDataSource,
		ResourceGroupDataSource,
		CloudProviderDataSource,
		NetworkSecurityGroupDataSource,
		RouteTableGroupDataSource,
		VirtualNetworkGroupDataSource,
	}
}

func (p *structureDeployProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		SubscriptionResource,
		ResourceGroupResource,
		CloudProviderResource,
		NetworkSecurityGroupResource,
		RouteTableResource,
		VirtualNetworkResource,
	}
}

func (p *structureDeployProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	// resp.Schema = schema.Schema{}

}

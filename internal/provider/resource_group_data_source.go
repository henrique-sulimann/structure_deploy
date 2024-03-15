package provider

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &resourceGroupDataSource{}

// var _ resource.ResourceWithConfigure = &subscriptionResource{}

type ResourceGroupDataSourceTenantSettings struct {
	Currency string `json:"currency"`
	TenantID string `json:"tenant_id"`
}
type ResourceGroupDataSourceSubscriptionSettings struct {
	SubscriptionID string `json:"subscription_id"`
}
type ResourceGroupDataSourceTenantProperties struct {
	Settings ResourceGroupDataSourceTenantSettings `json:"settings"`
}
type ResourceGroupDataSourceSubscriptionProperties struct {
	Settings ResourceGroupDataSourceSubscriptionSettings `json:"settings"`
}
type ResourceGroupDataSourceTenant struct {
	Alias       string                                  `json:"alias"`
	Description string                                  `json:"description"`
	ID          string                                  `json:"id"`
	Kind        string                                  `json:"kind"`
	Labels      map[string]interface{}                  `json:"labels"`
	Name        string                                  `json:"name"`
	Properties  ResourceGroupDataSourceTenantProperties `json:"properties"`
}
type ResourceGroupDataSourceSubscription struct {
	Alias       string                                        `json:"alias"`
	Description string                                        `json:"description"`
	ID          string                                        `json:"id"`
	Kind        string                                        `json:"kind"`
	Labels      map[string]interface{}                        `json:"labels"`
	Name        string                                        `json:"name"`
	Properties  ResourceGroupDataSourceSubscriptionProperties `json:"properties"`
}
type ResourceGroupDataSourceCloudProvider struct {
	Alias       string                 `json:"alias"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Kind        string                 `json:"kind"`
	Labels      map[string]interface{} `json:"labels"`
	Name        string                 `json:"name"`
	Properties  map[string]interface{} `json:"properties"`
}
type ResourceGroupDataSourceReferences struct {
	Cloudprovider ResourceGroupDataSourceCloudProvider `json:"cloudprovider"`
	Tenant        ResourceGroupDataSourceTenant        `json:"tenant"`
	Subscription  ResourceGroupDataSourceSubscription  `json:"subscription"`
}
type ResourceGroupDataSourcePropertiesSettings struct {
	Continent string `json:"continent"`
}
type ResourceGroupDataSourceProperties struct {
	References ResourceGroupDataSourceReferences         `json:"references"`
	Settings   ResourceGroupDataSourcePropertiesSettings `json:"settings"`
}

type ResourceGroupDataSourceAPIModelResponseBody struct {
	Alias       string                            `json:"alias"`
	Properties  ResourceGroupDataSourceProperties `json:"properties"`
	Description string                            `json:"description"`
	Kind        string                            `json:"kind"`
	Labels      map[string]interface{}            `json:"labels"`
	Name        string                            `json:"name"`
	ID          string                            `json:"id"`
}

type ResourceGroupDataSourceTenantPropertiesSettingsTF struct {
	Currency types.String `tfsdk:"currency"`
	TenantID types.String `tfsdk:"tenant_id"`
}
type ResourceGroupDataSourceTenantPropertiesTF struct {
	Settings *ResourceGroupDataSourceTenantPropertiesSettingsTF `tfsdk:"settings"`
}

type ResourceGroupDataSourceSubscriptionPropertiesSettingsTF struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
}
type ResourceGroupDataSourceSubscriptionPropertiesTF struct {
	Settings *ResourceGroupDataSourceSubscriptionPropertiesSettingsTF `tfsdk:"settings"`
}

type ResourceGroupDataSourceTenantTF struct {
	Alias       types.String                               `tfsdk:"alias"`
	Description types.String                               `tfsdk:"description"`
	ID          types.String                               `tfsdk:"id"`
	Kind        types.String                               `tfsdk:"kind"`
	Labels      types.Map                                  `tfsdk:"labels"`
	Name        types.String                               `tfsdk:"name"`
	Properties  *ResourceGroupDataSourceTenantPropertiesTF `tfsdk:"properties"`
}
type ResourceGroupDataSourceSubscriptionTF struct {
	Alias       types.String                                     `tfsdk:"alias"`
	Description types.String                                     `tfsdk:"description"`
	ID          types.String                                     `tfsdk:"id"`
	Kind        types.String                                     `tfsdk:"kind"`
	Labels      types.Map                                        `tfsdk:"labels"`
	Name        types.String                                     `tfsdk:"name"`
	Properties  *ResourceGroupDataSourceSubscriptionPropertiesTF `tfsdk:"properties"`
}
type ResourceGroupDataSourceCloudProviderTF struct {
	Alias       types.String `tfsdk:"alias"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Kind        types.String `tfsdk:"kind"`
	Labels      types.Map    `tfsdk:"labels"`
	Name        types.String `tfsdk:"name"`
	Properties  types.Map    `tfsdk:"properties"`
}
type ResourceGroupDataSourceReferencesTF struct {
	Cloudprovider *ResourceGroupDataSourceCloudProviderTF `tfsdk:"cloudprovider"`
	Tenant        *ResourceGroupDataSourceTenantTF        `tfsdk:"tenant"`
	Subscription  *ResourceGroupDataSourceSubscriptionTF  `tfsdk:"subscription"`
}

type ResourceGroupDataSourcePropertiesSettingsTF struct {
	Continent types.String `tfsdk:"continent"`
}
type ResourceGroupDataSourcePropertiesTF struct {
	References *ResourceGroupDataSourceReferencesTF         `tfsdk:"references"`
	Settings   *ResourceGroupDataSourcePropertiesSettingsTF `tfsdk:"settings"`
}

type ResourceGroupDataSourceTerraformModel struct {
	Alias       types.String                         `tfsdk:"alias"`
	Properties  *ResourceGroupDataSourcePropertiesTF `tfsdk:"properties"`
	Description types.String                         `tfsdk:"description"`
	Kind        types.String                         `tfsdk:"kind"`
	Labels      types.Map                            `tfsdk:"labels"`
	Name        types.String                         `tfsdk:"name"`
	ID          types.String                         `tfsdk:"id"`
}

type resourceGroupDataSource struct {
	provider structureDeployProvider
}

func ResourceGroupDataSource() datasource.DataSource {
	return &resourceGroupDataSource{}
}

func (e *resourceGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_group"
}
func (e *resourceGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"kind": schema.StringAttribute{
				Computed: true,
			},
			"labels": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"properties": schema.ObjectAttribute{
				Computed: true,
				AttributeTypes: map[string]attr.Type{
					"settings": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"continent": types.StringType,
						},
					},
					"references": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"cloudprovider": types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"alias":       types.StringType,
									"description": types.StringType,
									"id":          types.StringType,
									"kind":        types.StringType,
									"labels": types.MapType{
										ElemType: types.StringType,
									},
									"name": types.StringType,
									"properties": types.MapType{
										ElemType: types.StringType,
									},
								},
							},
							"tenant": types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"alias":       types.StringType,
									"description": types.StringType,
									"id":          types.StringType,
									"kind":        types.StringType,
									"labels": types.MapType{
										ElemType: types.StringType,
									},
									"name": types.StringType,
									"properties": types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"settings": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"currency":  types.StringType,
													"tenant_id": types.StringType,
												},
											},
										},
									},
								},
							},
							"subscription": types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"alias":       types.StringType,
									"description": types.StringType,
									"id":          types.StringType,
									"kind":        types.StringType,
									"labels": types.MapType{
										ElemType: types.StringType,
									},
									"name": types.StringType,
									"properties": types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"settings": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"subscription_id": types.StringType,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func (e *resourceGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var tfstate ResourceGroupDataSourceTerraformModel

	diags := req.Config.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Read resource using 3rd party API.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestData, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://localhost:8443/api/v1/resourcegroup/%s", tfstate.Name.ValueString()), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Get Resource",
			"An unexpected error occurred while get de resource information "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	client := &http.Client{}
	responseData, err := client.Do(requestData)
	body, err := ioutil.ReadAll(responseData.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Response Body",
			"An unexpected error occurred while read the response body. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	if responseData.StatusCode < 200 || responseData.StatusCode >= 300 {
		resp.Diagnostics.AddError(
			"Error From API Server",
			"An unexpected status code return from API Server "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+string(body),
		)
		return
	}
	fmt.Println("Reponse Body: " + string(body))
	defer func() {
		if responseData != nil && responseData.Body != nil {
			if cerr := responseData.Body.Close(); cerr != nil {
				resp.Diagnostics.AddError(
					"Error while Close Response Body",
					"An unexpected error occurred while Close response Body. "+
						"Please report this issue to the provider developers.\n\n"+
						"JSON Error: "+err.Error(),
				)
				return
			}
		}
	}()
	var responseBody ResourceGroupDataSourceAPIModelResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error while unmarshall de responseBody",
			"An unexpected error occurred while unmarshall de responseBody "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	tfstate = ResourceGroupDataSourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        types.StringValue(responseBody.Kind),
		Properties: &ResourceGroupDataSourcePropertiesTF{
			Settings: &ResourceGroupDataSourcePropertiesSettingsTF{
				Continent: types.StringValue(responseBody.Properties.Settings.Continent),
			},
			References: &ResourceGroupDataSourceReferencesTF{
				Cloudprovider: &ResourceGroupDataSourceCloudProviderTF{
					Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
					Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
					ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
				},
				Tenant: &ResourceGroupDataSourceTenantTF{
					Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
					Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
					ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
					Properties: &ResourceGroupDataSourceTenantPropertiesTF{
						Settings: &ResourceGroupDataSourceTenantPropertiesSettingsTF{
							Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
							TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
						},
					},
				},
				Subscription: &ResourceGroupDataSourceSubscriptionTF{
					Alias:       types.StringValue(responseBody.Properties.References.Subscription.Alias),
					Description: types.StringValue(responseBody.Properties.References.Subscription.Description),
					ID:          types.StringValue(responseBody.Properties.References.Subscription.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Subscription.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Subscription.Name),
					Properties: &ResourceGroupDataSourceSubscriptionPropertiesTF{
						Settings: &ResourceGroupDataSourceSubscriptionPropertiesSettingsTF{
							SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
						},
					},
				},
			},
		},
	}

	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	tfstate.Properties.References.Cloudprovider.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Labels)
	tfstate.Properties.References.Cloudprovider.Properties, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Properties)
	tfstate.Properties.References.Tenant.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Tenant.Labels)
	tfstate.Properties.References.Subscription.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Subscription.Labels)
	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

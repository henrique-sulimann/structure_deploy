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

var _ datasource.DataSource = &routeTableGroupDataSource{}

// var _ resource.ResourceWithConfigure = &subscriptionResource{}

type RouteTableGroupDataSourceTenantSettings struct {
	Currency string `json:"currency"`
	TenantID string `json:"tenant_id"`
}
type RouteTableGroupDataSourceSubscriptionSettings struct {
	SubscriptionID string `json:"subscription_id"`
}
type RouteTableGroupDataSourceResourceGroupSettings struct {
	Continent string `json:"continent"`
}
type RouteTableGroupDataSourceTenantProperties struct {
	Settings RouteTableGroupDataSourceTenantSettings `json:"settings"`
}
type RouteTableGroupDataSourceSubscriptionProperties struct {
	Settings RouteTableGroupDataSourceSubscriptionSettings `json:"settings"`
}
type RouteTableGroupDataSourceResourceGroupProperties struct {
	Settings RouteTableGroupDataSourceResourceGroupSettings `json:"settings"`
}
type RouteTableGroupDataSourceTenant struct {
	Alias       string                                    `json:"alias"`
	Description string                                    `json:"description"`
	ID          string                                    `json:"id"`
	Kind        string                                    `json:"kind"`
	Labels      map[string]interface{}                    `json:"labels"`
	Name        string                                    `json:"name"`
	Properties  RouteTableGroupDataSourceTenantProperties `json:"properties"`
}
type RouteTableGroupDataSourceSubscription struct {
	Alias       string                                          `json:"alias"`
	Description string                                          `json:"description"`
	ID          string                                          `json:"id"`
	Kind        string                                          `json:"kind"`
	Labels      map[string]interface{}                          `json:"labels"`
	Name        string                                          `json:"name"`
	Properties  RouteTableGroupDataSourceSubscriptionProperties `json:"properties"`
}
type RouteTableGroupDataSourceResourceGroup struct {
	Alias       string                                           `json:"alias"`
	Description string                                           `json:"description"`
	ID          string                                           `json:"id"`
	Kind        string                                           `json:"kind"`
	Labels      map[string]interface{}                           `json:"labels"`
	Name        string                                           `json:"name"`
	Properties  RouteTableGroupDataSourceResourceGroupProperties `json:"properties"`
}
type RouteTableGroupDataSourceCloudProvider struct {
	Alias       string                 `json:"alias"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Kind        string                 `json:"kind"`
	Labels      map[string]interface{} `json:"labels"`
	Name        string                 `json:"name"`
	Properties  map[string]interface{} `json:"properties"`
}
type RouteTableGroupDataSourceReferences struct {
	Cloudprovider RouteTableGroupDataSourceCloudProvider `json:"cloudprovider"`
	Tenant        RouteTableGroupDataSourceTenant        `json:"tenant"`
	Subscription  RouteTableGroupDataSourceSubscription  `json:"subscription"`
	ResourceGroup RouteTableGroupDataSourceResourceGroup `json:"resourcegroup"`
}
type RouteTableGroupDataSourcePropertiesSettings struct {
	RoutePropagation string `json:"route_propagation"`
	Status           string `json:"status"`
}
type RouteTableGroupDataSourceProperties struct {
	References RouteTableGroupDataSourceReferences         `json:"references"`
	Settings   RouteTableGroupDataSourcePropertiesSettings `json:"settings"`
}

type RouteTableGroupDataSourceAPIModelResponseBody struct {
	Alias       string                              `json:"alias"`
	Properties  RouteTableGroupDataSourceProperties `json:"properties"`
	Description string                              `json:"description"`
	Kind        string                              `json:"kind"`
	Labels      map[string]interface{}              `json:"labels"`
	Name        string                              `json:"name"`
	ID          string                              `json:"id"`
}

type RouteTableGroupDataSourceTenantPropertiesSettingsTF struct {
	Currency types.String `tfsdk:"currency"`
	TenantID types.String `tfsdk:"tenant_id"`
}
type RouteTableGroupDataSourceTenantPropertiesTF struct {
	Settings *RouteTableGroupDataSourceTenantPropertiesSettingsTF `tfsdk:"settings"`
}

type RouteTableGroupDataSourceSubscriptionPropertiesSettingsTF struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
}
type RouteTableGroupDataSourceSubscriptionPropertiesTF struct {
	Settings *RouteTableGroupDataSourceSubscriptionPropertiesSettingsTF `tfsdk:"settings"`
}
type RouteTableGroupDataSourceResourceGroupPropertiesSettingsTF struct {
	Continent types.String `tfsdk:"continent"`
}
type RouteTableGroupDataSourceResourceGroupPropertiesTF struct {
	Settings *RouteTableGroupDataSourceResourceGroupPropertiesSettingsTF `tfsdk:"settings"`
}

type RouteTableGroupDataSourceTenantTF struct {
	Alias       types.String                                 `tfsdk:"alias"`
	Description types.String                                 `tfsdk:"description"`
	ID          types.String                                 `tfsdk:"id"`
	Kind        types.String                                 `tfsdk:"kind"`
	Labels      types.Map                                    `tfsdk:"labels"`
	Name        types.String                                 `tfsdk:"name"`
	Properties  *RouteTableGroupDataSourceTenantPropertiesTF `tfsdk:"properties"`
}
type RouteTableGroupDataSourceSubscriptionTF struct {
	Alias       types.String                                       `tfsdk:"alias"`
	Description types.String                                       `tfsdk:"description"`
	ID          types.String                                       `tfsdk:"id"`
	Kind        types.String                                       `tfsdk:"kind"`
	Labels      types.Map                                          `tfsdk:"labels"`
	Name        types.String                                       `tfsdk:"name"`
	Properties  *RouteTableGroupDataSourceSubscriptionPropertiesTF `tfsdk:"properties"`
}
type RouteTableGroupDataSourceResourceGroupTF struct {
	Alias       types.String                                        `tfsdk:"alias"`
	Description types.String                                        `tfsdk:"description"`
	ID          types.String                                        `tfsdk:"id"`
	Kind        types.String                                        `tfsdk:"kind"`
	Labels      types.Map                                           `tfsdk:"labels"`
	Name        types.String                                        `tfsdk:"name"`
	Properties  *RouteTableGroupDataSourceResourceGroupPropertiesTF `tfsdk:"properties"`
}
type RouteTableGroupDataSourceCloudProviderTF struct {
	Alias       types.String `tfsdk:"alias"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Kind        types.String `tfsdk:"kind"`
	Labels      types.Map    `tfsdk:"labels"`
	Name        types.String `tfsdk:"name"`
	Properties  types.Map    `tfsdk:"properties"`
}
type RouteTableGroupDataSourceReferencesTF struct {
	Cloudprovider *RouteTableGroupDataSourceCloudProviderTF `tfsdk:"cloudprovider"`
	Tenant        *RouteTableGroupDataSourceTenantTF        `tfsdk:"tenant"`
	Subscription  *RouteTableGroupDataSourceSubscriptionTF  `tfsdk:"subscription"`
	ResourceGroup *RouteTableGroupDataSourceResourceGroupTF `tfsdk:"resourcegroup"`
}

type RouteTableGroupDataSourcePropertiesSettingsTF struct {
	RoutePropagation types.String `tfsdk:"route_propagation"`
	Status           types.String `tfsdk:"status"`
}
type RouteTableGroupDataSourcePropertiesTF struct {
	References *RouteTableGroupDataSourceReferencesTF         `tfsdk:"references"`
	Settings   *RouteTableGroupDataSourcePropertiesSettingsTF `tfsdk:"settings"`
}

type RouteTableGroupDataSourceTerraformModel struct {
	Alias       types.String                           `tfsdk:"alias"`
	Properties  *RouteTableGroupDataSourcePropertiesTF `tfsdk:"properties"`
	Description types.String                           `tfsdk:"description"`
	Kind        types.String                           `tfsdk:"kind"`
	Labels      types.Map                              `tfsdk:"labels"`
	Name        types.String                           `tfsdk:"name"`
	ID          types.String                           `tfsdk:"id"`
}

type routeTableGroupDataSource struct {
	provider structureDeployProvider
}

func RouteTableGroupDataSource() datasource.DataSource {
	return &routeTableGroupDataSource{}
}

func (e *routeTableGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_table"
}
func (e *routeTableGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
							"status":            types.StringType,
							"route_propagation": types.StringType,
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
							"resourcegroup": types.ObjectType{
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
													"continent": types.StringType,
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
func (e *routeTableGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var tfstate RouteTableGroupDataSourceTerraformModel

	diags := req.Config.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Read resource using 3rd party API.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestData, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://localhost:8443/api/v1/route_table/%s", tfstate.Name.ValueString()), nil)
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
	var responseBody RouteTableGroupDataSourceAPIModelResponseBody
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
	tfstate = RouteTableGroupDataSourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        types.StringValue(responseBody.Kind),
		Properties: &RouteTableGroupDataSourcePropertiesTF{
			Settings: &RouteTableGroupDataSourcePropertiesSettingsTF{
				Status:           types.StringValue(responseBody.Properties.Settings.Status),
				RoutePropagation: types.StringValue(responseBody.Properties.Settings.RoutePropagation),
			},
			References: &RouteTableGroupDataSourceReferencesTF{
				Cloudprovider: &RouteTableGroupDataSourceCloudProviderTF{
					Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
					Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
					ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
				},
				Tenant: &RouteTableGroupDataSourceTenantTF{
					Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
					Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
					ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
					Properties: &RouteTableGroupDataSourceTenantPropertiesTF{
						Settings: &RouteTableGroupDataSourceTenantPropertiesSettingsTF{
							Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
							TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
						},
					},
				},
				Subscription: &RouteTableGroupDataSourceSubscriptionTF{
					Alias:       types.StringValue(responseBody.Properties.References.Subscription.Alias),
					Description: types.StringValue(responseBody.Properties.References.Subscription.Description),
					ID:          types.StringValue(responseBody.Properties.References.Subscription.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Subscription.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Subscription.Name),
					Properties: &RouteTableGroupDataSourceSubscriptionPropertiesTF{
						Settings: &RouteTableGroupDataSourceSubscriptionPropertiesSettingsTF{
							SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
						},
					},
				},
				ResourceGroup: &RouteTableGroupDataSourceResourceGroupTF{
					Alias:       types.StringValue(responseBody.Properties.References.ResourceGroup.Alias),
					Description: types.StringValue(responseBody.Properties.References.ResourceGroup.Description),
					ID:          types.StringValue(responseBody.Properties.References.ResourceGroup.ID),
					Kind:        types.StringValue(responseBody.Properties.References.ResourceGroup.Kind),
					Name:        types.StringValue(responseBody.Properties.References.ResourceGroup.Name),
					Properties: &RouteTableGroupDataSourceResourceGroupPropertiesTF{
						Settings: &RouteTableGroupDataSourceResourceGroupPropertiesSettingsTF{
							Continent: types.StringValue(responseBody.Properties.References.ResourceGroup.Properties.Settings.Continent),
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
	tfstate.Properties.References.ResourceGroup.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.ResourceGroup.Labels)
	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

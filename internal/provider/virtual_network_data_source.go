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

var _ datasource.DataSource = &virtualNetworkGroupDataSource{}

// var _ resource.ResourceWithConfigure = &subscriptionResource{}

type VirtualNetworkGroupDataSourceTenantSettings struct {
	Currency string `json:"currency"`
	TenantID string `json:"tenant_id"`
}
type VirtualNetworkGroupDataSourceSubscriptionSettings struct {
	SubscriptionID string `json:"subscription_id"`
}
type VirtualNetworkGroupDataSourceResourceGroupSettings struct {
	Continent string `json:"continent"`
}
type VirtualNetworkGroupDataSourceTenantProperties struct {
	Settings VirtualNetworkGroupDataSourceTenantSettings `json:"settings"`
}
type VirtualNetworkGroupDataSourceSubscriptionProperties struct {
	Settings VirtualNetworkGroupDataSourceSubscriptionSettings `json:"settings"`
}
type VirtualNetworkGroupDataSourceResourceGroupProperties struct {
	Settings VirtualNetworkGroupDataSourceResourceGroupSettings `json:"settings"`
}
type VirtualNetworkGroupDataSourceTenant struct {
	Alias       string                                        `json:"alias"`
	Description string                                        `json:"description"`
	ID          string                                        `json:"id"`
	Kind        string                                        `json:"kind"`
	Labels      map[string]interface{}                        `json:"labels"`
	Name        string                                        `json:"name"`
	Properties  VirtualNetworkGroupDataSourceTenantProperties `json:"properties"`
}
type VirtualNetworkGroupDataSourceSubscription struct {
	Alias       string                                              `json:"alias"`
	Description string                                              `json:"description"`
	ID          string                                              `json:"id"`
	Kind        string                                              `json:"kind"`
	Labels      map[string]interface{}                              `json:"labels"`
	Name        string                                              `json:"name"`
	Properties  VirtualNetworkGroupDataSourceSubscriptionProperties `json:"properties"`
}
type VirtualNetworkGroupDataSourceResourceGroup struct {
	Alias       string                                               `json:"alias"`
	Description string                                               `json:"description"`
	ID          string                                               `json:"id"`
	Kind        string                                               `json:"kind"`
	Labels      map[string]interface{}                               `json:"labels"`
	Name        string                                               `json:"name"`
	Properties  VirtualNetworkGroupDataSourceResourceGroupProperties `json:"properties"`
}
type VirtualNetworkGroupDataSourceCloudProvider struct {
	Alias       string                 `json:"alias"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Kind        string                 `json:"kind"`
	Labels      map[string]interface{} `json:"labels"`
	Name        string                 `json:"name"`
	Properties  map[string]interface{} `json:"properties"`
}
type VirtualNetworkGroupDataSourceReferences struct {
	Cloudprovider VirtualNetworkGroupDataSourceCloudProvider `json:"cloudprovider"`
	Tenant        VirtualNetworkGroupDataSourceTenant        `json:"tenant"`
	Subscription  VirtualNetworkGroupDataSourceSubscription  `json:"subscription"`
	ResourceGroup VirtualNetworkGroupDataSourceResourceGroup `json:"resourcegroup"`
}
type VirtualNetworkGroupDataSourcePropertiesSettings struct {
	Cidr   []string `json:"cidr"`
	Status string   `json:"status"`
}
type VirtualNetworkGroupDataSourceProperties struct {
	References VirtualNetworkGroupDataSourceReferences         `json:"references"`
	Settings   VirtualNetworkGroupDataSourcePropertiesSettings `json:"settings"`
}

type VirtualNetworkGroupDataSourceAPIModelResponseBody struct {
	Alias       string                                  `json:"alias"`
	Properties  VirtualNetworkGroupDataSourceProperties `json:"properties"`
	Description string                                  `json:"description"`
	Kind        string                                  `json:"kind"`
	Labels      map[string]interface{}                  `json:"labels"`
	Name        string                                  `json:"name"`
	ID          string                                  `json:"id"`
}

type VirtualNetworkGroupDataSourceTenantPropertiesSettingsTF struct {
	Currency types.String `tfsdk:"currency"`
	TenantID types.String `tfsdk:"tenant_id"`
}
type VirtualNetworkGroupDataSourceTenantPropertiesTF struct {
	Settings *VirtualNetworkGroupDataSourceTenantPropertiesSettingsTF `tfsdk:"settings"`
}

type VirtualNetworkGroupDataSourceSubscriptionPropertiesSettingsTF struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
}
type VirtualNetworkGroupDataSourceSubscriptionPropertiesTF struct {
	Settings *VirtualNetworkGroupDataSourceSubscriptionPropertiesSettingsTF `tfsdk:"settings"`
}
type VirtualNetworkGroupDataSourceResourceGroupPropertiesSettingsTF struct {
	Continent types.String `tfsdk:"continent"`
}
type VirtualNetworkGroupDataSourceResourceGroupPropertiesTF struct {
	Settings *VirtualNetworkGroupDataSourceResourceGroupPropertiesSettingsTF `tfsdk:"settings"`
}

type VirtualNetworkGroupDataSourceTenantTF struct {
	Alias       types.String                                     `tfsdk:"alias"`
	Description types.String                                     `tfsdk:"description"`
	ID          types.String                                     `tfsdk:"id"`
	Kind        types.String                                     `tfsdk:"kind"`
	Labels      types.Map                                        `tfsdk:"labels"`
	Name        types.String                                     `tfsdk:"name"`
	Properties  *VirtualNetworkGroupDataSourceTenantPropertiesTF `tfsdk:"properties"`
}
type VirtualNetworkGroupDataSourceSubscriptionTF struct {
	Alias       types.String                                           `tfsdk:"alias"`
	Description types.String                                           `tfsdk:"description"`
	ID          types.String                                           `tfsdk:"id"`
	Kind        types.String                                           `tfsdk:"kind"`
	Labels      types.Map                                              `tfsdk:"labels"`
	Name        types.String                                           `tfsdk:"name"`
	Properties  *VirtualNetworkGroupDataSourceSubscriptionPropertiesTF `tfsdk:"properties"`
}
type VirtualNetworkGroupDataSourceResourceGroupTF struct {
	Alias       types.String                                            `tfsdk:"alias"`
	Description types.String                                            `tfsdk:"description"`
	ID          types.String                                            `tfsdk:"id"`
	Kind        types.String                                            `tfsdk:"kind"`
	Labels      types.Map                                               `tfsdk:"labels"`
	Name        types.String                                            `tfsdk:"name"`
	Properties  *VirtualNetworkGroupDataSourceResourceGroupPropertiesTF `tfsdk:"properties"`
}
type VirtualNetworkGroupDataSourceCloudProviderTF struct {
	Alias       types.String `tfsdk:"alias"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Kind        types.String `tfsdk:"kind"`
	Labels      types.Map    `tfsdk:"labels"`
	Name        types.String `tfsdk:"name"`
	Properties  types.Map    `tfsdk:"properties"`
}
type VirtualNetworkGroupDataSourceReferencesTF struct {
	Cloudprovider *VirtualNetworkGroupDataSourceCloudProviderTF `tfsdk:"cloudprovider"`
	Tenant        *VirtualNetworkGroupDataSourceTenantTF        `tfsdk:"tenant"`
	Subscription  *VirtualNetworkGroupDataSourceSubscriptionTF  `tfsdk:"subscription"`
	ResourceGroup *VirtualNetworkGroupDataSourceResourceGroupTF `tfsdk:"resourcegroup"`
}

type VirtualNetworkGroupDataSourcePropertiesSettingsTF struct {
	Cidr   types.List   `tfsdk:"cidr"`
	Status types.String `tfsdk:"status"`
}
type VirtualNetworkGroupDataSourcePropertiesTF struct {
	References *VirtualNetworkGroupDataSourceReferencesTF         `tfsdk:"references"`
	Settings   *VirtualNetworkGroupDataSourcePropertiesSettingsTF `tfsdk:"settings"`
}

type VirtualNetworkGroupDataSourceTerraformModel struct {
	Alias       types.String                               `tfsdk:"alias"`
	Properties  *VirtualNetworkGroupDataSourcePropertiesTF `tfsdk:"properties"`
	Description types.String                               `tfsdk:"description"`
	Kind        types.String                               `tfsdk:"kind"`
	Labels      types.Map                                  `tfsdk:"labels"`
	Name        types.String                               `tfsdk:"name"`
	ID          types.String                               `tfsdk:"id"`
}

type virtualNetworkGroupDataSource struct {
	provider structureDeployProvider
}

func VirtualNetworkGroupDataSource() datasource.DataSource {
	return &virtualNetworkGroupDataSource{}
}

func (e *virtualNetworkGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_network"
}
func (e *virtualNetworkGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
							"status": types.StringType,
							"cidr": types.ListType{
								ElemType: types.StringType,
							},
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
func (e *virtualNetworkGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var tfstate VirtualNetworkGroupDataSourceTerraformModel

	diags := req.Config.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Read resource using 3rd party API.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestData, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://localhost:8443/api/v1/virtualnetwork/%s", tfstate.Name.ValueString()), nil)
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
	var responseBody VirtualNetworkGroupDataSourceAPIModelResponseBody
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
	tfstate = VirtualNetworkGroupDataSourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        types.StringValue(responseBody.Kind),
		Properties: &VirtualNetworkGroupDataSourcePropertiesTF{
			Settings: &VirtualNetworkGroupDataSourcePropertiesSettingsTF{
				Status: types.StringValue(responseBody.Properties.Settings.Status),
			},
			References: &VirtualNetworkGroupDataSourceReferencesTF{
				Cloudprovider: &VirtualNetworkGroupDataSourceCloudProviderTF{
					Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
					Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
					ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
				},
				Tenant: &VirtualNetworkGroupDataSourceTenantTF{
					Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
					Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
					ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
					Properties: &VirtualNetworkGroupDataSourceTenantPropertiesTF{
						Settings: &VirtualNetworkGroupDataSourceTenantPropertiesSettingsTF{
							Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
							TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
						},
					},
				},
				Subscription: &VirtualNetworkGroupDataSourceSubscriptionTF{
					Alias:       types.StringValue(responseBody.Properties.References.Subscription.Alias),
					Description: types.StringValue(responseBody.Properties.References.Subscription.Description),
					ID:          types.StringValue(responseBody.Properties.References.Subscription.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Subscription.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Subscription.Name),
					Properties: &VirtualNetworkGroupDataSourceSubscriptionPropertiesTF{
						Settings: &VirtualNetworkGroupDataSourceSubscriptionPropertiesSettingsTF{
							SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
						},
					},
				},
				ResourceGroup: &VirtualNetworkGroupDataSourceResourceGroupTF{
					Alias:       types.StringValue(responseBody.Properties.References.ResourceGroup.Alias),
					Description: types.StringValue(responseBody.Properties.References.ResourceGroup.Description),
					ID:          types.StringValue(responseBody.Properties.References.ResourceGroup.ID),
					Kind:        types.StringValue(responseBody.Properties.References.ResourceGroup.Kind),
					Name:        types.StringValue(responseBody.Properties.References.ResourceGroup.Name),
					Properties: &VirtualNetworkGroupDataSourceResourceGroupPropertiesTF{
						Settings: &VirtualNetworkGroupDataSourceResourceGroupPropertiesSettingsTF{
							Continent: types.StringValue(responseBody.Properties.References.ResourceGroup.Properties.Settings.Continent),
						},
					},
				},
			},
		},
	}

	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	tfstate.Properties.Settings.Cidr, diags = types.ListValueFrom(ctx, types.StringType, responseBody.Properties.Settings.Cidr)
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

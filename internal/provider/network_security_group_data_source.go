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

var _ datasource.DataSource = &networkSecurityGroupDataSource{}

// var _ resource.ResourceWithConfigure = &subscriptionResource{}

type NetworkSecurityGroupDataSourceTenantSettings struct {
	Currency string `json:"currency"`
	TenantID string `json:"tenant_id"`
}
type NetworkSecurityGroupDataSourceSubscriptionSettings struct {
	SubscriptionID string `json:"subscription_id"`
}
type NetworkSecurityGroupDataSourceResourceGroupSettings struct {
	Continent string `json:"continent"`
}
type NetworkSecurityGroupDataSourceTenantProperties struct {
	Settings NetworkSecurityGroupDataSourceTenantSettings `json:"settings"`
}
type NetworkSecurityGroupDataSourceSubscriptionProperties struct {
	Settings NetworkSecurityGroupDataSourceSubscriptionSettings `json:"settings"`
}
type NetworkSecurityGroupDataSourceResourceGroupProperties struct {
	Settings NetworkSecurityGroupDataSourceResourceGroupSettings `json:"settings"`
}
type NetworkSecurityGroupDataSourceTenant struct {
	Alias       string                                         `json:"alias"`
	Description string                                         `json:"description"`
	ID          string                                         `json:"id"`
	Kind        string                                         `json:"kind"`
	Labels      map[string]interface{}                         `json:"labels"`
	Name        string                                         `json:"name"`
	Properties  NetworkSecurityGroupDataSourceTenantProperties `json:"properties"`
}
type NetworkSecurityGroupDataSourceSubscription struct {
	Alias       string                                               `json:"alias"`
	Description string                                               `json:"description"`
	ID          string                                               `json:"id"`
	Kind        string                                               `json:"kind"`
	Labels      map[string]interface{}                               `json:"labels"`
	Name        string                                               `json:"name"`
	Properties  NetworkSecurityGroupDataSourceSubscriptionProperties `json:"properties"`
}
type NetworkSecurityGroupDataSourceResourceGroup struct {
	Alias       string                                                `json:"alias"`
	Description string                                                `json:"description"`
	ID          string                                                `json:"id"`
	Kind        string                                                `json:"kind"`
	Labels      map[string]interface{}                                `json:"labels"`
	Name        string                                                `json:"name"`
	Properties  NetworkSecurityGroupDataSourceResourceGroupProperties `json:"properties"`
}
type NetworkSecurityGroupDataSourceCloudProvider struct {
	Alias       string                 `json:"alias"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Kind        string                 `json:"kind"`
	Labels      map[string]interface{} `json:"labels"`
	Name        string                 `json:"name"`
	Properties  map[string]interface{} `json:"properties"`
}
type NetworkSecurityGroupDataSourceReferences struct {
	Cloudprovider NetworkSecurityGroupDataSourceCloudProvider `json:"cloudprovider"`
	Tenant        NetworkSecurityGroupDataSourceTenant        `json:"tenant"`
	Subscription  NetworkSecurityGroupDataSourceSubscription  `json:"subscription"`
	ResourceGroup NetworkSecurityGroupDataSourceResourceGroup `json:"resourcegroup"`
}
type NetworkSecurityGroupDataSourcePropertiesSettings struct {
	Status string `json:"status"`
}
type NetworkSecurityGroupDataSourceProperties struct {
	References NetworkSecurityGroupDataSourceReferences         `json:"references"`
	Settings   NetworkSecurityGroupDataSourcePropertiesSettings `json:"settings"`
}

type NetworkSecurityGroupDataSourceAPIModelResponseBody struct {
	Alias       string                                   `json:"alias"`
	Properties  NetworkSecurityGroupDataSourceProperties `json:"properties"`
	Description string                                   `json:"description"`
	Kind        string                                   `json:"kind"`
	Labels      map[string]interface{}                   `json:"labels"`
	Name        string                                   `json:"name"`
	ID          string                                   `json:"id"`
}

type NetworkSecurityGroupDataSourceTenantPropertiesSettingsTF struct {
	Currency types.String `tfsdk:"currency"`
	TenantID types.String `tfsdk:"tenant_id"`
}
type NetworkSecurityGroupDataSourceTenantPropertiesTF struct {
	Settings *NetworkSecurityGroupDataSourceTenantPropertiesSettingsTF `tfsdk:"settings"`
}

type NetworkSecurityGroupDataSourceSubscriptionPropertiesSettingsTF struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
}
type NetworkSecurityGroupDataSourceSubscriptionPropertiesTF struct {
	Settings *NetworkSecurityGroupDataSourceSubscriptionPropertiesSettingsTF `tfsdk:"settings"`
}
type NetworkSecurityGroupDataSourceResourceGroupPropertiesSettingsTF struct {
	Continent types.String `tfsdk:"continent"`
}
type NetworkSecurityGroupDataSourceResourceGroupPropertiesTF struct {
	Settings *NetworkSecurityGroupDataSourceResourceGroupPropertiesSettingsTF `tfsdk:"settings"`
}

type NetworkSecurityGroupDataSourceTenantTF struct {
	Alias       types.String                                      `tfsdk:"alias"`
	Description types.String                                      `tfsdk:"description"`
	ID          types.String                                      `tfsdk:"id"`
	Kind        types.String                                      `tfsdk:"kind"`
	Labels      types.Map                                         `tfsdk:"labels"`
	Name        types.String                                      `tfsdk:"name"`
	Properties  *NetworkSecurityGroupDataSourceTenantPropertiesTF `tfsdk:"properties"`
}
type NetworkSecurityGroupDataSourceSubscriptionTF struct {
	Alias       types.String                                            `tfsdk:"alias"`
	Description types.String                                            `tfsdk:"description"`
	ID          types.String                                            `tfsdk:"id"`
	Kind        types.String                                            `tfsdk:"kind"`
	Labels      types.Map                                               `tfsdk:"labels"`
	Name        types.String                                            `tfsdk:"name"`
	Properties  *NetworkSecurityGroupDataSourceSubscriptionPropertiesTF `tfsdk:"properties"`
}
type NetworkSecurityGroupDataSourceResourceGroupTF struct {
	Alias       types.String                                             `tfsdk:"alias"`
	Description types.String                                             `tfsdk:"description"`
	ID          types.String                                             `tfsdk:"id"`
	Kind        types.String                                             `tfsdk:"kind"`
	Labels      types.Map                                                `tfsdk:"labels"`
	Name        types.String                                             `tfsdk:"name"`
	Properties  *NetworkSecurityGroupDataSourceResourceGroupPropertiesTF `tfsdk:"properties"`
}
type NetworkSecurityGroupDataSourceCloudProviderTF struct {
	Alias       types.String `tfsdk:"alias"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Kind        types.String `tfsdk:"kind"`
	Labels      types.Map    `tfsdk:"labels"`
	Name        types.String `tfsdk:"name"`
	Properties  types.Map    `tfsdk:"properties"`
}
type NetworkSecurityGroupDataSourceReferencesTF struct {
	Cloudprovider *NetworkSecurityGroupDataSourceCloudProviderTF `tfsdk:"cloudprovider"`
	Tenant        *NetworkSecurityGroupDataSourceTenantTF        `tfsdk:"tenant"`
	Subscription  *NetworkSecurityGroupDataSourceSubscriptionTF  `tfsdk:"subscription"`
	ResourceGroup *NetworkSecurityGroupDataSourceResourceGroupTF `tfsdk:"resourcegroup"`
}

type NetworkSecurityGroupDataSourcePropertiesSettingsTF struct {
	Status types.String `tfsdk:"status"`
}
type NetworkSecurityGroupDataSourcePropertiesTF struct {
	References *NetworkSecurityGroupDataSourceReferencesTF         `tfsdk:"references"`
	Settings   *NetworkSecurityGroupDataSourcePropertiesSettingsTF `tfsdk:"settings"`
}

type NetworkSecurityGroupDataSourceTerraformModel struct {
	Alias       types.String                                `tfsdk:"alias"`
	Properties  *NetworkSecurityGroupDataSourcePropertiesTF `tfsdk:"properties"`
	Description types.String                                `tfsdk:"description"`
	Kind        types.String                                `tfsdk:"kind"`
	Labels      types.Map                                   `tfsdk:"labels"`
	Name        types.String                                `tfsdk:"name"`
	ID          types.String                                `tfsdk:"id"`
}

type networkSecurityGroupDataSource struct {
	provider structureDeployProvider
}

func NetworkSecurityGroupDataSource() datasource.DataSource {
	return &networkSecurityGroupDataSource{}
}

func (e *networkSecurityGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_security_group"
}
func (e *networkSecurityGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
func (e *networkSecurityGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var tfstate NetworkSecurityGroupDataSourceTerraformModel

	diags := req.Config.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Read resource using 3rd party API.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestData, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://localhost:8443/api/v1/nsg/%s", tfstate.Name.ValueString()), nil)
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
	var responseBody NetworkSecurityGroupDataSourceAPIModelResponseBody
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
	tfstate = NetworkSecurityGroupDataSourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        types.StringValue(responseBody.Kind),
		Properties: &NetworkSecurityGroupDataSourcePropertiesTF{
			Settings: &NetworkSecurityGroupDataSourcePropertiesSettingsTF{
				Status: types.StringValue(responseBody.Properties.Settings.Status),
			},
			References: &NetworkSecurityGroupDataSourceReferencesTF{
				Cloudprovider: &NetworkSecurityGroupDataSourceCloudProviderTF{
					Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
					Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
					ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
				},
				Tenant: &NetworkSecurityGroupDataSourceTenantTF{
					Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
					Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
					ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
					Properties: &NetworkSecurityGroupDataSourceTenantPropertiesTF{
						Settings: &NetworkSecurityGroupDataSourceTenantPropertiesSettingsTF{
							Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
							TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
						},
					},
				},
				Subscription: &NetworkSecurityGroupDataSourceSubscriptionTF{
					Alias:       types.StringValue(responseBody.Properties.References.Subscription.Alias),
					Description: types.StringValue(responseBody.Properties.References.Subscription.Description),
					ID:          types.StringValue(responseBody.Properties.References.Subscription.ID),
					Kind:        types.StringValue(responseBody.Properties.References.Subscription.Kind),
					Name:        types.StringValue(responseBody.Properties.References.Subscription.Name),
					Properties: &NetworkSecurityGroupDataSourceSubscriptionPropertiesTF{
						Settings: &NetworkSecurityGroupDataSourceSubscriptionPropertiesSettingsTF{
							SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
						},
					},
				},
				ResourceGroup: &NetworkSecurityGroupDataSourceResourceGroupTF{
					Alias:       types.StringValue(responseBody.Properties.References.ResourceGroup.Alias),
					Description: types.StringValue(responseBody.Properties.References.ResourceGroup.Description),
					ID:          types.StringValue(responseBody.Properties.References.ResourceGroup.ID),
					Kind:        types.StringValue(responseBody.Properties.References.ResourceGroup.Kind),
					Name:        types.StringValue(responseBody.Properties.References.ResourceGroup.Name),
					Properties: &NetworkSecurityGroupDataSourceResourceGroupPropertiesTF{
						Settings: &NetworkSecurityGroupDataSourceResourceGroupPropertiesSettingsTF{
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

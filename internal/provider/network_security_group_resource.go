package provider

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkSecurityGroupResourceSubscriptionSettings struct {
	SubscriptionID string `json:"subscription_id"`
}
type NetworkSecurityGroupResourceSubscriptionProperties struct {
	Settings NetworkSecurityGroupResourceSubscriptionSettings `json:"settings"`
}
type NetworkSecurityGroupResourceResourceGroupSettings struct {
	Continent string `json:"continent"`
}
type NetworkSecurityGroupResourceResourceGroupProperties struct {
	Settings NetworkSecurityGroupResourceResourceGroupSettings `json:"settings"`
}
type NetworkSecurityGroupResourceTenantSettings struct {
	Currency string `json:"currency"`
	TenantID string `json:"tenant_id"`
}
type NetworkSecurityGroupResourceTenantProperties struct {
	Settings NetworkSecurityGroupResourceTenantSettings `json:"settings"`
}
type NetworkSecurityGroupResourceTenant struct {
	Alias       string                                       `json:"alias"`
	Description string                                       `json:"description"`
	ID          string                                       `json:"id"`
	Kind        string                                       `json:"kind"`
	Labels      map[string]interface{}                       `json:"labels"`
	Name        string                                       `json:"name"`
	Properties  NetworkSecurityGroupResourceTenantProperties `json:"properties"`
}
type NetworkSecurityGroupResourceSubscription struct {
	Alias       string                                             `json:"alias"`
	Description string                                             `json:"description"`
	ID          string                                             `json:"id"`
	Kind        string                                             `json:"kind"`
	Labels      map[string]interface{}                             `json:"labels"`
	Name        string                                             `json:"name"`
	Properties  NetworkSecurityGroupResourceSubscriptionProperties `json:"properties"`
}
type NetworkSecurityGroupResourceResourceGroup struct {
	Alias       string                                              `json:"alias"`
	Description string                                              `json:"description"`
	ID          string                                              `json:"id"`
	Kind        string                                              `json:"kind"`
	Labels      map[string]interface{}                              `json:"labels"`
	Name        string                                              `json:"name"`
	Properties  NetworkSecurityGroupResourceResourceGroupProperties `json:"properties"`
}
type NetworkSecurityGroupResourceCloudProvider struct {
	Alias       string                 `json:"alias"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Kind        string                 `json:"kind"`
	Labels      map[string]interface{} `json:"labels"`
	Name        string                 `json:"name"`
	Properties  map[string]interface{} `json:"properties"`
}
type NetworkSecurityGroupResourceReferences struct {
	Cloudprovider NetworkSecurityGroupResourceCloudProvider `json:"cloudprovider"`
	Tenant        NetworkSecurityGroupResourceTenant        `json:"tenant"`
	Subscription  NetworkSecurityGroupResourceSubscription  `json:"subscription"`
	ResourceGroup NetworkSecurityGroupResourceResourceGroup `json:"resourcegroup"`
}
type NetworkSecurityGroupResourcePropertiesSettings struct {
	Status string `json:"status"`
}
type NetworkSecurityGroupResourceProperties struct {
	// References NetworkSecurityGroupResourceReferences         `json:"references"`
	Settings NetworkSecurityGroupResourcePropertiesSettings `json:"settings"`
}
type NetworkSecurityGroupResourcePropertiesRequestBody struct {
	Settings NetworkSecurityGroupResourcePropertiesSettings `json:"settings"`
}
type NetworkSecurityGroupResourceAPIModelRequestBody struct {
	Alias        string                                            `json:"alias"`
	Dependencies map[string]interface{}                            `json:"dependencies"`
	Description  string                                            `json:"description"`
	Kind         string                                            `json:"kind"`
	Labels       map[string]interface{}                            `json:"labels"`
	Name         string                                            `json:"name"`
	ReleaseDate  string                                            `json:"release_date"`
	Version      string                                            `json:"version"`
	ID           string                                            `json:"id"`
	Properties   NetworkSecurityGroupResourcePropertiesRequestBody `json:"properties"`
}
type NetworkSecurityGroupResourceAPIModelResponseBody struct {
	Alias        string                                 `json:"alias"`
	Dependencies map[string]interface{}                 `json:"dependencies"`
	Description  string                                 `json:"description"`
	Kind         string                                 `json:"kind"`
	Labels       map[string]interface{}                 `json:"labels"`
	Name         string                                 `json:"name"`
	ReleaseDate  string                                 `json:"release_date"`
	Version      string                                 `json:"version"`
	ID           string                                 `json:"id"`
	Properties   NetworkSecurityGroupResourceProperties `json:"properties"`
}

type NetworkSecurityGroupResourceSubscriptionPropertiesSettingsTF struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
}
type NetworkSecurityGroupResourceSubscriptionPropertiesTF struct {
	Settings *NetworkSecurityGroupResourceSubscriptionPropertiesSettingsTF `tfsdk:"settings"`
}
type NetworkSecurityGroupResourceResourceGroupPropertiesSettingsTF struct {
	Continent types.String `tfsdk:"continent"`
}
type NetworkSecurityGroupResourceResourceGroupPropertiesTF struct {
	Settings *NetworkSecurityGroupResourceResourceGroupPropertiesSettingsTF `tfsdk:"settings"`
}
type NetworkSecurityGroupResourceTenantPropertiesSettingsTF struct {
	Currency types.String `tfsdk:"currency"`
	TenantID types.String `tfsdk:"tenant_id"`
}
type NetworkSecurityGroupResourceTenantPropertiesTF struct {
	Settings *NetworkSecurityGroupResourceTenantPropertiesSettingsTF `tfsdk:"settings"`
}
type NetworkSecurityGroupResourceTenantTF struct {
	Alias       types.String                                    `tfsdk:"alias"`
	Description types.String                                    `tfsdk:"description"`
	ID          types.String                                    `tfsdk:"id"`
	Kind        types.String                                    `tfsdk:"kind"`
	Labels      types.Map                                       `tfsdk:"labels"`
	Name        types.String                                    `tfsdk:"name"`
	Properties  *NetworkSecurityGroupResourceTenantPropertiesTF `tfsdk:"properties"`
}
type NetworkSecurityGroupResourceSubscriptionTF struct {
	Alias       types.String                                          `tfsdk:"alias"`
	Description types.String                                          `tfsdk:"description"`
	ID          types.String                                          `tfsdk:"id"`
	Kind        types.String                                          `tfsdk:"kind"`
	Labels      types.Map                                             `tfsdk:"labels"`
	Name        types.String                                          `tfsdk:"name"`
	Properties  *NetworkSecurityGroupResourceSubscriptionPropertiesTF `tfsdk:"properties"`
}
type NetworkSecurityGroupResourceResourceGroupTF struct {
	Alias       types.String                                           `tfsdk:"alias"`
	Description types.String                                           `tfsdk:"description"`
	ID          types.String                                           `tfsdk:"id"`
	Kind        types.String                                           `tfsdk:"kind"`
	Labels      types.Map                                              `tfsdk:"labels"`
	Name        types.String                                           `tfsdk:"name"`
	Properties  *NetworkSecurityGroupResourceResourceGroupPropertiesTF `tfsdk:"properties"`
}
type NetworkSecurityGroupResourceCloudProviderTF struct {
	Alias       types.String `tfsdk:"alias"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Kind        types.String `tfsdk:"kind"`
	Labels      types.Map    `tfsdk:"labels"`
	Name        types.String `tfsdk:"name"`
	Properties  types.Map    `tfsdk:"properties"`
}

type NetworkSecurityGroupResourceReferencesTF struct {
	Cloudprovider *NetworkSecurityGroupResourceCloudProviderTF `tfsdk:"cloudprovider"`
	Tenant        *NetworkSecurityGroupResourceTenantTF        `tfsdk:"tenant"`
	Subscription  *NetworkSecurityGroupResourceSubscriptionTF  `tfsdk:"subscription"`
	ResourceGroup *NetworkSecurityGroupResourceResourceGroupTF `tfsdk:"resourcegroup"`
}

type NetworkSecurityGroupResourcePropertiesSettingsTF struct {
	Status types.String `tfsdk:"status"`
}

type NetworkSecurityGroupResourcePropertiesTF struct {
	// References *NetworkSecurityGroupResourceReferencesTF         `tfsdk:"references"`
	Settings *NetworkSecurityGroupResourcePropertiesSettingsTF `tfsdk:"settings"`
}
type NetworkSecurityGroupResourceTerraformModel struct {
	Alias        types.String                              `tfsdk:"alias"`
	Dependencies types.Map                                 `tfsdk:"dependencies"`
	Description  types.String                              `tfsdk:"description"`
	Kind         types.String                              `tfsdk:"kind"`
	Labels       types.Map                                 `tfsdk:"labels"`
	Name         types.String                              `tfsdk:"name"`
	ReleaseDate  types.String                              `tfsdk:"release_date"`
	Version      types.String                              `tfsdk:"version"`
	ID           types.String                              `tfsdk:"id"`
	Status       types.String                              `tfsdk:"status"`
	Properties   *NetworkSecurityGroupResourcePropertiesTF `tfsdk:"properties"`
}

var _ resource.Resource = &networkSecurityGroupResource{}

// var _ resource.ResourceWithConfigure = &subscriptionResource{}

type networkSecurityGroupResource struct {
	provider structureDeployProvider
}

func NetworkSecurityGroupResource() resource.Resource {
	return &networkSecurityGroupResource{}
}

func (e *networkSecurityGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_security_group"
}

func (e *networkSecurityGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				Required: true,
			},
			"dependencies": schema.MapAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"kind": schema.StringAttribute{
				Required: true,
			},
			"labels": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"release_date": schema.StringAttribute{
				Optional: true,
			},
			"version": schema.StringAttribute{
				Optional: true,
			},
			"status": schema.StringAttribute{
				Required: true,
			},
			"properties": schema.ObjectAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				AttributeTypes: map[string]attr.Type{
					"settings": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"status": types.StringType,
						},
					},
					// "references": types.ObjectType{
					// 	AttrTypes: map[string]attr.Type{
					// 		"cloudprovider": types.ObjectType{
					// 			AttrTypes: map[string]attr.Type{
					// 				"alias":       types.StringType,
					// 				"description": types.StringType,
					// 				"id":          types.StringType,
					// 				"kind":        types.StringType,
					// 				"labels": types.MapType{
					// 					ElemType: types.StringType,
					// 				},
					// 				"name": types.StringType,
					// 				"properties": types.MapType{
					// 					ElemType: types.StringType,
					// 				},
					// 			},
					// 		},
					// 		"tenant": types.ObjectType{
					// 			AttrTypes: map[string]attr.Type{
					// 				"alias":       types.StringType,
					// 				"description": types.StringType,
					// 				"id":          types.StringType,
					// 				"kind":        types.StringType,
					// 				"labels": types.MapType{
					// 					ElemType: types.StringType,
					// 				},
					// 				"name": types.StringType,
					// 				"properties": types.ObjectType{
					// 					AttrTypes: map[string]attr.Type{
					// 						"settings": types.ObjectType{
					// 							AttrTypes: map[string]attr.Type{
					// 								"currency":  types.StringType,
					// 								"tenant_id": types.StringType,
					// 							},
					// 						},
					// 					},
					// 				},
					// 			},
					// 		},
					// 		"subscription": types.ObjectType{
					// 			AttrTypes: map[string]attr.Type{
					// 				"alias":       types.StringType,
					// 				"description": types.StringType,
					// 				"id":          types.StringType,
					// 				"kind":        types.StringType,
					// 				"labels": types.MapType{
					// 					ElemType: types.StringType,
					// 				},
					// 				"name": types.StringType,
					// 				"properties": types.ObjectType{
					// 					AttrTypes: map[string]attr.Type{
					// 						"settings": types.ObjectType{
					// 							AttrTypes: map[string]attr.Type{
					// 								"subscription_id": types.StringType,
					// 							},
					// 						},
					// 					},
					// 				},
					// 			},
					// 		},
					// 		"resourcegroup": types.ObjectType{
					// 			AttrTypes: map[string]attr.Type{
					// 				"alias":       types.StringType,
					// 				"description": types.StringType,
					// 				"id":          types.StringType,
					// 				"kind":        types.StringType,
					// 				"labels": types.MapType{
					// 					ElemType: types.StringType,
					// 				},
					// 				"name": types.StringType,
					// 				"properties": types.ObjectType{
					// 					AttrTypes: map[string]attr.Type{
					// 						"settings": types.ObjectType{
					// 							AttrTypes: map[string]attr.Type{
					// 								"continent": types.StringType,
					// 							},
					// 						},
					// 					},
					// 				},
					// 			},
					// 		},
					// 	},
					// },
				},
			},
		},
	}
}

func (e *networkSecurityGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfstate NetworkSecurityGroupResourceTerraformModel
	// req.Config.Get(ctx, &plan)
	diags := req.Config.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	dependencies := make(map[string]interface{})
	for key, value := range tfstate.Dependencies.Elements() {
		dependencies[key] = strings.Replace(string(value.String()), "\"", "", -1)
	}
	labels := make(map[string]interface{})
	for key, value := range tfstate.Labels.Elements() {
		labels[key] = strings.Replace(string(value.String()), "\"", "", -1)
	}

	requestBody := &NetworkSecurityGroupResourceAPIModelRequestBody{
		Alias:        tfstate.Alias.ValueString(),
		Dependencies: dependencies,
		Description:  tfstate.Description.ValueString(),
		Kind:         tfstate.Kind.ValueString(),
		Labels:       labels,
		Name:         tfstate.Name.ValueString(),
		ReleaseDate:  tfstate.ReleaseDate.ValueString(),
		Version:      tfstate.Version.ValueString(),
		Properties: NetworkSecurityGroupResourcePropertiesRequestBody{
			Settings: NetworkSecurityGroupResourcePropertiesSettings{
				Status: tfstate.Status.ValueString(),
			},
		},
	}

	// Create resource using 3rd party API.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable marshall JSON Body",
			"An unexpected error occurred while marshall JSON Body. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	fmt.Println(string(requestBodyBytes))
	requestData, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://localhost:8443/api/v1/nsg", bytes.NewReader(requestBodyBytes))

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Resource",
			"An unexpected error occurred while send POST request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	client := &http.Client{}
	responseData, err := client.Do(requestData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Resource",
			"An unexpected error occurred while send POST request (client.Do). "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	defer func() {
		if responseData != nil && responseData.Body != nil {
			if cerr := responseData.Body.Close(); cerr != nil {
				resp.Diagnostics.AddError(
					"Error while Close response body",
					"An unexpected error occurred while Close response Body. "+
						"Please report this issue to the provider developers.\n\n"+
						"JSON Error: "+err.Error(),
				)
				return
			}
		}
	}()
	body, err := ioutil.ReadAll(responseData.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read response body",
			"An unexpected error occurred while Read response body. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	if responseData.StatusCode < 200 || responseData.StatusCode >= 300 {
		resp.Diagnostics.AddError(
			"Error from API Server",
			"An unexpected status code return from API Server "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+string(body),
		)
		return
	}
	var responseBody NetworkSecurityGroupResourceAPIModelResponseBody
	err = json.Unmarshal(body, &responseBody)

	// err = json.NewDecoder(responseData.Body).Decode(&responseBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error while unmarshall de responseBody",
			"An unexpected error occurred while unmarshall de responseBody "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	tfstate = NetworkSecurityGroupResourceTerraformModel{
		ID:           types.StringValue(responseBody.ID),
		Alias:        types.StringValue(responseBody.Alias),
		Name:         types.StringValue(responseBody.Name),
		Description:  types.StringValue(responseBody.Description),
		Kind:         types.StringValue(responseBody.Kind),
		Status:       tfstate.Status,
		Version:      tfstate.Version,
		Dependencies: tfstate.Dependencies,
		ReleaseDate:  tfstate.ReleaseDate,
		Properties: &NetworkSecurityGroupResourcePropertiesTF{
			Settings: &NetworkSecurityGroupResourcePropertiesSettingsTF{
				Status: types.StringValue(responseBody.Properties.Settings.Status),
			},
			// References: &NetworkSecurityGroupResourceReferencesTF{
			// 	Cloudprovider: &NetworkSecurityGroupResourceCloudProviderTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
			// 	},
			// 	Tenant: &NetworkSecurityGroupResourceTenantTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceTenantPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceTenantPropertiesSettingsTF{
			// 				Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
			// 				TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
			// 			},
			// 		},
			// 	},
			// 	Subscription: &NetworkSecurityGroupResourceSubscriptionTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceSubscriptionPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceSubscriptionPropertiesSettingsTF{
			// 				SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
			// 			},
			// 		},
			// 	},
			// 	ResourceGroup: &NetworkSecurityGroupResourceResourceGroupTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceResourceGroupPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceResourceGroupPropertiesSettingsTF{
			// 				Continent: types.StringValue(responseBody.Properties.References.ResourceGroup.Properties.Settings.Continent),
			// 			},
			// 		},
			// 	},
			// },
		},
	}
	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	// tfstate.Properties.References.Cloudprovider.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Labels)
	// tfstate.Properties.References.Cloudprovider.Properties, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Properties)
	// tfstate.Properties.References.Tenant.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Tenant.Labels)
	// tfstate.Properties.References.Subscription.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Subscription.Labels)
	// tfstate.Properties.References.ResourceGroup.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.ResourceGroup.Labels)
	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
}

func (e *networkSecurityGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfstate NetworkSecurityGroupResourceTerraformModel

	diags := req.State.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Read resource using 3rd party API.
	var requestBody NetworkSecurityGroupResourceAPIModelRequestBody
	requestBody.Name = strings.Replace(tfstate.Name.ValueString(), " ", "%20", -1)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestData, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://localhost:8443/api/v1/nsg/%s", requestBody.Name), nil)
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
	var responseBody NetworkSecurityGroupResourceAPIModelResponseBody
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
	tfstate = NetworkSecurityGroupResourceTerraformModel{
		ID:           types.StringValue(responseBody.ID),
		Alias:        types.StringValue(responseBody.Alias),
		Name:         types.StringValue(responseBody.Name),
		Description:  types.StringValue(responseBody.Description),
		Kind:         types.StringValue(responseBody.Kind),
		Status:       tfstate.Status,
		Version:      tfstate.Version,
		Dependencies: tfstate.Dependencies,
		ReleaseDate:  tfstate.ReleaseDate,
		Properties: &NetworkSecurityGroupResourcePropertiesTF{
			Settings: &NetworkSecurityGroupResourcePropertiesSettingsTF{
				Status: types.StringValue(responseBody.Properties.Settings.Status),
			},
			// References: &NetworkSecurityGroupResourceReferencesTF{
			// 	Cloudprovider: &NetworkSecurityGroupResourceCloudProviderTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
			// 	},
			// 	Tenant: &NetworkSecurityGroupResourceTenantTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceTenantPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceTenantPropertiesSettingsTF{
			// 				Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
			// 				TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
			// 			},
			// 		},
			// 	},
			// 	Subscription: &NetworkSecurityGroupResourceSubscriptionTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceSubscriptionPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceSubscriptionPropertiesSettingsTF{
			// 				SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
			// 			},
			// 		},
			// 	},
			// 	ResourceGroup: &NetworkSecurityGroupResourceResourceGroupTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceResourceGroupPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceResourceGroupPropertiesSettingsTF{
			// 				Continent: types.StringValue(responseBody.Properties.References.ResourceGroup.Properties.Settings.Continent),
			// 			},
			// 		},
			// 	},
			// },
		},
	}
	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	// tfstate.Properties.References.Cloudprovider.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Labels)
	// tfstate.Properties.References.Cloudprovider.Properties, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Properties)
	// tfstate.Properties.References.Tenant.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Tenant.Labels)
	// tfstate.Properties.References.Subscription.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Subscription.Labels)
	// tfstate.Properties.References.ResourceGroup.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.ResourceGroup.Labels)

	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (e *networkSecurityGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var tfstate NetworkSecurityGroupResourceTerraformModel

	diags := req.Plan.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Update resource using 3rd party API.
	dependencies := make(map[string]interface{})
	for key, value := range tfstate.Dependencies.Elements() {
		dependencies[key] = strings.Replace(string(value.String()), "\"", "", -1)
	}
	labels := make(map[string]interface{})
	for key, value := range tfstate.Labels.Elements() {
		labels[key] = strings.Replace(string(value.String()), "\"", "", -1)
	}
	// d := json.Unmarshal([]byte(config.Dependencies.String()), &)

	requestBody := &NetworkSecurityGroupResourceAPIModelRequestBody{
		ID:           tfstate.ID.ValueString(),
		Alias:        tfstate.Alias.ValueString(),
		Dependencies: dependencies,
		Description:  tfstate.Description.ValueString(),
		Kind:         tfstate.Kind.ValueString(),
		Labels:       labels,
		Name:         tfstate.Name.ValueString(),
		ReleaseDate:  tfstate.ReleaseDate.ValueString(),
		Version:      tfstate.Version.ValueString(),
		Properties: NetworkSecurityGroupResourcePropertiesRequestBody{
			Settings: NetworkSecurityGroupResourcePropertiesSettings{
				Status: tfstate.Status.ValueString(),
			},
		},
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error occurred while marshall JSON Body",
			"An unexpected error occurred while marshall JSON Body. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	requestData, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/nsg/", "https://localhost:8443"), bytes.NewReader(requestBodyBytes))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Send PUT Request",
			"An unexpected error occurred while send PUT request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	requestData.Header.Add("Content-Type", "application/json")
	responseData, err := client.Do(requestData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error while Making Request",
			"An unexpected error occurred while making request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	defer func() {
		if responseData != nil && responseData.Body != nil {
			if cerr := responseData.Body.Close(); cerr != nil {
				resp.Diagnostics.AddError(
					"Unable to Close Response Body",
					"An unexpected error occurred while Close response Body. "+
						"Please report this issue to the provider developers.\n\n"+
						"JSON Error: "+err.Error(),
				)
				return
			}
		}
	}()
	body, err := ioutil.ReadAll(responseData.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Reponse Body",
			"An unexpected error occurred while Read response body. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	if responseData.StatusCode < 200 || responseData.StatusCode >= 300 {
		resp.Diagnostics.AddError(
			"Error From API Server",
			"An unexpected status code return from API "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+string(body),
		)
		return
	}
	var responseBody NetworkSecurityGroupResourceAPIModelResponseBody
	err = json.Unmarshal(body, &responseBody)
	// err = json.NewDecoder(responseData.Body).Decode(&responseBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error while decode Response Body",
			"An unexpected error occurred while decode responseBody "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	tfstate = NetworkSecurityGroupResourceTerraformModel{
		ID:           types.StringValue(responseBody.ID),
		Alias:        types.StringValue(responseBody.Alias),
		Name:         types.StringValue(responseBody.Name),
		Description:  types.StringValue(responseBody.Description),
		Kind:         types.StringValue(responseBody.Kind),
		Status:       tfstate.Status,
		Version:      tfstate.Version,
		Dependencies: tfstate.Dependencies,
		ReleaseDate:  tfstate.ReleaseDate,
		Properties: &NetworkSecurityGroupResourcePropertiesTF{
			Settings: &NetworkSecurityGroupResourcePropertiesSettingsTF{
				Status: types.StringValue(responseBody.Properties.Settings.Status),
			},
			// References: &NetworkSecurityGroupResourceReferencesTF{
			// 	Cloudprovider: &NetworkSecurityGroupResourceCloudProviderTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Cloudprovider.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Cloudprovider.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Cloudprovider.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Cloudprovider.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Cloudprovider.Name),
			// 	},
			// 	Tenant: &NetworkSecurityGroupResourceTenantTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceTenantPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceTenantPropertiesSettingsTF{
			// 				Currency: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.Currency),
			// 				TenantID: types.StringValue(responseBody.Properties.References.Tenant.Properties.Settings.TenantID),
			// 			},
			// 		},
			// 	},
			// 	Subscription: &NetworkSecurityGroupResourceSubscriptionTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceSubscriptionPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceSubscriptionPropertiesSettingsTF{
			// 				SubscriptionID: types.StringValue(responseBody.Properties.References.Subscription.Properties.Settings.SubscriptionID),
			// 			},
			// 		},
			// 	},
			// 	ResourceGroup: &NetworkSecurityGroupResourceResourceGroupTF{
			// 		Alias:       types.StringValue(responseBody.Properties.References.Tenant.Alias),
			// 		Description: types.StringValue(responseBody.Properties.References.Tenant.Description),
			// 		ID:          types.StringValue(responseBody.Properties.References.Tenant.ID),
			// 		Kind:        types.StringValue(responseBody.Properties.References.Tenant.Kind),
			// 		Name:        types.StringValue(responseBody.Properties.References.Tenant.Name),
			// 		Properties: &NetworkSecurityGroupResourceResourceGroupPropertiesTF{
			// 			Settings: &NetworkSecurityGroupResourceResourceGroupPropertiesSettingsTF{
			// 				Continent: types.StringValue(responseBody.Properties.References.ResourceGroup.Properties.Settings.Continent),
			// 			},
			// 		},
			// 	},
			// },
		},
	}
	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	// tfstate.Properties.References.Cloudprovider.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Labels)
	// tfstate.Properties.References.Cloudprovider.Properties, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Properties)
	// tfstate.Properties.References.Tenant.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Tenant.Labels)
	// tfstate.Properties.References.Subscription.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Subscription.Labels)
	// tfstate.Properties.References.ResourceGroup.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.ResourceGroup.Labels)
	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
}

func (e *networkSecurityGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkSecurityGroupResourceTerraformModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	var requestBody NetworkSecurityGroupResourceAPIModelRequestBody
	requestBody.ID = state.ID.ValueString()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	requestData, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/nsg/%s", "https://localhost:8443", requestBody.ID), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while send DELETE request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
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
			"Error from API Server",
			"An unexpected status code return from API "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+string(body),
		)
		return
	}
	defer responseData.Body.Close()
	// Delete resource using 3rd party API.
}

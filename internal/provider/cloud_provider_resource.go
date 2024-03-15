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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudProviderResourceAPIModelRequestBody struct {
	Alias       string                 `json:"alias"`
	Description string                 `json:"description"`
	Kind        string                 `json:"kind"`
	Labels      map[string]interface{} `json:"labels"`
	Name        string                 `json:"name"`
	ReleaseDate string                 `json:"release_date"`
	Version     string                 `json:"version"`
	ID          string                 `json:"id"`
}
type CloudProviderResourceAPIModelResponseBody struct {
	Alias       string                 `json:"alias"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Labels      map[string]interface{} `json:"labels"`
	ID          string                 `json:"id"`
	// Kind        string                 `json:"kind"`
}

type CloudProviderResourceTerraformModel struct {
	Alias       types.String `tfsdk:"alias"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Kind        types.String `tfsdk:"kind"`
	Labels      types.Map    `tfsdk:"labels"`
	ReleaseDate types.String `tfsdk:"release_date"`
	Version     types.String `tfsdk:"version"`
	ID          types.String `tfsdk:"id"`
}

var _ resource.Resource = &cloudProviderResource{}

// var _ resource.ResourceWithConfigure = &subscriptionResource{}

type cloudProviderResource struct {
	provider structureDeployProvider
}

func CloudProviderResource() resource.Resource {
	return &cloudProviderResource{}
}

func (e *cloudProviderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_provider"
}

func (e *cloudProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				Required: true,
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
		},
	}
}

func (e *cloudProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfstate CloudProviderResourceTerraformModel
	// req.Config.Get(ctx, &plan)
	diags := req.Config.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	labels := make(map[string]interface{})
	for key, value := range tfstate.Labels.Elements() {
		labels[key] = strings.Replace(string(value.String()), "\"", "", -1)
	}

	requestBody := &CloudProviderResourceAPIModelRequestBody{
		Alias:       tfstate.Alias.ValueString(),
		Description: tfstate.Description.ValueString(),
		Kind:        tfstate.Kind.ValueString(),
		Labels:      labels,
		Name:        tfstate.Name.ValueString(),
		ReleaseDate: tfstate.ReleaseDate.ValueString(),
		Version:     tfstate.Version.ValueString(),
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
	requestData, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://localhost:8443/api/v1/cloudProvider", bytes.NewReader(requestBodyBytes))

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
	var responseBody CloudProviderResourceAPIModelResponseBody
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
	tfstate = CloudProviderResourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        tfstate.Kind,
		Version:     tfstate.Version,
		ReleaseDate: tfstate.ReleaseDate,
	}
	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
}

func (e *cloudProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfstate CloudProviderResourceTerraformModel

	diags := req.State.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Read resource using 3rd party API.
	var requestBody CloudProviderResourceAPIModelRequestBody
	requestBody.Name = strings.Replace(tfstate.Name.ValueString(), " ", "%20", -1)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	requestData, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://localhost:8443/api/v1/cloudProvider/%s", requestBody.Name), nil)
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
	var responseBody CloudProviderResourceAPIModelResponseBody
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
	tfstate = CloudProviderResourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        tfstate.Kind,
		Version:     tfstate.Version,
		ReleaseDate: tfstate.ReleaseDate,
	}
	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	// tfstate.Properties.References.Cloudprovider.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Labels)
	// tfstate.Properties.References.Cloudprovider.Properties, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Properties)
	// tfstate.Properties.References.Tenant.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Tenant.Labels)

	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (e *cloudProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var tfstate CloudProviderResourceTerraformModel

	diags := req.Plan.Get(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Update resource using 3rd party API.
	labels := make(map[string]interface{})
	for key, value := range tfstate.Labels.Elements() {
		labels[key] = strings.Replace(string(value.String()), "\"", "", -1)
	}
	// d := json.Unmarshal([]byte(config.Dependencies.String()), &)

	requestBody := &CloudProviderResourceAPIModelRequestBody{
		ID:          tfstate.ID.ValueString(),
		Alias:       tfstate.Alias.ValueString(),
		Description: tfstate.Description.ValueString(),
		Kind:        tfstate.Kind.ValueString(),
		Labels:      labels,
		Name:        tfstate.Name.ValueString(),
		ReleaseDate: tfstate.ReleaseDate.ValueString(),
		Version:     tfstate.Version.ValueString(),
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
	requestData, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/cloudProvider/", "https://localhost:8443"), bytes.NewReader(requestBodyBytes))
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
	var responseBody CloudProviderResourceAPIModelResponseBody
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
	tfstate = CloudProviderResourceTerraformModel{
		ID:          types.StringValue(responseBody.ID),
		Alias:       types.StringValue(responseBody.Alias),
		Name:        types.StringValue(responseBody.Name),
		Description: types.StringValue(responseBody.Description),
		Kind:        tfstate.Kind,
		Version:     tfstate.Version,
		ReleaseDate: tfstate.ReleaseDate,
	}
	tfstate.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Labels)
	// tfstate.Properties.References.Cloudprovider.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Labels)
	// tfstate.Properties.References.Cloudprovider.Properties, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Cloudprovider.Properties)
	// tfstate.Properties.References.Tenant.Labels, diags = types.MapValueFrom(ctx, types.StringType, responseBody.Properties.References.Tenant.Labels)
	diags = resp.State.Set(ctx, &tfstate)
	resp.Diagnostics.Append(diags...)
}

func (e *cloudProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CloudProviderResourceTerraformModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	var requestBody CloudProviderResourceAPIModelRequestBody
	requestBody.ID = state.ID.ValueString()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	requestData, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/cloudProvider/%s", "https://localhost:8443", requestBody.ID), nil)
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

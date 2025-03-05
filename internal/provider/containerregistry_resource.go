package provider

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &ContainerregistryResource{}

func NewContainerregistryResource() resource.Resource {
	return &ContainerregistryResource{}
}

// ContainerregistryResource defines the resource implementation.
type ContainerregistryResource struct {
}

// ExampleResourceModel describes the resource data model.
type ContainerregistryResourceModel struct {
	ImageTarball     types.String `tfsdk:"image_tarball"`
	ImageTarballHash types.String `tfsdk:"image_tarball_hash"`
	RemoteTag        types.String `tfsdk:"remote_tag"`
	Id               types.String `tfsdk:"id"`
}

func (r *ContainerregistryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (r *ContainerregistryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Sample resource in the Terraform provider containerregistry.",

		Attributes: map[string]schema.Attribute{
			"image_tarball": schema.StringAttribute{
				MarkdownDescription: "Image tarball thing.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"image_tarball_hash": schema.StringAttribute{
				MarkdownDescription: "Hash of the image tarball.",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"remote_tag": schema.StringAttribute{
				MarkdownDescription: "The tag to save the image to.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of this resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ContainerregistryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContainerregistryResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tag, err := name.NewTag(data.RemoteTag.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}
	img, err := tarball.ImageFromPath(data.ImageTarball.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}
	err = remote.Write(tag, img, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}

	idFromAPI := data.RemoteTag.ValueString()
	data.Id = types.StringValue(idFromAPI)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContainerregistryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContainerregistryResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tag, err := name.NewTag(data.RemoteTag.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}
	_, err = remote.Head(tag, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContainerregistryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ContainerregistryResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContainerregistryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContainerregistryResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tag, err := name.NewTag(data.RemoteTag.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}
	err = remote.Delete(tag, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		resp.Diagnostics.AddError("It break", err.Error())
		return
	}
}

package provider

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceContainerregistry() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider containerregistry.",

		CreateContext: resourceContainerRegistryCreate,
		ReadContext:   resourceContainerRegistryRead,
		// UpdateContext: resourceContainerRegistryUpdate,
		DeleteContext: resourceContainerRegistryDelete,

		Schema: map[string]*schema.Schema{
			"image_tarball": {
				// This description is used by the documentation generator and the language server.
				Description: "Image tarball thing.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"image_tarball_hash": {
				// This description is used by the documentation generator and the language server.
				Description: "Hash of the image tarball.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"remote_tag": {
				// This description is used by the documentation generator and the language server.
				Description: "The tag to save the image to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceContainerRegistryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	tag, err := name.NewTag(d.Get("remote_tag").(string))
	if err != nil {
		return diag.Errorf(err.Error())
	}
	img, err := tarball.ImageFromPath(d.Get("image_tarball").(string), nil)
	if err != nil {
		return diag.Errorf(err.Error())
	}
	err = remote.Write(tag, img, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return diag.Errorf(err.Error())
	}

	idFromAPI := d.Get("remote_tag").(string)
	d.SetId(idFromAPI)

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	return nil
}

func resourceContainerRegistryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	tag, err := name.NewTag(d.Get("remote_tag").(string))
	if err != nil {
		return diag.Errorf(err.Error())
	}
	_, err = remote.Head(tag, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return diag.Errorf(err.Error())
	}

	return nil
}

// func resourceContainerRegistryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
// 	// use the meta value to retrieve your client from the provider configure method
// 	// client := meta.(*apiClient)

// 	return diag.Errorf("not implemented")
// }

func resourceContainerRegistryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	tag, err := name.NewTag(d.Get("remote_tag").(string))
	if err != nil {
		return diag.Errorf(err.Error())
	}
	err = remote.Delete(tag, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return diag.Errorf(err.Error())
	}

	return nil
}

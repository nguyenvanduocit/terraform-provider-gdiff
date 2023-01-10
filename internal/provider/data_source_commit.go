package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-gdiff/internal/client"
)

func dataSourceGdiff() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "This data source provides the last commit of a file or folder",

		ReadContext: dataSourceGdiffRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Description: "Absolute path to the file,directory.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hash": {
				Description: "The hash of the last commit",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceGdiffRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	apiClient := meta.(*client.ApiClient)

	path := d.Get("path").(string)

	commit, err := apiClient.GetLastCommit(path)
	if err != nil {
		return diag.FromErr(err)
	}

	if commit != nil {
		d.SetId(commit.String())
		d.Set("hash", commit.String())
		return nil
	}

	// The commit does not contain the path, so we can't use it as ID, try to get th existing ID
	if hash, ok := d.GetOk("hash"); ok {
		d.SetId(hash.(string))
		d.Set("hash", hash.(string))
	}

	return nil
}

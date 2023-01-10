package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGdiff() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider gdiff.",

		ReadContext: dataSourceGdiffRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Description: "Absolute path to the file,directory.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceGdiffRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}

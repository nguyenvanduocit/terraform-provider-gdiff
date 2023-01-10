package provider

import (
	"context"
	"github.com/hashicorp/terraform-provider-gdiff/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"git_path": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Absolute path to the git folder",
					DefaultFunc: schema.EnvDefaultFunc("GIT_PATH", nil),
				},
				"diff_mode": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "DiffMode of the git repository. e.g. 'tag' or 'commit', 'stage', 'dirty'",
					DefaultFunc: schema.EnvDefaultFunc("GIT_DIFF_MODE", "tag"),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"gdiff_commit": dataSourceGdiff(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		return &client.ApiClient{
			DiffMode:     client.DiffMode(d.Get("diff_mode").(string)),
			GitPath:      d.Get("git_path").(string),
			ResourceData: d,
		}, nil
	}
}

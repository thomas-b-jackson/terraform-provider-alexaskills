package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scg/va/ask_client"
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
			DataSourcesMap: map[string]*schema.Resource{
				"alexaskills_skill_resource":             dataSourceSkills(),
				"alexaskills_interaction_model_resource": dataSourceInteractionModel(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"alexaskills_skill_resource":             resourceSkills(),
				"alexaskills_interaction_model_resource": resourceInteractionModel(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {

		var diags diag.Diagnostics
		askClient, err := ask_client.NewClient()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create ask client",
				Detail:   "Unable to create ask client",
			})
			return nil, diags
		}

		return askClient, diags
	}
}

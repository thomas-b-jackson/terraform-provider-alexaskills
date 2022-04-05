package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scg/va/smapi_client"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func Provider(version string) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SMAPI_TOKEN", nil),
			},
			"vendor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"alexaskills_skill_resource":             dataSourceSkills(),
			"alexaskills_interaction_model_resource": dataSourceInteractionModel(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"alexaskills_skill_resource":             resourceSkills(),
			"alexaskills_interaction_model_resource": resourceInteractionModel(),
		},
		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	token := d.Get("token").(string)
	vendorId := d.Get("vendor_id").(string)

	smapiClient, err := smapi_client.NewClient(token, vendorId)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create smapi client",
			Detail:   "Unable to create smapi client",
		})
		return nil, diags
	}

	return smapiClient, diags
}

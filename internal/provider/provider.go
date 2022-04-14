package provider

import (
	"context"
	"fmt"
	"log"

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
			"vendor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lwa_client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LWA_CLIENT_ID", nil),
			},
			"lwa_client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LWA_CLIENT_SECRET", nil),
			},
			"lwa_refresh_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LWA_REFRESH_TOKEN", nil),
			},
			"lwa_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LWA_ACCESS_TOKEN", nil),
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

	vendorId := d.Get("vendor_id").(string)
	lwa_access_token := d.Get("lwa_token").(string)

	log.Printf("[DEBUG] token is %s", lwa_access_token)

	if lwa_access_token == "" {

		lwa_client_id := d.Get("lwa_client_id").(string)
		lwa_client_secret := d.Get("lwa_client_secret").(string)
		lwa_refresh_token := d.Get("lwa_refresh_token").(string)

		// create an lwa access token
		generated_access_token, err := smapi_client.GetLwaToken(lwa_client_id,
			lwa_client_secret,
			lwa_refresh_token)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create LWA access token",
				Detail:   fmt.Sprintf("Unable to create LWA access token, err: %s", err),
			})
			return nil, diags
		}

		lwa_access_token = generated_access_token
	}

	smapiClient, err := smapi_client.NewClient(lwa_access_token, vendorId)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create smapi client",
			Detail:   fmt.Sprintf("Unable to create smapi client, err: %s", err),
		})
		return nil, diags
	}

	return smapiClient, diags
}

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scg/va/smapi_client"
)

func dataSourceSkills() *schema.Resource {

	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "a data source returning details on a alexa skill",
		ReadContext: dataSourceSkillsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the skill.",
			},
			"manifest": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: skillSchemaMap,
				},
			},
		},
	}
}

func dataSourceSkillsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	skillID := d.Get("id").(string)

	smapiClient := meta.(*smapi_client.SMAPIClient)

	skillManifest, err := smapiClient.GetSkillManifest(skillID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get requested skill manifest",
			Detail:   fmt.Sprintf("Unable to get requested skill manifest, err: %s", err),
		})
		return diags
	}

	flattenedSkillManifest := flattenSkillManifest(skillManifest)

	if err := d.Set("manifest", flattenedSkillManifest); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set requested skill manifest",
			Detail:   fmt.Sprintf("Unable to set requested skill manifest, err: %s", err),
		})
		return diags
	}

	d.SetId(skillID)

	return diags
}

var skillSchemaMap = map[string]*schema.Schema{
	"manifest_version": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"publishing_information": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem:     lexPubInfoResource,
	},
	"apis": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     lexApiResource,
	},
}

var lexPubInfoResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"locales": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"en_us": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"summary": {
									Type:     schema.TypeString,
									Required: true,
								},
								"example_phrases": {
									Type:     schema.TypeList,
									Required: true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"name": {
									Type:     schema.TypeString,
									Required: true,
								},
								"description": {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		"is_available_worldwide": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"testing_instructions": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"category": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"distribution_countries": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	},
}

var lexApiResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"custom": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"endpoint": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"uri": {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					"interfaces": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	},
}

func flattenSkillManifest(skillManifest smapi_client.SkillManifest) (flattened []map[string]interface{}) {

	flattened = []map[string]interface{}{
		{
			"manifest_version":       skillManifest.ManifestVersion,
			"publishing_information": flattenPublishingInformation(skillManifest.PublishingInformation),
			"apis":                   flattenApis(skillManifest.Apis),
		},
	}

	return
}

func flattenPublishingInformation(publishingInformation smapi_client.PublishingInformation) (flattened []map[string]interface{}) {

	flattened = []map[string]interface{}{
		{
			"locales":                flattenLocales(publishingInformation.Locales),
			"category":               publishingInformation.Category,
			"distribution_countries": publishingInformation.DistributionCountries,
			"is_available_worldwide": publishingInformation.IsAvailableWorldwide,
			"testing_instructions":   publishingInformation.TestingInstructions,
		},
	}

	return
}

func flattenLocales(locales smapi_client.Locales) (flattened []map[string]interface{}) {

	flattened = []map[string]interface{}{{
		"en_us": []map[string]interface{}{{
			"name":            locales.EnglishUS.Name,
			"summary":         locales.EnglishUS.Summary,
			"description":     locales.EnglishUS.Description,
			"example_phrases": flattenExamplePhrases(locales.EnglishUS.ExamplePhrases),
		}},
	}}

	return
}

func flattenExamplePhrases(phrasesIn []string) (phrases []interface{}) {
	phrases = make([]interface{}, len(phrasesIn))

	for idx, phrase := range phrasesIn {
		phrases[idx] = phrase
	}
	return
}

func flattenApis(apis smapi_client.Apis) (flattened []map[string]interface{}) {

	flattened = []map[string]interface{}{{
		"custom": []map[string]interface{}{{
			"endpoint": []map[string]interface{}{{
				"uri": apis.Custom.Endpoint.Uri,
			}},
			"interfaces": apis.Custom.Interfaces,
		}},
	}}

	return
}

func ExpandSkillManifest(in []interface{}) *smapi_client.SkillManifest {

	m := in[0].(map[string]interface{})

	manifest := smapi_client.SkillManifest{}

	if v, ok := m["publishing_information"].([]interface{}); ok {
		manifest.PublishingInformation = *ExpandPublishingInformation(v)
	}

	if v, ok := m["apis"].([]interface{}); ok {
		manifest.Apis = *ExpandApis(v)
	}

	if v, ok := m["manifest_version"]; ok {
		manifest.ManifestVersion = v.(string)
	}

	return &manifest
}

func ExpandPublishingInformation(in []interface{}) *smapi_client.PublishingInformation {

	pubInfo := &smapi_client.PublishingInformation{}

	m := in[0].(map[string]interface{})

	if v, ok := m["locales"].([]interface{}); ok && len(v) > 0 {
		pubInfo.Locales = *expandLocales(v)
	}

	if v, ok := m["is_available_worldwide"].(bool); ok {
		pubInfo.IsAvailableWorldwide = v
	}

	if v, ok := m["testing_instructions"].(string); ok {
		pubInfo.TestingInstructions = v
	}

	if v, ok := m["category"].(string); ok {
		pubInfo.Category = v
	}

	return pubInfo
}

func ExpandApis(in []interface{}) *smapi_client.Apis {

	apis := &smapi_client.Apis{}

	m := in[0].(map[string]interface{})

	if v, ok := m["custom"].([]interface{}); ok && len(v) > 0 {
		customApi := expandCustomApi(v)
		apis.Custom = *customApi
	}
	return apis
}

func expandCustomApi(in []interface{}) *smapi_client.CustomApi {

	customApi := &smapi_client.CustomApi{}

	m := in[0].(map[string]interface{})

	if v, ok := m["endpoint"].([]interface{}); ok && len(v) > 0 {
		endpoint := expandEndpoint(v)
		customApi.Endpoint = *endpoint
	}
	return customApi
}

func expandEndpoint(in []interface{}) *smapi_client.Endpoint {

	endpoint := &smapi_client.Endpoint{}

	m := in[0].(map[string]interface{})

	if v, ok := m["uri"].(string); ok {
		endpoint.Uri = v
	}

	return endpoint
}

func expandLocales(in []interface{}) *smapi_client.Locales {

	locales := &smapi_client.Locales{}

	m := in[0].(map[string]interface{})

	if v, ok := m["en_us"].([]interface{}); ok && len(v) > 0 {
		enUS := explandEnUSLocal(v)
		locales.EnglishUS = *enUS
	}

	return locales
}

func explandEnUSLocal(in []interface{}) *smapi_client.EnglishUSLocal {

	enUS := &smapi_client.EnglishUSLocal{}

	m := in[0].(map[string]interface{})

	if v, ok := m["name"].(string); ok {
		enUS.Name = v
	}

	if v, ok := m["summary"].(string); ok {
		enUS.Summary = v
	}

	if v, ok := m["description"].(string); ok {
		enUS.Description = v
	}

	if v, ok := m["example_phrases"].([]interface{}); ok {
		enUS.ExamplePhrases = make([]string, len(v))

		for idx, phrase := range v {
			enUS.ExamplePhrases[idx] = phrase.(string)
		}
	}

	return enUS
}

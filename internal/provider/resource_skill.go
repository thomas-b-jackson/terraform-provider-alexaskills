package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scg/va/smapi_client"
)

func resourceSkills() *schema.Resource {
	return &schema.Resource{
		Description: "Alex skill resource",

		CreateContext: resourceSkillsCreate,
		ReadContext:   resourceSkillsRead,
		UpdateContext: resourceSkillsUpdate,
		DeleteContext: resourceSkillsDelete,

		Schema: map[string]*schema.Schema{
			"manifest": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: skillSchemaMap,
				},
			},
		},
	}
}

func resourceSkillsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	manifest := ExpandSkillManifest(d.Get("manifest").([]interface{}))

	smapiClient := meta.(*smapi_client.SMAPIClient)

	skillId, err := smapiClient.CreateSkill(*manifest)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create requested skill",
			Detail:   fmt.Sprintf("Unable to create requested skill, err: %s", err),
		})
		return diags
	}

	d.SetId(skillId)

	return diags
}

func resourceSkillsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	skillID := d.Id()

	smapiClient := meta.(*smapi_client.SMAPIClient)

	skillManifest, err := smapiClient.GetSkillManifest(skillID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get requested skill",
			Detail:   fmt.Sprintf("Unable to get requested skill, err: %s", err),
		})
		return diags
	}

	flattenedSkillManifest := flattenSkillManifest(skillManifest)

	if err := d.Set("manifest", flattenedSkillManifest); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set skill manifest",
			Detail:   fmt.Sprintf("Unable to set skill manifest, err: %s", err),
		})
		return diags
	}

	d.SetId(skillID)

	return diags
}

func resourceSkillsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceSkillsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	smapiClient := meta.(*smapi_client.SMAPIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	skillID := d.Id()

	err := smapiClient.DeleteSkill(skillID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete requested skill",
			Detail:   fmt.Sprintf("Unable to delete requested skill, err: %s", err),
		})
	}

	return diags
}

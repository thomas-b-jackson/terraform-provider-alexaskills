package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scg/va/smapi_client"
)

func dataSourceInteractionModel() *schema.Resource {

	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "a data source returning details on a alexa skill",
		ReadContext: dataSourceInteractionModelRead,
		Schema:      interactionModelSchema,
	}
}

var interactionModelSchema = map[string]*schema.Schema{
	"skill_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"interaction_model": {
		Type: schema.TypeList,
		// optional so same schema can be used with model data and resources
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"language_model": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"invocation_name": {
								Type:     schema.TypeString,
								Required: true,
							},
							"types": {
								Type:     schema.TypeList,
								MaxItems: 1,
								Required: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:     schema.TypeString,
											Required: true,
										},
										"values": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
									},
								},
							},
							"intents": {
								Type:     schema.TypeList,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:     schema.TypeString,
											Required: true,
										},
										"samples": {
											Type:     schema.TypeList,
											Required: true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"slots": {
											Type:     schema.TypeList,
											Optional: true,
											Elem:     slotResource,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

var slotResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

func dataSourceInteractionModelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	skillID := d.Get("skill_id").(string)

	smapiClient := meta.(*smapi_client.SMAPIClient)

	model, err := smapiClient.GetInteractionModel(skillID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get requested interaction model",
			Detail:   fmt.Sprintf("Unable to get requested interaction model, err: %s", err),
		})
		return diags
	}

	modelItems := flattenInteractionModels(model)

	if err := d.Set("interaction_model", modelItems); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read requested model",
			Detail:   fmt.Sprintf("Unable to read requested model, err: %s", err),
		})
		return diags
	}

	d.SetId(skillID)

	return diags
}

func flattenInteractionModels(model smapi_client.InteractionModelObj) (flattened []map[string][]map[string]interface{}) {

	flattened = []map[string][]map[string]interface{}{
		{
			"language_model": {{
				"invocation_name": model.InteractionModel.LanguageModel.InvocationName,
				"types":           flattenTypes(model.InteractionModel.LanguageModel.Types),
				"intents":         flattenIntents(model.InteractionModel.LanguageModel.Intents),
			}},
		},
	}

	return
}

func flattenTypes(modelTypes []smapi_client.Types) (flattened []map[string]interface{}) {

	if len(modelTypes) > 0 {
		log.Printf("[DEBUG] skills model:\n%+v\n", modelTypes)
		flattened = []map[string]interface{}{
			{
				"name":   modelTypes[0].Name,
				"values": flattenTypeValues(modelTypes[0].Values),
			},
		}
	} else {
		flattened = make([]map[string]interface{}, 0)
	}
	return
}

func flattenTypeValues(type_values []smapi_client.TypeValues) []string {

	typeValues := make([]string, len(type_values))

	for idx, type_value := range type_values {
		typeValues[idx] = type_value.Name.Value
	}

	return typeValues
}

func flattenIntents(intents []smapi_client.Intents) []map[string]interface{} {

	intents_slice := make([]map[string]interface{}, len(intents))

	for idx, intent := range intents {
		intent_map := make(map[string]interface{})
		intent_map["name"] = intent.Name

		samples := make([]interface{}, len(intent.Samples))

		for sample_idx, sample := range intent.Samples {
			samples[sample_idx] = sample
		}

		intent_map["samples"] = samples

		intent_map["slots"] = flattenSlots(intent.Slots)
		intents_slice[idx] = intent_map
	}

	return intents_slice
}

func flattenSlots(slots []smapi_client.Slot) []map[string]interface{} {

	slots_slice := make([]map[string]interface{}, len(slots))

	for idx, slot := range slots {
		slot_map := make(map[string]interface{})
		slot_map["name"] = slot.Name
		slot_map["type"] = slot.Type
		slots_slice[idx] = slot_map
	}

	return slots_slice
}

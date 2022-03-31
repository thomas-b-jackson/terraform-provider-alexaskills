package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scg/va/smapi_client"
)

func resourceInteractionModel() *schema.Resource {
	return &schema.Resource{
		Description: "Interaction model associated with an alexa skills.",

		CreateContext: resourceInteractionModelCreate,
		ReadContext:   resourceInteractionModelRead,
		UpdateContext: resourceInteractionModelUpdate,
		DeleteContext: resourceInteractionModelDelete,

		Schema: interactionModelSchema,
	}
}

func resourceInteractionModelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	interactionModel := ExpandInteractionModel(d.Get("interaction_model").([]interface{}))

	skillId := d.Get("skill_id").(string)

	smapiClient := meta.(*smapi_client.SMAPIClient)

	err := smapiClient.UpdateInteractionModel(skillId, *interactionModel)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create/update requested model",
			Detail:   fmt.Sprintf("Unable to create/update requested model, err: %s", err),
		})
		return diags
	}

	d.SetId(skillId)

	return diags
}

func resourceInteractionModelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return dataSourceInteractionModelRead(ctx, d, meta)
}

func resourceInteractionModelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// same as a model create
	return resourceInteractionModelCreate(ctx, d, meta)
}

func resourceInteractionModelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// you can delete a skill, but you cannot actually delete the interaction model
	// associated with a skill (you can only update it)
	var diags diag.Diagnostics

	return diags
}

func ExpandInteractionModel(in []interface{}) *smapi_client.InteractionModel {

	model := &smapi_client.InteractionModel{}

	m := in[0].(map[string]interface{})

	if v, ok := m["language_model"].([]interface{}); ok && len(v) > 0 {
		model.LanguageModel = *expandLanguageModel(v)
	}

	return model
}

func expandLanguageModel(in []interface{}) *smapi_client.LanguageModel {

	languageModel := &smapi_client.LanguageModel{}

	m := in[0].(map[string]interface{})

	if v, ok := m["invocation_name"].(string); ok {
		languageModel.InvocationName = v
	}

	if v, ok := m["types"].([]interface{}); ok && len(v) > 0 {
		types := expandTypes(v)
		languageModel.Types = *types
	}

	if v, ok := m["intents"].([]interface{}); ok && len(v) > 0 {
		intents := expandIntents(v)
		languageModel.Intents = *intents
	}

	return languageModel
}

func expandTypes(in []interface{}) *[]smapi_client.Types {

	types := make([]smapi_client.Types, len(in))

	for idx, this_in := range in {
		m := this_in.(map[string]interface{})

		if v, ok := m["name"].(string); ok {
			types[idx].Name = v
		}

		if v, ok := m["values"].([]interface{}); ok {

			typeValues := expandTypeValues(v)
			types[idx].Values = *typeValues
		}
	}

	return &types
}

func expandTypeValues(in []interface{}) *[]smapi_client.TypeValues {

	typeValues := make([]smapi_client.TypeValues, len(in))

	for idx, this_in := range in {

		typeValues[idx] = smapi_client.TypeValues{
			Name: smapi_client.Value{
				Value: this_in.(string),
			},
		}
	}

	return &typeValues
}

func expandIntents(in []interface{}) *[]smapi_client.Intents {

	intents := make([]smapi_client.Intents, len(in))

	for idx, this_in := range in {
		m := this_in.(map[string]interface{})

		if v, ok := m["name"].(string); ok {
			intents[idx].Name = v
		}

		if v, ok := m["samples"].([]interface{}); ok {
			samples := make([]string, len(v))

			for idx, sample := range v {
				samples[idx] = sample.(string)
			}

			intents[idx].Samples = samples
		} else {
			// samples are optional, but log the content of the samples just in case
			// there was an issue casting the samples data
			log.Printf("[DEBUG] samples extract failed:\n%s\n", v)
		}

		if v, ok := m["slots"].([]interface{}); ok {
			intents[idx].Slots = *expandSlots(v)
		}
	}

	return &intents
}

func expandSlots(in []interface{}) *[]smapi_client.Slot {

	slots := make([]smapi_client.Slot, len(in))

	for idx, this_in := range in {
		m := this_in.(map[string]interface{})

		if v, ok := m["name"].(string); ok {
			slots[idx].Name = v
		}

		if v, ok := m["type"].(string); ok {
			slots[idx].Type = v
		}
	}

	return &slots
}

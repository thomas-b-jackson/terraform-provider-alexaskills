---
page_title: "alexaskills_interaction_model_resource Data Source - terraform-provider-alexaskills"
subcategory: ""
description: |-
  Returns manifest for an alexa skill's interaction model
---

# Data Source `alexaskills_interaction_model_resource`

Returns manifest for an alexa skill's interaction model

## Example Usage

```terraform
data "alexaskills_interaction_model_resource" "example" {
  id = "foo"
}
```

## Schema

### Required

- **id** (String, Required) ID of an alexa skill.

### Optional

None
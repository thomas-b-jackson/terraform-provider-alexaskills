---
page_title: "alexaskills_interaction_model_resource Resource - terraform-provider-alexaskills"
subcategory: ""
description: |-
  Updates the interaction model associated with an alexa skill
---

# Data Source `alexaskills_interaction_model_resource`

Updates the interaction model associated with an alexa skill

## Example Usage

```terraform
data "alexaskills_interaction_model_resource" "example" {
  
  skill_id = resource.alexaskills_skill_resource.hello_world.id

  interaction_model {
    // see examples/resources
  }
}
```

## Schema

### Required

- **skill_id** (String, Required) ID of an alexa skill.

### Optional

None
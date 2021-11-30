---
page_title: "alexaskills_skill_resource Data Source - terraform-provider-alexaskills"
subcategory: ""
description: |-
  Returns manifest for an alexa skill
---

# Data Source `alexaskills_skill_resource`

Returns manifest for an alexa skill.

## Example Usage

```terraform
data "alexaskills_skill_resource" "example" {
  id = "foo"
}
```

## Schema

### Required

- **id** (String, Required) ID of an alexa skill.

### Optional

None
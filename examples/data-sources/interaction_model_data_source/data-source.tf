terraform {
  required_providers {
    alexaskills = {
      source  = "localhost/va/alexaskills"
      version = "0.2.0-beta4"
    }
  }
}

provider "alexaskills" {
  vendor_id = "M3VEWOQC3LNOOF"
}

data "alexaskills_interaction_model_resource" "va_demo" {
  skill_id = "amzn1.ask.skill.0c29632d-efd8-4d0c-a3b8-a7988c097a74"
}

output "model_demo" {
  value = data.alexaskills_interaction_model_resource.va_demo
}
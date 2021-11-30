terraform {
  required_providers {
    alexaskills = {
      version = "0.1"
      source  = "scg.com/va/alexaskills"
    }
  }
}

provider "alexaskills" {
}

data "alexaskills_interaction_model_resource" "va_demo" {
  skill_id = "amzn1.ask.skill.d5a0762e-bfed-4800-9d2a-bebbc748674c"
}

output "model_demo" {
  value = data.alexaskills_interaction_model_resource.va_demo
}
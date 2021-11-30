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

data "alexaskills_skill_resource" "va_demo" {
  id = "amzn1.ask.skill.0de4190b-8137-4f9b-b7cf-489ed1653637"
}

output "skill_demo" {
  value = data.alexaskills_skill_resource.va_demo
}
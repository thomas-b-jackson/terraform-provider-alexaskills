terraform {
  required_providers {
    alexaskills = {
      version = "0.2.0-beta3"
      source  = "scg.com/va/alexaskills"
    }
  }
}

provider "alexaskills" {
  vendor_id = "M3VEWOQC3LNOOF"
}

data "alexaskills_skill_resource" "va_demo" {
  id = "amzn1.ask.skill.95387949-b517-4fcf-b9be-4ff3f1d4fe3d"
}

output "skill_demo" {
  value = data.alexaskills_skill_resource.va_demo
}
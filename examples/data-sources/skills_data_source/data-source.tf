terraform {
  required_providers {
    alexaskills = {
      version = "0.2"
      source  = "scg.com/va/alexaskills"
    }
  }
}

provider "alexaskills" {
  vendor_id = "M3VEWOQC3LNOOF"
}

data "alexaskills_skill_resource" "va_demo" {
  id = "amzn1.ask.skill.0c29632d-efd8-4d0c-a3b8-a7988c097a74"
}

output "skill_demo" {
  value = data.alexaskills_skill_resource.va_demo
}
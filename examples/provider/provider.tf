terraform {
  required_providers {
    alexaskills = {
      version = "0.1"
      source  = "scg.com/va/alexaskills"
    }
  }
}

provider "alexaskills" {
  token = ""
  vendorid = ""
}
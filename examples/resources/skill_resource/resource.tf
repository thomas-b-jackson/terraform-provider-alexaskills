terraform {
  required_providers {
    alexaskills = {
      version = "0.2-beta3"
      source  = "scg.com/va/alexaskills"
    }
  }
}

provider "alexaskills" {
  vendor_id = "M3VEWOQC3LNOOF"
}

resource "alexaskills_skill_resource" "hello_world" {

  manifest {
    manifest_version = "1.0"

    publishing_information {
      locales {
        en_us {
          summary         = "Sample Short Description"
          example_phrases = ["alexa open hello world", "hello", "help"]
          name            = "hello world example"
          description     = "Sample Full Description"
          small_icon_uri  = "https://wrapper-dev4.dev.shcva.socalgas.com/APP_ICON.png"
          large_icon_uri  = "https://wrapper-dev4.dev.shcva.socalgas.com/APP_ICON_LARGE.png"
        }
      }

      is_available_worldwide = true
      testing_instructions   = "Sample Testing Instructions."
      category               = "KNOWLEDGE_AND_TRIVIA"
      distribution_countries = []
    }

    privacy_and_compliance {
      allows_purchases = false
      uses_personal_info = false
      is_child_directed = false
      is_export_compliant = true
      contains_ads = false
    }

    apis {
      custom {
        endpoint {
          uri = "arn:aws:lambda:us-west-2:111365482541:function:QnABot-FulfillmentLambda-JMeqq75oakh2:live"
        }
        interfaces = []
      }
    }
  }
}
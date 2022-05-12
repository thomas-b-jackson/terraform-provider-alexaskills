terraform {
  required_providers {
    alexaskills = {
      version = "0.2-beta4"
      source  = "localhost/va/alexaskills"
    }

    time = {
      source = "hashicorp/time"
      version = "0.7.2"
    }
  }
}

provider "alexaskills" {
	vendor_id = "M3VEWOQC3LNOOF"
}

resource "alexaskills_skill_resource" "integration_test" {

  manifest {
    manifest_version = "1.0"

    publishing_information {
      locales {
        en_us {
          summary         = "Sample Short Description"
          example_phrases = ["alexa open integration test", "hello", "help"]
          name            = "integration test"
          description     = "Sample Full Description"
          small_icon_uri  = "https://wrapper-dev.dev.shcva.socalgas.com/APP_ICON.png"
          large_icon_uri  = "https://wrapper-dev.dev.shcva.socalgas.com/APP_ICON_LARGE.png"
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

resource "time_sleep" "wait_10_seconds" {
  depends_on = [alexaskills_skill_resource.integration_test]

  create_duration = "10s"
}

resource "alexaskills_interaction_model_resource" "integration_test" {
  
	# need the skill's id and to wait for the skill for be fully created
	depends_on = [alexaskills_skill_resource.integration_test,
                time_sleep.wait_10_seconds]
  
	# skill id
	skill_id = alexaskills_skill_resource.integration_test.id
  
	interaction_model {
  
	  language_model {
		invocation_name = "integration test"
  
		types {
		  name = "FAQ"
		  values = ["this", "that"]
		}
  
		intents {
		  slots {
			name = "QnA_slot"
			type = "FAQ"
		  }
  
		  name = "Qna_intent"
  
		  samples = ["{QnA_slot}"]
		}
  
		intents {
		  name = "AMAZON.StopIntent"
          samples = []
		}
  
		intents {
		  name = "AMAZON.RepeatIntent"
          samples = []
		}
  
		intents {
		  name = "AMAZON.FallbackIntent"
          samples = []
		}
  
		intents {
		  name = "AMAZON.CancelIntent"
          samples = []
		}

        intents {
		  name = "AMAZON.NavigateHomeIntent"
          samples = []
		}
	  }
	}
}

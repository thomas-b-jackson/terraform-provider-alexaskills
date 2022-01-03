terraform {
  required_providers {
    alexaskills = {
      version = "0.1.0-beta1"
      source  = "thomas-b-jackson/va/alexaskills"
    }
  }
}

provider "alexaskills" {
}

resource "alexaskills_skill_resource" "hello_world" {

	manifest {
	  manifest_version = "1.0"
  
	  publishing_information {
		locales {
		  en_us {
			summary         = "Sample Short Description"
			example_phrases = ["help my gas is leaking", 
							  "pilot light is out", 
							  "can't remember my password"]
			name            = "TestAccInteractionModel"
			description     = "Sample Full Description"
		  }
		}
  
		is_available_worldwide = true
		testing_instructions   = "Sample Testing Instructions."
		category               = "KNOWLEDGE_AND_TRIVIA"
		distribution_countries = []
	  }
  
	  apis {
		custom {
		  endpoint {
			uri = "arn:aws:lambda:us-west-2:580753938011:function:serverlessrepo-alexa-skil-alexaskillskitnodejsfact-pymFhOcUAodv"
		  }
		  interfaces = []
		}
	  }
	}
  }

resource "alexaskills_interaction_model_resource" "hello_world" {
  
	# need the skill's id
	depends_on = [resource.alexaskills_skill_resource.hello_world]
  
	# skill id
	skill_id = resource.alexaskills_skill_resource.hello_world.id
  
	interaction_model {
  
	  language_model {
		invocation_name = "my socal gas"
  
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

  resource "null_resource" "null_resource_simple" {
    depends_on = [resource.alexaskills_interaction_model_resource.hello_world]
    provisioner "local-exec" {
        command = "sleep 30s"
    }
  }
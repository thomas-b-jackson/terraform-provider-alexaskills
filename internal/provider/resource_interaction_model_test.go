package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccInteractionModel(t *testing.T) {
	// t.Skip("need to figure out how to run this test w/out actually needing a real skill to test against")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInteractionModel,
			},
		},
	})
}

const testAccResourceInteractionModel = `

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
`

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

data "alexaskills_skill_resource" "hello_world" {
  id = "amzn1.ask.skill.594c2523-5808-4edc-95e5-3be333699d5e"
}

resource "alexaskills_interaction_model_resource" "hello_world" {

  # add interaction model to the skill
  skill_id = data.alexaskills_skill_resource.hello_world.id

  interaction_model {

    language_model {
      invocation_name = "hello my world"

      types {
        name = "FAQ"
        values = [
          "alexa open hello world", 
          "hello world", 
          "help",
          "help me"]
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
        name ="AMAZON.NavigateHomeIntent"
        samples = []
      }
    }
  }
}


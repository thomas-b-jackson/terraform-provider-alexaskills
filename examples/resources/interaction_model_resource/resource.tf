terraform {
  required_providers {
    alexaskills = {
      version = "0.1.0-beta1"
      source  = "thomas-b-jackson/va/alexaskills"
    }

    aws = {}
  }
}

provider "alexaskills" {
}

provider "aws" {
  region = "us-west-2"
}

# create a bot/intent/slot to use with model example
module "bot" {
  source = "./lex_bot"
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
          name            = "my socal gas"
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
          # TODO: generalize. currently using the lambda for the QnABot stack in the Logic AWS account
          uri = "arn:aws:lambda:us-west-2:111365482541:function:scg-shcva-dev-wus2-lambda-fulfillment"
        }
        interfaces = []
      }
    }
  }
}

resource "alexaskills_interaction_model_resource" "hello_world" {

  # need the intents from the bot
  depends_on = [module.bot,
                resource.alexaskills_skill_resource.hello_world]

  # add interaction model to the skill
  skill_id = resource.alexaskills_skill_resource.hello_world.id

  interaction_model {

    language_model {
      invocation_name = "my socal gas"

      types {
        name = "FAQ"
        values = [for q in module.bot.bot_questions : q.value]
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


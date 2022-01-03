terraform {
  required_providers {
    alexaskills = {
      version = "0.1-beta0"
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
          example_phrases = ["alexa open hello world", "hello", "help"]
          name            = "hello world example"
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
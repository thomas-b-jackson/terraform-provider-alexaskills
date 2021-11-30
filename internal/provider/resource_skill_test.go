package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSkills(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSkills,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"alexaskills_skill_resource.foo", "id", regexp.MustCompile("amzn1.ask.skill.*")),
				),
			},
		},
	})
}

const testAccResourceSkills = `
resource "alexaskills_skill_resource" "foo" {
	manifest {
		manifest_version = "1.0"

		publishing_information {
			locales {
				en_us {
					summary         = "Sample Short Description"
					example_phrases = ["alexa open hello world", "hello", "help"]
					name            = "TestAccResourceSkills"
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
`

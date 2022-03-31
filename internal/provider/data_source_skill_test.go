package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSkill(t *testing.T) {
	// t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSkill,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.alexaskills_skill_resource.foo", "id", regexp.MustCompile("amzn1.ask.skill.*")),
				),
			},
		},
	})
}

const testAccDataSourceSkill = `

provider "alexaskills" {
  token = "foo"
	vendor_id = "bar"
}

resource "alexaskills_skill_resource" "foo" {
	manifest {
		manifest_version = "1.0"

		publishing_information {
			locales {
				en_us {
					summary         = "Sample Short Description"
					example_phrases = ["alexa open hello world", "hello", "help"]
					name            = "TestAccDataSourceSkill"
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

data "alexaskills_skill_resource" "foo" {
  id = resource.alexaskills_skill_resource.foo.id
}
`

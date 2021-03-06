package smapi_client

import (
	"fmt"
	"testing"
)

func TestGetSkill(t *testing.T) {

	f := func(token string, verb string, url string, payload []byte) (SMAPIResponse, error) {

		var vendorSkills = `{"Skills":
		[{
			"ID":"amzn1.ask.skill.70946374-9b01-4e18-b5e0-e0ec292a0f53",
			"LastUpdated": "2021-09-13T16:33:35.245Z", 
			"NameByLocale": {
				 "EnglishUS": "QnA Bot"
			},
			"PublicationStatus": "DEVELOPMENT",
			"Stage": "development", 
			"ASIN": ""
		}]
		}`

		return SMAPIResponse{200, []byte(vendorSkills)}, nil
	}

	smapiClient, _ := NewTestClient(f)

	skills, err := smapiClient.GetSkill("amzn1.ask.skill.70946374-9b01-4e18-b5e0-e0ec292a0f53")

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if len(skills.Skills) == 0 {
		t.Log("no skills returned")
		t.Fail()
	}

	fmt.Printf("%+v\n", skills)
}

func TestCreateSkill(t *testing.T) {

	f := func(token string, verb string, url string, payload []byte) (SMAPIResponse, error) {

		var createSkillResponse = `{
			"skillId": "amzn1.ask.skill.1572d73d-0c2e-49d7-9a6b-c19dcb4383c0"
		}`

		return SMAPIResponse{202, []byte(createSkillResponse)}, nil
	}

	smapiClient, _ := NewTestClient(f)

	skillManifest := SkillManifest{
		PublishingInformation: PublishingInformation{
			Locales: Locales{
				EnglishUS: EnglishUSLocal{
					Summary:        "Sample Short Description",
					ExamplePhrases: [](string){"alexa open hello world", "hello", "help"},
					Name:           "",
					Description:    "Sample Full Description",
				},
			},
			IsAvailableWorldwide:  true,
			TestingInstructions:   "some instructions",
			Category:              "wheels",
			DistributionCountries: [](string){},
		},
		Apis: Apis{
			Custom: CustomApi{
				Endpoint: Endpoint{
					Uri: "https://foo.bar",
				},
				Interfaces: [](string){},
			},
		},
	}

	// happy path
	skillId, err := smapiClient.CreateSkill(skillManifest)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if skillId != "amzn1.ask.skill.1572d73d-0c2e-49d7-9a6b-c19dcb4383c0" {
		t.Log("skill ID not returned")
		t.Fail()
	}

	// unhappy-path
	unhappy := func(token string, verb string, url string, payload []byte) (SMAPIResponse, error) {

		var createSkillResponse = `{
			"headers": [
			  {
				"key": "content-type",
				"value": "application/json"
			  },
			  {
				"key": "content-length",
				"value": "66"
			  },
			  {
				"key": "connection",
				"value": "close"
			  }
			],
			"statusCode": 400
		  }`

		return SMAPIResponse{400, []byte(createSkillResponse)}, nil
	}

	smapiClient, _ = NewTestClient(unhappy)

	_, err = smapiClient.CreateSkill(skillManifest)

	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

func TestDeleteSkill(t *testing.T) {

	f := func(token string, verb string, url string, payload []byte) (SMAPIResponse, error) {

		var deleteSkillResponse = `{
			"headers": [
			  {
				"key": "content-type",
				"value": "application/json"
			  },
			  {
				"key": "content-length",
				"value": "66"
			  },
			  {
				"key": "connection",
				"value": "close"
			  }
			],
			"statusCode": 204
		  }`

		return SMAPIResponse{204, []byte(deleteSkillResponse)}, nil
	}

	smapiClient, _ := NewTestClient(f)

	err := smapiClient.DeleteSkill("some-skill-id")

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	unhappy := func(token string, verb string, url string, payload []byte) (SMAPIResponse, error) {

		var deleteSkillResponse = `{
			"headers": [
			  {
				"key": "content-type",
				"value": "application/json"
			  },
			  {
				"key": "content-length",
				"value": "66"
			  },
			  {
				"key": "connection",
				"value": "close"
			  }
			],
			"statusCode": 500
		  }`

		return SMAPIResponse{500, []byte(deleteSkillResponse)}, nil
	}

	smapiClient, _ = NewTestClient(unhappy)

	err = smapiClient.DeleteSkill("some-skill-id")

	if err == nil {
		t.Log("error should be not be nil", err)
		t.Fail()
	}
}

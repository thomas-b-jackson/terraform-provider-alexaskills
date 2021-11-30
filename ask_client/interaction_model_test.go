package ask_client

import (
	"fmt"
	"testing"
)

func TestGetInteractionModel(t *testing.T) {

	f := func(name string, arg ...string) (string, error) {

		var interactionModel = `{
			"interactionModel": {
			  "languageModel": {
				"invocationName": "my logic",
				"intents": [
				  {
					"name": "AMAZON.CancelIntent",
					"samples": []
				  },
				  {
					"name": "AMAZON.HelpIntent",
					"samples": []
				  },
				  {
					"name": "AMAZON.StopIntent",
					"samples": []
				  },
				  {
					"name": "PaidHolidayIntent",
					"slots": [],
					"samples": [
					  "holidays",
					  "holidays this month",
					  "next holiday",
					  "when is the next holiday"
					]
				  },
				  {
					"name": "TicketIntent",
					"slots": [],
					"samples": [
					  "about IT tickets",
					  "how do I file a ticket",
					  "help with IT issue",
					  "get help from help desk"
					]
				  },
				  {
					"name": "BenefitsEnrollmentIntent",
					"slots": [],
					"samples": [
					  "when is open enrollment",
					  "changing my benefits",
					  "when can I change my benefits",
					  "enroll for benefits",
					  "how can I change my benefits",
					  "changing health care benefits"
					]
				  },
				  {
					"name": "AMAZON.NavigateHomeIntent",
					"samples": []
				  }
				],
				"types": []
			  }
			},
			"version": "3"
		  }`

		return interactionModel, nil
	}

	askClient, _ := NewTestClient(f)

	model, err := askClient.GetInteractionModel("amzn1.ask.skill.70946374-9b01-4e18-b5e0-e0ec292a0f53")

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if len(model.InteractionModel.LanguageModel.Intents) == 0 {
		t.Log("no intents returned")
		t.Fail()
	}

	fmt.Printf("%+v\n", model)
}

func TestCreateInteractionModel(t *testing.T) {

	happyPath := func(name string, arg ...string) (string, error) {

		var createSkillResponse = `{
			"body": {
			  "skillId": "amzn1.ask.skill.1572d73d-0c2e-49d7-9a6b-c19dcb4383c0"
			},
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
			"statusCode": 202
		  }`

		return createSkillResponse, nil
	}

	askClient, _ := NewTestClient(happyPath)

	interactionModel := InteractionModel{
		LanguageModel: LanguageModel{
			InvocationName: "my socal gas",
			Intents: []Intents{
				{
					Name:    "AMAZON.CancelIntent",
					Samples: []string{},
				},
			},
			Types: []Types{
				{
					Name: "FAQ",
					Values: []TypeValues{
						{
							Name: Value{
								Value: "I forgot my password",
							},
						},
						{
							Name: Value{
								Value: "Can't remember my password",
							},
						},
					},
				},
			},
		},
	}

	// happy path
	err := askClient.UpdateInteractionModel("amzn1.ask.skill.1572d73d-0c2e-49d7-9a6b-c19dcb4383c0",
		interactionModel)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	// unhappy-path
	unhappy := func(name string, arg ...string) (string, error) {

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

		return createSkillResponse, nil
	}

	askClient, _ = NewTestClient(unhappy)

	err = askClient.UpdateInteractionModel("amzn1.ask.skill.1572d73d-0c2e-49d7-9a6b-c19dcb4383c0",
		interactionModel)

	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}
}

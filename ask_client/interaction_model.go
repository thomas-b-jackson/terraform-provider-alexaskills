package ask_client

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Intents struct {
	Name    string   `json:"name"`
	Samples []string `json:"samples"`
	Slots   []Slot   `json:"slots"`
}

type Slot struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type LanguageModel struct {
	InvocationName string    `json:"invocationName"`
	Intents        []Intents `json:"intents"`
	Types          []Types   `json:"types"`
}

type Types struct {
	Name   string       `json:"name"`
	Values []TypeValues `json:"values"`
}

type TypeValues struct {
	Name Value `json:"name"`
}

type Value struct {
	Value string `json:"value"`
}
type InteractionModel struct {
	LanguageModel LanguageModel `json:"languageModel"`
}

type InteractionModelObj struct {
	InteractionModel InteractionModel `json:"interactionModel"`
}

type UpdateModelResponse struct {
	StatusCode int `json:"statusCode"`
	Body       struct {
		ID string `json:"skillId"`
	} `json:"body"`
}

// wait up to this many seconds for a model build to complete
const ModelBuildTimeoutSec = 30

func (c *AskClient) GetInteractionModel(skillId string) (InteractionModelObj, error) {

	askOutput, err := c.Exec("ask",
		"smapi",
		"get-interaction-model",
		"--stage",
		"development",
		"--locale",
		"en-US",
		"-s",
		skillId)

	var model InteractionModelObj

	if err != nil {
		log.Printf("[DEBUG] get skill model raw output:\n%s\n", askOutput)
		return model, err
	}

	err = json.Unmarshal([]byte(askOutput), &model)

	return model, err
}

func (c *AskClient) UpdateInteractionModel(skillId string, model InteractionModel) error {

	// include the outer tag in the model
	interactionModelObj := InteractionModelObj{model}

	modelBytes, err := json.Marshal(interactionModelObj)

	if err != nil {
		log.Printf("[DEBUG] interaction model update marshalled bytes:\n%s\n", modelBytes)
		log.Printf("[DEBUG] interaction model update:\n%+v\n", model)
		return err
	}

	// --full-response will ensure pure json is returned
	askOutput, err := c.Exec("ask",
		"smapi",
		"set-interaction-model",
		"--full-response",
		"--stage",
		"development",
		"--locale",
		"en-US",
		"--skill-id",
		skillId,
		"--interaction-model",
		string(modelBytes))

	if err != nil {
		log.Printf("[DEBUG] interaction model update raw output:\n%s\n", askOutput)
		log.Printf("[DEBUG] interaction model update marshalled bytes:\n%s\n", modelBytes)
		return err
	}

	var updateModelResponse UpdateModelResponse

	// load the response into a struct
	err = json.Unmarshal([]byte(askOutput), &updateModelResponse)

	if err != nil {
		// un-marshall failed
		log.Printf("[DEBUG] interaction model update raw output:\n%s\n", askOutput)
		log.Printf("[DEBUG] interaction model update output:\n%+v\n", updateModelResponse)
		return err
	}

	if updateModelResponse.StatusCode != 202 {
		// ask cli returned an unhappy response code
		return fmt.Errorf("ask interaction model update command failed with output:\n%s", askOutput)
	}

	// Wait (up to a threshold) until the model can be read back out.
	// This is needed because a model build is initiated once the model is received,
	// and the model is not readable (or usable) until the build is complete
	expiredTimeSec := 0
	for {
		_, err := c.GetInteractionModel(skillId)
		if err == nil || expiredTimeSec >= ModelBuildTimeoutSec {
			break
		}
		// sleep for X seconds
		sleepDurationSec := 5
		time.Sleep(time.Duration(sleepDurationSec) * time.Second)
		expiredTimeSec += expiredTimeSec
	}

	return err
}

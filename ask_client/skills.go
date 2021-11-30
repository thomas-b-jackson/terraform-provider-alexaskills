package ask_client

import (
	"encoding/json"
	"fmt"
	"log"
)

type VendorSkills struct {
	Skills []struct {
		ID           string `json:"skillId"`
		LastUpdated  string `json:"lastUpdated"`
		NameByLocale struct {
			EnglishUS string `json:"en-US"`
		}
		PublicationStatus string `json:"publicationStatus"`
		Stage             string `json:"stage"`
		ASIN              string `json:"asin"`
	} `json:"skills"`
}

type CreateSkillResponse struct {
	StatusCode int `json:"statusCode"`
	Body       struct {
		ID string `json:"skillId"`
	} `json:"body"`
}

type DeleteSkillResponse struct {
	StatusCode int `json:"statusCode"`
}

type EnglishUSLocal struct {
	Summary        string   `json:"summary"`
	ExamplePhrases []string `json:"examplePhrases"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
}

type Locales struct {
	EnglishUS EnglishUSLocal `json:"en-US"`
}

type PublishingInformation struct {
	Locales               Locales  `json:"locales"`
	IsAvailableWorldwide  bool     `json:"isAvailableWorldwide"`
	TestingInstructions   string   `json:"testingInstructions"`
	Category              string   `json:"category"`
	DistributionCountries []string `json:"distributionCountries"`
}

type Endpoint struct {
	Uri string `json:"uri"`
}

type CustomApi struct {
	Endpoint   Endpoint `json:"endpoint"`
	Interfaces []string `json:"interfaces"`
}

type Apis struct {
	Custom CustomApi `json:"custom"`
}

type SkillManifest struct {
	PublishingInformation PublishingInformation `json:"publishingInformation"`
	Apis                  Apis                  `json:"apis"`
	ManifestVersion       string                `json:"manifestVersion"`
}

type SkillManifestWrapper struct {
	Manifest SkillManifest `json:"manifest"`
}

func (c *AskClient) GetSkill(skillId string) (VendorSkills, error) {

	askOutput, err := c.Exec("ask",
		"smapi",
		"list-skills-for-vendor",
		"-s",
		skillId)

	var skills VendorSkills

	if err != nil {
		log.Printf("[DEBUG] skills list raw output:\n%s\n", askOutput)
		return skills, err
	}

	err = json.Unmarshal([]byte(askOutput), &skills)

	return skills, err
}

func (c *AskClient) GetSkillManifest(skillId string) (SkillManifest, error) {

	askOutput, err := c.Exec("ask",
		"smapi",
		"get-skill-manifest",
		"-s",
		skillId)

	var wrappedSkill SkillManifestWrapper

	if err != nil {
		log.Printf("[DEBUG] skills get manifest raw output:\n%s\n", askOutput)
		return SkillManifest{}, err
	}

	err = json.Unmarshal([]byte(askOutput), &wrappedSkill)

	return wrappedSkill.Manifest, err
}

func (c *AskClient) CreateSkill(skillManifest SkillManifest) (string, error) {

	manifestComplete := SkillManifestWrapper{skillManifest}

	manifestBytes, err := json.Marshal(manifestComplete)

	var createSkillResponse CreateSkillResponse
	var skillId string

	if err != nil {
		log.Printf("[DEBUG] skills marshalled manifest:\n%s\n", manifestBytes)
		log.Printf("[DEBUG] skills manifest:\n%+v\n", manifestComplete)
		return skillId, err
	}

	// --full-response will ensure pure json is returned
	askOutput, err := c.Exec("ask",
		"smapi",
		"create-skill-for-vendor",
		"--full-response",
		"--manifest",
		string(manifestBytes))

	if err != nil {
		log.Printf("[DEBUG] skills create raw output:\n%s\n", askOutput)
		return skillId, err
	}

	// load the response into a struct
	err = json.Unmarshal([]byte(askOutput), &createSkillResponse)

	if err != nil {
		// un-marshall failed
		log.Printf("[DEBUG] skills create raw output:\n%s\n", askOutput)
		log.Printf("[DEBUG] skills create output:\n%+v\n", createSkillResponse)
		return skillId, err
	}

	if createSkillResponse.StatusCode != 202 {
		// ask cli returned an unhappy response code
		return skillId, fmt.Errorf("ask skill creation command failed with output:\n%s", askOutput)
	}

	skillId = createSkillResponse.Body.ID

	return skillId, err
}

func (c *AskClient) DeleteSkill(skillId string) error {

	askOutput, err := c.Exec("ask",
		"smapi",
		"delete-skill",
		"--full-response",
		"--skill-id",
		skillId)

	if err != nil {
		log.Printf("[DEBUG] skills delete raw output:\n%s\n", askOutput)
		return err
	}

	var deleteSkillResponse DeleteSkillResponse

	err = json.Unmarshal([]byte(askOutput), &deleteSkillResponse)

	if err != nil {
		log.Printf("[DEBUG] skills deletion unsuccessful. raw output:\n%s\n", askOutput)
		return err
	}

	if deleteSkillResponse.StatusCode != 204 {
		// ask cli returned an unhappy response code
		return fmt.Errorf("ask skill deletion command failed with output:\n%s", askOutput)
	}

	return err
}

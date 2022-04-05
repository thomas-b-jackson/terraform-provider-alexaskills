package smapi_client

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
	ID string `json:"skillId"`
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

type SkillManifestUpdateWrapper struct {
	Manifest SkillManifest `json:"manifest"`
}

type SkillManifestWrapper struct {
	VendorId string        `json:"vendorId"`
	Manifest SkillManifest `json:"manifest"`
}

func (c *SMAPIClient) GetSkill(skillId string) (VendorSkills, error) {

	smapiResponse, err := c.Get(
		"/v1/skills?vendorId=" + c.vendorId + "&skillId=" + skillId)

	var skills VendorSkills

	if err != nil {
		log.Printf("[DEBUG] skills list raw output:\n%+v\n", smapiResponse)
		return skills, err
	}

	err = json.Unmarshal(smapiResponse.Body, &skills)

	return skills, err
}

func (c SMAPIClient) GetSkillManifest(skillId string) (SkillManifest, error) {

	smapiResponse, err := c.Get(
		"/v1/skills/" + skillId + "/stages/development/manifest")

	if err != nil {
		log.Printf("[DEBUG] Error getting skill manifest for :%s\n", skillId)
		return SkillManifest{}, err
	}

	var wrappedSkill SkillManifestWrapper

	err = json.Unmarshal(smapiResponse.Body, &wrappedSkill)

	return wrappedSkill.Manifest, err
}

func (c *SMAPIClient) CreateSkill(skillManifest SkillManifest) (string, error) {

	manifestComplete := SkillManifestWrapper{c.vendorId, skillManifest}

	manifestBytes, err := json.Marshal(manifestComplete)

	var createSkillResponse CreateSkillResponse
	var skillId string

	if err != nil {
		log.Printf("[DEBUG] skills marshalled manifest:\n%s\n", manifestBytes)
		log.Printf("[DEBUG] skills manifest:\n%+v\n", manifestComplete)
		return skillId, err
	}

	smapiResponse, err := c.Post(
		"/v1/skills",
		manifestBytes)

	if err != nil {
		log.Printf("[DEBUG] skills create raw output:\n%+v\n", smapiResponse)
		return skillId, err
	}

	// load the response into a struct
	err = json.Unmarshal(smapiResponse.Body, &createSkillResponse)

	if err != nil {
		// un-marshall failed
		log.Printf("[DEBUG] skills create unmarshal failure ... raw output:\n%+v\n", smapiResponse)
		return skillId, err
	}

	if smapiResponse.Status != 202 {
		// smapi returned an unhappy response code
		return skillId, fmt.Errorf("skill creation failed with response code: %d and message: \n%s\n", smapiResponse.Status, smapiResponse.Body)
	}

	skillId = createSkillResponse.ID

	return skillId, err
}

func (c *SMAPIClient) UpdateSkill(skillId string, skillManifest SkillManifest) error {

	manifestBytes, err := json.Marshal(SkillManifestUpdateWrapper{skillManifest})

	if err != nil {
		log.Printf("[DEBUG] skills update marshalled manifest:\n%s\n", manifestBytes)
		log.Printf("[DEBUG] skills update manifest:\n%+v\n", skillManifest)
		return err
	}

	smapiResponse, err := c.Put(
		"/v1/skills/"+skillId+"/stages/development/manifest",
		manifestBytes)

	if err != nil {
		log.Printf("[DEBUG] skills update raw output:\n%+v\n", smapiResponse)
		return err
	}

	if smapiResponse.Status != 202 {
		// smapi returned an unhappy response code
		log.Printf("[DEBUG] skills update manifest:\n%s\n", manifestBytes)
		return fmt.Errorf("skill update failed with response code: %d, and body:\n%s", smapiResponse.Status, smapiResponse.Body)
	}

	return err
}

func (c *SMAPIClient) DeleteSkill(skillId string) error {

	smapiResponse, err := c.Delete(
		"/v1/skills/" + skillId)

	if err != nil {
		log.Printf("[DEBUG] skills delete raw output:\n%+v\n", smapiResponse)
		return err
	}

	if smapiResponse.Status != 204 {
		// smapi returned an unhappy response code
		return fmt.Errorf("skill deletion command failed with output:\n%s\n", smapiResponse)
	}

	return err
}

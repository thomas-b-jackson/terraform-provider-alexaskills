package smapi_client

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

const baseURL = "http://api.amazonalexa.com"
const token = "Atza|IwEBIL9MIQsfiawXiC4kreIgWmHjSNOcfWFEEbkOu2esPe_E1-O20zhJxRGGelkiavzyODLDaINmqhRei_lzUDZiQW6BtuRxeYCHWz4d57kFeZWrMyfZ0jUeha_oY649qGz79n9XQUtx8v4_uGZDPEJPB-V2vv8-_d4C_8CVhMuFE9D7dg7Lrf-Jua84_JqKJF14TKMlRk3WxzkBL0tiARKkXMT3wrPLLMQjCByWLkHrkIkKcrb0KoRy5Mwz3TFfqq1yKThqESE_2i5SuddsZyE5scIkR52fZ1OtWLW_3PbHyn9dIjzWui7gQWC6cYck1PRf7bvQPEMmcWIC0UKe0BDUTkmX-nuL0fwpKs84OSoTTv3ISmJ-tjJomCiDLvEdI816tMmrhGonvukmKxGn7Fw-pYLr4z94ZeMJ_bEPPuZfS1zFq4jTteznRCjK1D1g00dOM5iEsx9uRLViGOg607y1l1xD1eefBY_XZawPnrZbLc69jSqv2juaDhR2siBU37NUM1Poa_NmqGfYltgu0ixpxj9_gQCY8hRXEt55eocFmNeNavwrQmk7Mg8qKSwjvOZhtnQKPMzxqD8gimIVqDHegjFQr1bvabczJYrz7ztBXpVHq-Vg91d9ANGMXeQY68OPiuTAl_MBEmBkri-HJjqudxh0"

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


func (c *SMAPIClient) GetSkill(token string, vendorId string, skillId string) (VendorSkills, error) {

	body, err := c.Exec(
		token,
		"GET",
		baseURL + "/v1/skills?vendorId=" + vendorId + "&skillId=" + skillid)

	var skills VendorSkills

	if err != nil {
		log.Printf("[DEBUG] skills list raw output:\n%s\n", askOutput)
		return skills, err
	}

	err = json.Unmarshal([]byte(body), &skills)

	return skills, err
}


func (c SMAPIClient) GetSkillManifest(token string, skillId string) (SkillManifest, error) {

	body, err := c.Exec(
		token,
		"GET",
		baseURL + "v1/skills/" + skillId + "/stages/development/manifest")

	if err != nil {
		log.Printf("[DEBUG] Error getting skill manifest for :%s\n",
			skillId)
		return SkillManifest{}, err
	}

	var wrappedSkill SkillManifestWrapper

	err = json.Unmarshal([]byte(body), &wrappedSkill)

	return wrappedSkill.Manifest, err
}


func (c *SMAPIClient) CreateSkill(token string, skillManifest SkillManifest) (string, error) {

	manifestComplete := SkillManifestWrapper{skillManifest}

	manifestBytes, err := json.Marshal(manifestComplete)

	var createSkillResponse CreateSkillResponse
	var skillId string

	if err != nil {
		log.Printf("[DEBUG] skills marshalled manifest:\n%s\n", manifestBytes)
		log.Printf("[DEBUG] skills manifest:\n%+v\n", manifestComplete)
		return skillId, err
	}

	body, err := c.Exec(
		token,
		"POST",
		baseURL + "v1/skills",
		string(manifestBytes))

	if err != nil {
		log.Printf("[DEBUG] skills create raw output:\n%s\n", body)
		return skillId, err
	}

	// load the response into a struct
	err = json.Unmarshal([]byte(body), &createSkillResponse)

	if err != nil {
		// un-marshall failed
		log.Printf("[DEBUG] skills create raw output:\n%s\n", body)
		log.Printf("[DEBUG] skills create output:\n%+v\n", createSkillResponse)
		return skillId, err
	}

	if createSkillResponse.StatusCode != 202 {
		// ask cli returned an unhappy response code
		return skillId, fmt.Errorf("ask skill creation command failed with output:\n%s", body)
	}

	skillId = createSkillResponse.Body.ID

	return skillId, err
}


func (c *SMAPIClient) DeleteSkill(token string, skillId string) error {

	body, err := c.Exec(
		token,
		"DELETE",
		baseURL + "v1/skills/" + skillid + "/",)

	if err != nil {
		log.Printf("[DEBUG] skills delete raw output:\n%s\n", body)
		return err
	}

	var deleteSkillResponse DeleteSkillResponse

	err = json.Unmarshal([]byte(body), &deleteSkillResponse)

	if err != nil {
		log.Printf("[DEBUG] skills deletion unsuccessful. raw output:\n%s\n", body)
		return err
	}

	if deleteSkillResponse.StatusCode != 204 {
		// ask cli returned an unhappy response code
		return fmt.Errorf("ask skill deletion command failed with output:\n%s", body)
	}

	return err
}

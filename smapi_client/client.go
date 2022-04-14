package smapi_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SMAPIResponse struct {
	Status int
	Body   []byte
}

type LWATokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Doer func(token string, verb string, url string, payload []byte) (SMAPIResponse, error)

func do(token string, verb string, url string, payload []byte) (SMAPIResponse, error) {

	const baseUrl = "https://api.amazonalexa.com"

	// log.Printf("[DEBUG] token used: %s\n", token)

	hc := &http.Client{}

	var payloadReader io.Reader

	if payload != nil {
		payloadReader = bytes.NewBuffer(payload)
	}

	request, err := http.NewRequest(verb, baseUrl+url, payloadReader)

	if err != nil {
		log.Printf("[DEBUG] Error building %s for %s\n", verb, url)
	}

	if token != "" {
		request.Header.Add("Authorization", token)
	}

	response, err := hc.Do(request)
	if err != nil {
		log.Printf("[DEBUG] Error doing %s for %s\n", verb, url)
		log.Fatal(err)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return SMAPIResponse{response.StatusCode, bytes}, err
}

type SMAPIClient struct {
	do       Doer
	token    string
	vendorId string
}

func NewClient(token string, vendorId string) (*SMAPIClient, error) {
	c := SMAPIClient{do, token, vendorId}
	return &c, nil
}

func (c *SMAPIClient) Get(url string) (SMAPIResponse, error) {
	return c.do(c.token, "GET", url, nil)
}

func (c *SMAPIClient) Post(url string, payload []byte) (SMAPIResponse, error) {
	return c.do(c.token, "POST", url, payload)
}

func (c *SMAPIClient) Put(url string, payload []byte) (SMAPIResponse, error) {
	return c.do(c.token, "PUT", url, payload)
}

func (c *SMAPIClient) Delete(url string) (SMAPIResponse, error) {
	return c.do(c.token, "DELETE", url, nil)
}

func GetLwaToken(
	lwa_client_id string,
	lwa_client_secret string,
	lwa_refresh_token string) (string, error) {

	const tokenUrl = "https://api.amazon.com/auth/o2/token"

	hc := &http.Client{}

	var payloadReader io.Reader

	request, err := http.NewRequest("POST", tokenUrl, payloadReader)

	if err != nil {
		return "", fmt.Errorf("error creating request for refresh token. error: %s", err)
	}

	query := request.URL.Query()
	query.Add("grant_type", "refresh_token")
	query.Add("refresh_token", lwa_refresh_token)
	query.Add("client_id", lwa_client_id)
	query.Add("client_secret", lwa_client_secret)
	request.URL.RawQuery = query.Encode()

	response, err := hc.Do(request)
	if err != nil {
		return "", fmt.Errorf("error retrieving refresh token. error: %v", err)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading bytes for refresh token. error: %s", err)
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("error retrieving refresh token. Status code: %d, response body: %s", response.StatusCode, string(bytes))
	}

	var lwaTokenResponse LWATokenResponse

	err = json.Unmarshal(bytes, &lwaTokenResponse)

	return lwaTokenResponse.AccessToken, err
}

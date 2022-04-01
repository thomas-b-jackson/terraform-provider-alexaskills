package smapi_client

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type Doer func(token string, verb string, url string, payload []byte) (string, error)

func do(token string, verb string, url string, payload []byte) (string, error) {

	const baseUrl = "https://api.amazonalexa.com"

	log.Printf("[DEBUG] token used: %s\n", token)

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
	body := string(bytes)
	// fmt.Println(body)

	return body, err
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

func (c *SMAPIClient) Get(url string) (string, error) {
	return c.do(c.token, "GET", url, nil)
}

func (c *SMAPIClient) Post(url string, payload []byte) (string, error) {
	return c.do(c.token, "POST", url, payload)
}

func (c *SMAPIClient) Put(url string, payload []byte) (string, error) {
	return c.do(c.token, "PUT", url, payload)
}

func (c *SMAPIClient) Delete(url string) (string, error) {
	return c.do(c.token, "DELETE", url, nil)
}

package smapi_client

import (
	"context"
	"os/exec"
	"time"
	"bytes"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Executer func(token string, verb string, url string, payload ...string) (string, error)

func osExec(token string, verb string, url string, payload ...string) (body, error) {
	hc := &http.Client{}

	request, err := http.NewRequest(verb, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("[DEBUG] Error building %s for %s\n", verb, url)
	}

	if token != "" {
		request.Header.Add("Authorization", token)
	}

	response, err := hc.Do(request)
	if err != nil {
		log.Printf("[DEBUG] Error doing %s for %s\n", verb, url )
		log.Fatal(err)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	body = string(bytes)
	fmt.Println(string(body))

	return body, err
}


type SMAPIClient struct {
	exec Executer
	token string
	vendorId string
}

func NewClient(token string, vendorId string) (*SMAPIClient, error) {
	c := SMAPIClient{osExec, token, vendorId}
	return &c, nil
}

func (c *SMAPIClient) Exec(token string, arg ...string) (string, error) {
	return c.exec(token, arg...)
}

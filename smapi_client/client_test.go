package smapi_client

import "testing"

// allow tests to pass in their own exec function
func NewSMAPITestClient(f func(token string, arg ...string) (string,
	error))(*AskClient, error) {

	c := SMAPIClient{f}
	return &c, nil
}

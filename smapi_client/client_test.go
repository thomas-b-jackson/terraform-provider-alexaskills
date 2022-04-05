package smapi_client

import "testing"

// allow tests to pass in their own do-er function
func NewTestClient(f func(token string, verb string, url string, payload []byte) (string, error)) (*SMAPIClient, error) {

	c := SMAPIClient{f, "foo", "bar"}
	return &c, nil
}

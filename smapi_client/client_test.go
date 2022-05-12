package smapi_client

// allow tests to pass in their own do-er function
func NewTestClient(f func(token string, verb string, url string, payload []byte) (SMAPIResponse, error)) (*SMAPIClient, error) {

	c := SMAPIClient{f, "foo", "bar"}
	return &c, nil
}

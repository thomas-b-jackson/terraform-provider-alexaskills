package ask_client

// allow tests to pass in their own exec function
func NewTestClient(f func(name string, arg ...string) (string, error)) (*AskClient, error) {
	c := AskClient{f}
	return &c, nil
}

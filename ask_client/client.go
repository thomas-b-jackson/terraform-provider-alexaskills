package ask_client

import (
	"context"
	"os/exec"
	"time"
)

type Executer func(name string, arg ...string) (string, error)

func osExec(name string, arg ...string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	// Create the command with our context
	cmd := exec.CommandContext(ctx, name, arg...)
	out, err := cmd.Output()
	return string(out), err
}

type AskClient struct {
	exec Executer
}

func NewClient() (*AskClient, error) {
	c := AskClient{osExec}
	return &c, nil
}

func (c *AskClient) Exec(name string, arg ...string) (string, error) {
	return c.exec(name, arg...)
}

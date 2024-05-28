package exec

import (
	"os"
	"os/exec"
)

type Command struct {
	Name string
	Args []string
}

func (c *Command) Execute() error {
	cmd := exec.Command(c.Name, c.Args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

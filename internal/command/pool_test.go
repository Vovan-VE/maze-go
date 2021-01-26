package command

import (
	"os"
	"testing"
)

func init() {
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	stdout = f
	stderr = f
}

func TestPoolRun(t *testing.T) {
	_, _ = Run([]string{"maze"})
	_, _ = Run([]string{"maze", "gen"})
	_, _ = Run([]string{"maze", "gen", "-s5x4"})
	_, _ = Run([]string{"maze", "help", "gen"})
	_, _ = Run([]string{"maze", "-h", "gen"})
	_, _ = Run([]string{"maze", "-h", "-x"})
	_, _ = Run([]string{"maze", "foo", "bar", "lol"})
	t.Skip("result is not checked")
}

func TestPoolCommand(t *testing.T) {
	pool := &tPool{}

	c := pool.Command("unknown1")
	if c != nil {
		t.Error("There must be no such command")
	}

	c = pool.Command(cmdGen)
	if c == nil {
		t.Errorf("There is no %s command", cmdGen)
	}
	if _, ok := c.(*Generate); !ok {
		t.Errorf("Command %s is not *Generate", cmdGen)
	}
}

func TestPoolCommands(t *testing.T) {
	pool := &tPool{}
	cmdMap := pool.Commands()

	c, ok := cmdMap[cmdGen]
	if !ok {
		t.Errorf("There is no %s command", cmdGen)
	}
	if _, ok := c.(*Generate); !ok {
		t.Errorf("Command %s is not *Generate", cmdGen)
	}
}

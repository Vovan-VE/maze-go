package command

import (
	"testing"
)

func TestHelpRun(t *testing.T) {
	help := NewHelp(&testHelpPool{}, "help")
	if _, err := help.Run([]string{"-x"}); err == nil || err.Error() != "flag provided but not defined: -x" {
		t.Errorf("unexpected error: %v", err)
	}
	if code, err := help.Run([]string{"-l"}); code != 0 || err != nil {
		t.Errorf("unexpected code %d: %v", code, err)
	}
	if code, err := help.Run([]string{"foo"}); code != 0 || err != nil {
		t.Errorf("unexpected code %d: %v", code, err)
	}
	if _, err := help.Run([]string{"unknown1"}); err == nil || err.Error() != "unknown help topic \"unknown1\"" {
		t.Errorf("unexpected error: %v", err)
	}
	if code, err := help.Run(nil); code != 0 || err != nil {
		t.Errorf("unexpected code %d: %v", code, err)
	}
}

func TestHelpUsage(t *testing.T) {
	if "" == NewHelp(nil, "help").Usage() {
		t.Error("empty usage")
	}
}

var testHelpCommands = []string{cmdGen, "foo"}

type testHelpPool struct {
}

func (t *testHelpPool) Command(name string) Commander {
	for _, n := range testHelpCommands {
		if n == name {
			return &testHelpCommander{name}
		}
	}
	return nil
}

func (t *testHelpPool) Commands() map[string]Commander {
	res := make(map[string]Commander, len(testHelpCommands))
	for _, name := range testHelpCommands {
		res[name] = &testHelpCommander{name}
	}
	return res
}

//var tHelpCmd1 = &testHelpCommander{"h1"}

type testHelpCommander struct {
	name string
}

func (c *testHelpCommander) Run(args []string) (int, error) {
	return 0, nil
}

func (c *testHelpCommander) Name() string {
	return c.name
}

func (c *testHelpCommander) Usage() string {
	return ""
}

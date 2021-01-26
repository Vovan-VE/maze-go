package command

import (
	"testing"
)

func TestName(t *testing.T) {
	var p Pooler
	c := newCommand(p, "foo-bar")
	if c.Name() != "foo-bar" {
		t.Error("Wrong name returned")
	}
	if c.Usage() != "" {
		t.Error("empty fallback usage")
	}
}

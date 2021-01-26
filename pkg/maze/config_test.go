package maze

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	if c.Width != 0 {
		t.Error("Bad width initialization")
	}
	if c.Height != 0 {
		t.Error("Bad height initialization")
	}
	if c.BranchLength != 0 {
		t.Error("Bad branchLength initialization")
	}
	if c.Format != "" {
		t.Error("Bad format initialization")
	}
	if len(c.Options) != 0 {
		t.Error("Bap options initialization")
	}
}

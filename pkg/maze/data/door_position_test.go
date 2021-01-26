package data

import (
	"testing"
)

func TestNewDoorPosition(t *testing.T) {
	d := NewDoorPosition(LEFT, 42)
	if d.Side() != LEFT {
		t.Error("Wrong side")
	}
	if d.Offset() != 42 {
		t.Error("Wrong offset")
	}
}

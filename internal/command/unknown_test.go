package command

import (
	"testing"
)

func TestUnknown(t *testing.T) {
	cmd := NewUnknown(nil, "foo-bar")
	code, err := cmd.Run(nil)
	if 0 == code {
		t.Errorf("code is %d", code)
	}
	if nil == err || err.Error() != "Unknown command: foo-bar" {
		t.Errorf("unexpected error: %v", err)
	}
	if "" != cmd.Usage() {
		t.Error("usage is useless")
	}
}

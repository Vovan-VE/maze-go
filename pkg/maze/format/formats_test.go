package format

import (
	"testing"
)

func TestHasExporter(t *testing.T) {
	if !HasExporter("text") {
		t.Error("Incorrect, the IS text exporter")
	}

	if HasExporter("unknown1") {
		t.Error("Incorrect, the ISN'T unknown1 exporter")
	}
}

func TestNewExporter(t *testing.T) {
	ex := NewExporter("text")
	if ex == nil {
		t.Error("Exporter was not created")
	}

	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "Unknown format name: unknown1" {
			panic(e)
		}
	}()
	ex = NewExporter("unknown1")
	if ex == nil {
		t.Error("Unknown exporter return nil")
	}
	t.Error("Unknown exporter created")
}

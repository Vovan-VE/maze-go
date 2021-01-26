package cli

import (
	"flag"
	"testing"
)

func TestNewMapValue(t *testing.T) {
	m := NewMapValue()
	if 0 != len(m.values) {
		t.Error("Must be init with empty values map")
	}
}

func TestValues(t *testing.T) {
	m := NewMapValue()
	m.Set("foo=42")
	m.Set("bar=")
	values := m.Values()
	if len(values) != 2 {
		t.Error("Incorrect length")
	}
	if v, ok := values["foo"]; !ok || v != "42" {
		t.Error("name `foo` must be set and have `42` value")
	}
	if v, ok := values["bar"]; !ok || v != "" {
		t.Error("name `bar` must be set and have `` value")
	}
}

func TestString(t *testing.T) {
	m := NewMapValue()
	m.Set("foo=42")
	m.Set("bar=")
	m.Set("baz")
	if "" != m.String() {
		t.Error("String representation must be empty because it's impossible to show all pairs in one flag")
	}
}

func TestSet(t *testing.T) {
	m := NewMapValue()
	m.Set("foo=42")
	m.Set("bar=")
	m.Set("baz")
	m.Set("foo=37")

	if len(m.values) != 3 {
		t.Error("length must be 3 since only 3 different names was used")
	}
	if v, ok := m.values["foo"]; !ok || v != "37" {
		t.Error("name `foo` must be set and have `37` value")
	}
	if v, ok := m.values["bar"]; !ok || v != "" {
		t.Error("name `bar` must be set and have `` value")
	}
	if v, ok := m.values["baz"]; !ok || v != "" {
		t.Error("name `baz` must be set and have `` value")
	}
}

func TestFlagParsing(t *testing.T) {
	m := NewMapValue()
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.Var(m, "c", "config option")
	e := fs.Parse([]string{
		"-c", "foo=42",
		"-c", "bar=lorem ipsum",
		"-c", "baz=",
		"-c", "lol",
		"-c", "foo=37",
	})
	if e != nil {
		t.Errorf("Parse failed: %v", e)
	}

	if len(m.values) != 4 {
		t.Error("length must be 4 since only 4 different names was used")
	}
	if v, ok := m.values["foo"]; !ok || v != "37" {
		t.Error("name `foo` must be set and have `37` value")
	}
	if v, ok := m.values["bar"]; !ok || v != "lorem ipsum" {
		t.Error("name `bar` must be set and have `lorem ipsum` value")
	}
	if v, ok := m.values["baz"]; !ok || v != "" {
		t.Error("name `baz` must be set and have `` value")
	}
	if v, ok := m.values["lol"]; !ok || v != "" {
		t.Error("name `lol` must be set and have `` value")
	}
}

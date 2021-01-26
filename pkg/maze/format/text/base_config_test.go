package text

import (
	"testing"
)

func TestNewBaseConfig(t *testing.T) {
	actual := newBaseConfig()
	expected := baseConfig{"#", "i", "E"}
	if *actual != expected {
		t.Errorf("Expected %v but saw %v", expected, *actual)
	}
}

func TestConfigure(t *testing.T) {
	c := newBaseConfig()
	c.configure(map[string]string{
		"in":   "()",
		"out":  "[]",
		"wall": "##",
	})
	if c.wall != "##" {
		t.Error("Wall did not changed")
	}
	if c.in != "()" {
		t.Error("In did not changed")
	}
	if c.out != "[]" {
		t.Error("Out did not changed")
	}

	c.configure(map[string]string{
		"wall": "@@@",
	})
	if c.wall != "@@@" {
		t.Error("Wall did not changed")
	}
	if c.in != "()" {
		t.Error("In must not be changed")
	}
	if c.out != "[]" {
		t.Error("Out must not be changed")
	}

	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "Unknown option \"foo\"" && e != "Unknown option \"bar\"" {
			panic(e)
		}
	}()
	c.configure(map[string]string{
		"foo": "lorem",
		"bar": "ipsum",
	})
}

func TestConfigureEmpty(t *testing.T) {
	c := newBaseConfig()
	names := []string{"wall", "in", "out"}
	for _, name := range names {
		func() {
			defer func() {
				e := recover()
				if e == nil {
					t.Error("Did not panic")
				}
				if e != "Value for \""+name+"\" must not be empty" {
					panic(e)
				}
			}()
			c.configure(map[string]string{name: ""})
		}()
	}
}

func TestRepeatStringToLength(t *testing.T) {
	cases := []struct {
		in  string
		n   int
		out string
	}{
		{"a", 0, ""},
		{"a", 3, "aaa"},
		{"abc", 5, "abcab"},
		{"abc", 8, "abcabcab"},
		{"abc", 9, "abcabcabc"},
		{"йщ№", 8, "йщ№йщ№йщ"},
	}
	for i, v := range cases {
		actual := repeatStringToLength(v.in, v.n)
		if actual != v.out {
			t.Errorf("case %d: expected %q but saw %q", i, v.out, actual)
		}
	}
}

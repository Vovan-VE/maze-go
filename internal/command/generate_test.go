package command

import (
	"fmt"
	"os"
	"testing"
)

type cfgT struct {
	w, h, bl int
	f        string
	op       map[string]string
}
type caseT struct {
	args []string
	cfg  *cfgT
}

type caseE struct {
	args []string
	err  string
}

func init() {
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	stdout = f
	stderr = f
}

func TestGenerateRun(t *testing.T) {
	g := NewGenerate(nil, "gen")
	ret, err := g.Run([]string{})
	if err != nil || ret != 0 {
		t.Errorf("code %d, error %v", ret, err)
	}

	ret, err = g.Run([]string{"--fail"})
	if err == nil || err.Error() != "flag provided but not defined: --fail" {
		t.Errorf("code %d, error %v", ret, err)
	}
}

func TestGenerateUsage(t *testing.T) {
	if "" == NewGenerate(nil, "gen").Usage() {
		t.Error("Usage should not be empty")
	}
}

func diffMaps(expected, actual map[string]string) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("expected length is %d but saw length %d", len(expected), len(actual))
	}
	for k, v := range expected {
		a, ok := actual[k]
		if !ok {
			return fmt.Errorf("expected to have key %q but it's not", k)
		}
		if a != v {
			return fmt.Errorf("expected value of key %q is %q but saw %q", k, v, a)
		}
	}
	return nil
}

func TestInitConfigOk(t *testing.T) {
	cases := []caseT{
		caseT{
			[]string{},
			&cfgT{30, 10, 10, "text", nil},
		},
		caseT{
			[]string{"-W42", "-H", "37", "-f", "text"},
			&cfgT{42, 37, 10, "text", nil},
		},
		caseT{
			[]string{"--width=42", "--height", "37", "-B", "23"},
			&cfgT{42, 37, 23, "text", nil},
		},
		caseT{
			[]string{"-s42x37", "--branch-length", "42", "--format", "text"},
			&cfgT{42, 37, 42, "text", nil},
		},
		caseT{
			[]string{"--size", "42x37", "-cfoo=bar", "-c", "lorem=ipsum"},
			&cfgT{42, 37, 10, "text", map[string]string{"foo": "bar", "lorem": "ipsum"}},
		},
	}
	for i, v := range cases {
		res, e := initConfig("cmd", v.args)
		if e != nil {
			t.Errorf("case %d: unexpected error %v", i, e)
		}
		if res.Width != v.cfg.w {
			t.Errorf("case %d: width expected %d, but was %d", i, v.cfg.w, res.Width)
		}
		if res.Height != v.cfg.h {
			t.Errorf("case %d: height expected %d, but was %d", i, v.cfg.h, res.Height)
		}
		if res.BranchLength != v.cfg.bl {
			t.Errorf("case %d: bl expected %d, but was %d", i, v.cfg.bl, res.BranchLength)
		}
		if res.Format != v.cfg.f {
			t.Errorf("case %d: format expected %q, but was %q", i, v.cfg.f, res.Format)
		}
		if err := diffMaps(v.cfg.op, res.Options); err != nil {
			t.Errorf("case %d: %s", i, err.Error())
		}
	}
}

func TestInitConfigFail(t *testing.T) {
	cases := []caseE{
		caseE{
			[]string{"-x"},
			"flag provided but not defined: -x",
		},
		caseE{
			[]string{"-s", "1x2x3"},
			"size must be in form `<WIDTH>x<HEIGHT>`,  like `30x10`",
		},
		caseE{
			[]string{"-s-2x3"},
			"<WIDTH> must be positive integer in `-s` (`--size`)",
		},
		caseE{
			[]string{"-s2x-3"},
			"<HEIGHT> must be positive integer in `-s` (`--size`)",
		},
		caseE{
			[]string{"-W-2"},
			"<WIDTH> must be greater then zero in `-W` (`--width`)",
		},
		caseE{
			[]string{"-H-3"},
			"<HEIGHT> must be greater then zero in `-H` (`--height`)",
		},
		caseE{
			[]string{"-Bfail"},
			"<BL> must be either integer greater then 1 (number of CELLs), " +
				"a string `max` to set <WIDTH>*<HEIGHT>, " +
				"or decimal from 0 to 1 as fraction of max in `-B` (`--branch-length`)",
		},
		caseE{
			[]string{"-flul"},
			"unknown format name in `-F` (`--format`)",
		},
	}
	for i, v := range cases {
		_, e := initConfig("cmd", v.args)
		if e == nil || e.Error() != v.err {
			t.Errorf("case %d: expected error %v but saw %v", i, v.err, e)
		}
	}
}

func TestParseBranchLength(t *testing.T) {
	for s, i := range map[string]int{
		"max":    100,
		"1":      1,
		"100":    100,
		"101":    100,
		"0.01":   1,
		"1.0":    100,
		"0.2249": 22,
		"0.225":  23,
	} {
		if n, e := parseBranchLength(s, 100); e != nil || n != i {
			t.Errorf("case %q: failed to parse correct value: n=%d; e=%v", s, n, e)
		}
	}
	msg := "must be either integer greater then 1 (number of CELLs), " +
		"a string `max` to set <WIDTH>*<HEIGHT>, " +
		"or decimal from 0 to 1 as fraction of max"
	if n, e := parseBranchLength("1.foo", 100); e == nil || e.Error() != msg {
		t.Errorf("unexpected result: n=%d; e=%v", n, e)
	}
}

func TestParsePlusIntValue(t *testing.T) {
	for s, i := range map[string]int{"1": 1, "65535": 65535} {
		if n, e := parsePlusIntValue(s); e != nil || n != i {
			t.Errorf("case %q: failed to parse correct value: n=%d; e=%v", s, n, e)
		}
	}
	if n, e := parsePlusIntValue("0"); e == nil || e.Error() != "must be greater then zero" {
		t.Errorf("unexpected result: n=%d; e=%v", n, e)
	}
	if n, e := parsePlusIntValue("10h"); e == nil || e.Error() != "must be positive integer" {
		t.Errorf("unexpected result: n=%d; e=%v", n, e)
	}
	if n, e := parsePlusIntValue("65536"); e == nil || e.Error() != "is out of range" {
		t.Errorf("unexpected result: n=%d; e=%v", n, e)
	}
}

func TestValidatePlusIntValue(t *testing.T) {
	for _, n := range []int{1, 2, 0xFFFE, 0xFFFF} {
		if e := validatePlusIntValue(n); e != nil {
			t.Errorf("unexpected error for %d: %v", n, e)
		}
	}
	for _, n := range []int{-2, -1, 0} {
		e := validatePlusIntValue(n)
		if e == nil || e.Error() != "must be greater then zero" {
			t.Errorf("unexpected error for %d: %v", n, e)
		}
	}
	for _, n := range []int{0x10000, 0x10001} {
		e := validatePlusIntValue(n)
		if e == nil || e.Error() != "is too big" {
			t.Errorf("unexpected error for %d: %v", n, e)
		}
	}
}

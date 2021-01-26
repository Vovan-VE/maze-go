package cli

import (
	"flag"
	"testing"
)

func TestIsFlagSetStd(t *testing.T) {
	fs := flag.NewFlagSet("foo", flag.PanicOnError)
	fs.Int("width", 10, "width of rect")
	fs.String("name", "", "name hint here")

	fs.Parse([]string{""})
	if IsFlagSet(fs, "width") {
		t.Error("width did not set yet")
	}
	if IsFlagSet(fs, "name") {
		t.Error("name did not set yet")
	}

	fs.Parse([]string{"--width", "42"})
	if !IsFlagSet(fs, "width") {
		t.Error("width must be set")
	}
	if IsFlagSet(fs, "name") {
		t.Error("name did not set yet")
	}

	if IsFlagSet(fs, "bad") {
		t.Error("unknown flag cannot be set")
	}
}

// func TestIsFlagSetGetopt(t *testing.T) {
// 	fs := getopt.NewFlagSet("foo", flag.PanicOnError)
// 	fs.Int("width", 10, "width of rect")
// 	fs.String("name", "", "name hint here")
// 	fs.Alias("w", "width")
// 	fs.Alias("n", "name")
//
// 	//fs.Parse([]string{""})
// 	if IsFlagSet(fs, "width") || IsFlagSet(fs, "w") {
// 		t.Error("width did not set yet")
// 	}
// 	if IsFlagSet(fs, "name") || IsFlagSet(fs, "n") {
// 		t.Error("name did not set yet")
// 	}
//
// 	e := fs.Parse([]string{"--width", "tt", "rest", "args"})
// 	if e != nil {
// 		t.Errorf("Parse failed: %v", e)
// 	}
// 	if !IsFlagSet(fs, "width") {
// 		t.Errorf("width must be set, %v, %v", fs, fs.FlagSet)
// 	}
// 	if !IsFlagSet(fs, "w") {
// 		t.Error("w must be set")
// 	}
// 	if IsFlagSet(fs, "name") {
// 		t.Error("name did not set yet")
// 	}
// 	if IsFlagSet(fs, "n") {
// 		t.Error("name did not set yet")
// 	}
//
// 	if IsFlagSet(fs, "bad") {
// 		t.Error("unknown flag cannot be set")
// 	}
// }

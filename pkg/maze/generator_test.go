package maze

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/Vovan-VE/maze-go/pkg/maze/format/text"
)

func TestNewGenerator(t *testing.T) {
	w, h := 10, 10
	config := NewConfig()
	config.Width = w
	config.Height = h
	config.BranchLength = 2
	gen := NewGenerator(config)
	maze := gen.Generate()
	exporter := text.NewExporter()
	result := exporter.ExportMaze(maze)

	if n := strings.Count(result, "i"); n != 1 {
		t.Errorf("Found entrances: %d", n)
	}
	if n := strings.Count(result, "E"); n != 1 {
		t.Errorf("Found exits: %d", n)
	}

	lines := strings.Split(result, "\n")
	if n := len(lines); n != h*2+1 {
		t.Errorf("There are %d lines, but %d expected", n, h*2+1)
	}
	for y, line := range lines {
		if n := len(line); n != w*2+1 {
			t.Errorf("line %d is %d length, but %d expected", y, n, w*2+1)
		}
	}

	// ####### edge
	// # # # # space
	// ####### mid
	// # # # # space
	// ####### edge
	reEdge := regexp.MustCompile(fmt.Sprintf("^#([#iE]#){%d}$", w))
	reMid := regexp.MustCompile(fmt.Sprintf("^#([ #]#){%d}$", w))
	reSpace := regexp.MustCompile(fmt.Sprintf("^[#iE]( [ #]){%d} [#iE]$", w-1))
	for y := 0; y < h; y++ {
		re := reMid
		if y == 0 {
			re = reEdge
		}
		if !re.MatchString(lines[y*2]) {
			t.Errorf("A `wall` line %d (y=%d) is %q", y*2, y, lines[y*2])
		}
		if !reSpace.MatchString(lines[y*2+1]) {
			t.Errorf("A `space` line %d (y=%d) is %q", y*2+1, y, lines[y*2+1])
		}
	}
	if !reEdge.MatchString(lines[h*2]) {
		t.Errorf("A bottom `wall` line %d is %q", h*2, lines[h*2])
	}

}

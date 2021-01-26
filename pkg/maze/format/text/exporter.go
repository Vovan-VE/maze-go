package text

import (
	"strings"
	"unicode/utf8"

	"github.com/Vovan-VE/maze-go/pkg/maze/data"
)

// Exporter implements MazeExporter for text format
type Exporter struct {
	*baseConfig
}

// NewExporter creates new Exporter
func NewExporter() *Exporter {
	return &Exporter{newBaseConfig()}
}

// ConfigureExport configures an instance with given values
func (c *Exporter) ConfigureExport(config map[string]string) {
	c.configure(config)
}

// ExportMaze does export given Maze in appropriate format
func (c *Exporter) ExportMaze(maze *data.Maze) string {
	w := maze.Width()
	h := maze.Height()

	charsCount := utf8.RuneCountInString(c.wall)
	space := strings.Repeat(" ", charsCount)

	lines := make([][]string, h*2+1)
	for i := range lines {
		line := make([]string, w*2+1)
		for j := range line {
			line[j] = c.wall
		}
		lines[i] = line
	}

	for y := 0; y < h; y++ {
		sy := y*2 + 1
		for x := 0; x < w; x++ {
			sx := x*2 + 1
			lines[sy][sx] = space

			cell := maze.Cell(x, y)
			if y == 0 && !cell.TopWall {
				lines[sy-1][sx] = space
			}
			if x == 0 && !cell.LeftWall {
				lines[sy][sx-1] = space
			}
			if !cell.BottomWall {
				lines[sy+1][sx] = space
			}
			if !cell.RightWall {
				lines[sy][sx+1] = space
			}
		}
	}

	if in := maze.Entrance(); in != nil {
		markDoor(lines, in, repeatStringToLength(c.in, charsCount))
	}
	if out := maze.Exit(); out != nil {
		markDoor(lines, out, repeatStringToLength(c.out, charsCount))
	}

	str := make([]string, len(lines))
	for y, line := range lines {
		str[y] = strings.Join(line, "")
	}

	return strings.Join(str, "\n")
}

func markDoor(lines [][]string, door *data.DoorPosition, str string) {
	var x, y int
	switch door.Side() {
	case data.TOP:
		x = door.Offset()*2 + 1
		y = 0
	case data.RIGHT:
		x = len(lines[0]) - 1
		y = door.Offset()*2 + 1
	case data.BOTTOM:
		x = door.Offset()*2 + 1
		y = len(lines) - 1
	case data.LEFT:
		x = 0
		y = door.Offset()*2 + 1
	}
	lines[y][x] = str
}

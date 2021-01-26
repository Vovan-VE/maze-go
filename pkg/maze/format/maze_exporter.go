package format

import (
	"github.com/Vovan-VE/maze-go/pkg/maze/data"
)

// MazeExporter allows to export a maze to specific format
type MazeExporter interface {
	// ConfigureExport configures an instance with given values
	ConfigureExport(config map[string]string)
	// ExportMaze does export given Maze in appropriate format
	ExportMaze(maze *data.Maze) string
}

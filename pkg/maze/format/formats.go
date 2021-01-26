package format

import (
	"github.com/Vovan-VE/maze-go/pkg/maze/format/text"
)

type ctor func() MazeExporter

var exporters = map[string]ctor{
	"text": func() MazeExporter {
		return text.NewExporter()
	},
}

// HasExporter checks if MazeExporter exists for given format
func HasExporter(format string) bool {
	_, ok := exporters[format]
	return ok
}

// NewExporter creates new MazeExporter for given format
func NewExporter(format string) MazeExporter {
	create, ok := exporters[format]
	if !ok {
		panic("Unknown format name: " + format)
	}
	return create()
}

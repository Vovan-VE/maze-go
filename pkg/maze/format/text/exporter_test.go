package text

import (
	"strings"
	"testing"

	"github.com/Vovan-VE/maze-go/pkg/maze/data"
)

func TestConfigureExport(t *testing.T) {
	maze := data.NewMaze(1, 1)
	maze.SetEntrance(data.LEFT, 0)
	maze.SetExit(data.BOTTOM, 0)
	exporter := NewExporter()
	exporter.ConfigureExport(map[string]string{
		"wall": "▒▒",
		"in":   "()",
		"out":  "[]",
	})
	expected := strings.Join([]string{
		"▒▒▒▒▒▒",
		"()  ▒▒",
		"▒▒[]▒▒",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}
}

func TestExportAllWalls(t *testing.T) {
	maze := data.NewMaze(3, 3)
	exporter := NewExporter()
	expected := strings.Join([]string{
		"#######",
		"# # # #",
		"#######",
		"# # # #",
		"#######",
		"# # # #",
		"#######",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}

	exporter.wall = "▒"
	expected = strings.Join([]string{
		"▒▒▒▒▒▒▒",
		"▒ ▒ ▒ ▒",
		"▒▒▒▒▒▒▒",
		"▒ ▒ ▒ ▒",
		"▒▒▒▒▒▒▒",
		"▒ ▒ ▒ ▒",
		"▒▒▒▒▒▒▒",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}
}

func TestExportMultichar(t *testing.T) {
	maze := data.NewMaze(3, 3)
	maze.SetEntrance(data.LEFT, 0)
	maze.SetExit(data.RIGHT, 2)

	exporter := NewExporter()
	exporter.wall = "▓█▓"
	exporter.in = "()"

	expected := strings.Join([]string{
		"▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓",
		"()(   ▓█▓   ▓█▓   ▓█▓",
		"▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓",
		"▓█▓   ▓█▓   ▓█▓   ▓█▓",
		"▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓",
		"▓█▓   ▓█▓   ▓█▓   EEE",
		"▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓▓█▓",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}
}

func TestExport(t *testing.T) {
	maze := data.NewMaze(3, 3)
	exporter := NewExporter()

	// remove some walls to be like a real maze
	expected := strings.Join([]string{
		"#######",
		"#   # #",
		"### # #",
		"# #   #",
		"# # ###",
		"#     #",
		"#######",
	}, "\n")
	maze.RemoveWalls(0, 0, data.RIGHT)
	maze.RemoveWalls(1, 0, data.BOTTOM)
	maze.RemoveWalls(2, 0, data.BOTTOM)
	maze.RemoveWalls(1, 1, data.RIGHT)
	maze.RemoveWalls(1, 1, data.BOTTOM)
	maze.RemoveWalls(0, 1, data.BOTTOM)
	maze.RemoveWalls(0, 2, data.RIGHT)
	maze.RemoveWalls(1, 2, data.RIGHT)
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}

	// remove some outer walls to test all its possible cases
	expected = strings.Join([]string{
		"# ### #",
		"    #  ",
		"### # #",
		"# #   #",
		"# # ###",
		"       ",
		"# ### #",
	}, "\n")
	maze.RemoveOuterWall(0, 0, data.TOP)
	maze.RemoveOuterWall(0, 0, data.LEFT)
	maze.RemoveOuterWall(2, 0, data.TOP)
	maze.RemoveOuterWall(2, 0, data.RIGHT)
	maze.RemoveOuterWall(0, 2, data.BOTTOM)
	maze.RemoveOuterWall(0, 2, data.LEFT)
	maze.RemoveOuterWall(2, 2, data.RIGHT)
	maze.RemoveOuterWall(2, 2, data.BOTTOM)
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}

	// remove all the rest walls to check
	expected = strings.Join([]string{
		"# # # #",
		"       ",
		"# # # #",
		"       ",
		"# # # #",
		"       ",
		"# # # #",
	}, "\n")
	maze.RemoveWalls(0, 0, data.BOTTOM)
	maze.RemoveWalls(1, 0, data.RIGHT)
	maze.RemoveWalls(0, 1, data.RIGHT)
	maze.RemoveWalls(2, 1, data.BOTTOM)
	maze.RemoveOuterWall(1, 0, data.TOP)
	maze.RemoveOuterWall(0, 1, data.LEFT)
	maze.RemoveOuterWall(2, 1, data.RIGHT)
	maze.RemoveOuterWall(1, 2, data.BOTTOM)
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}
}

func TestExportDoors(t *testing.T) {
	exporter := NewExporter()

	maze := data.NewMaze(3, 2)
	maze.SetEntrance(data.LEFT, 0)
	maze.SetExit(data.RIGHT, 0)
	expected := strings.Join([]string{
		"#######",
		"i # # E",
		"#######",
		"# # # #",
		"#######",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}

	maze = data.NewMaze(3, 2)
	maze.SetEntrance(data.LEFT, 1)
	maze.SetExit(data.RIGHT, 1)
	expected = strings.Join([]string{
		"#######",
		"# # # #",
		"#######",
		"i # # E",
		"#######",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}

	maze = data.NewMaze(3, 2)
	maze.SetEntrance(data.TOP, 0)
	maze.SetExit(data.BOTTOM, 0)
	expected = strings.Join([]string{
		"#i#####",
		"# # # #",
		"#######",
		"# # # #",
		"#E#####",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}

	maze = data.NewMaze(3, 2)
	maze.SetEntrance(data.TOP, 2)
	maze.SetExit(data.BOTTOM, 2)
	expected = strings.Join([]string{
		"#####i#",
		"# # # #",
		"#######",
		"# # # #",
		"#####E#",
	}, "\n")
	if result := exporter.ExportMaze(maze); result != expected {
		t.Errorf("Expected <<<<\n%s\n>>>>\nbut saw <<<<\n%s\n>>>>", expected, result)
	}
}

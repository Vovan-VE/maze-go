package maze

import (
	"math/rand"

	"github.com/Vovan-VE/maze-go/pkg/maze/data"
)

// Generator can generate randomized Maze
type Generator struct {
	config *Config
}

// NewGenerator creates a new Generator with specified Config
func NewGenerator(config *Config) *Generator {
	return &Generator{config}
}

// Generate creates new randomized Maze
func (g *Generator) Generate() *data.Maze {
	free := data.NewCellsSet()
	maze := data.NewMaze(g.config.Width, g.config.Height)
	for _, cell := range maze.AllCells() {
		free.Add(cell)
	}

	length := g.config.BranchLength
	frontier := data.NewCellsSet()

	start := free.Random()
	free.Remove(start)
	frontier.Add(start)

	for !free.IsEmpty() {
		current := frontier.Random()
		for L := 0; L < length; L++ {
			next, at := findAdjacentFree(maze, free, current, data.RandomDirection())
			maze.RemoveWalls(current.X(), current.Y(), at)
			free.Remove(next)
			if free.IsEmpty() {
				break
			}
			updateFrontiers(maze, frontier, free, next)
			if !cellHasFreeAdjacent(maze, free, next) {
				break
			}
			frontier.Add(next)
			current = next
		}
	}

	maze.SetEntrance(data.LEFT, rand.Intn(maze.Height()))
	maze.SetExit(data.RIGHT, rand.Intn(maze.Height()))

	return maze
}

func findAdjacentFree(
	maze *data.Maze,
	free *data.CellsSet,
	current *data.Cell,
	at data.Direction,
) (next *data.Cell, to data.Direction) {
	x, y := current.XY()
	for n := 0; n < 4; n++ {
		next := maze.AdjacentCell(x, y, at)
		if next != nil && free.Has(next) {
			return next, at
		}
		at = at.Next()
	}
	panic("There are no free cells ever?")
}

func updateFrontiers(maze *data.Maze, frontier, free *data.CellsSet, cell *data.Cell) {
	x, y := cell.XY()
	d := data.TOP
	for n := 0; n < 4; n++ {
		next := maze.AdjacentCell(x, y, d)
		if next != nil && frontier.Has(next) && !cellHasFreeAdjacent(maze, free, next) {
			frontier.Remove(next)
		}
		d = d.Next()
	}
}

func cellHasFreeAdjacent(maze *data.Maze, free *data.CellsSet, cell *data.Cell) bool {
	x, y := cell.XY()
	d := data.TOP
	for n := 0; n < 4; n++ {
		next := maze.AdjacentCell(x, y, d)
		if next != nil && free.Has(next) {
			return true
		}
		d = d.Next()
	}
	return false
}

package data

import (
	"fmt"
)

type cellsBase struct {
	cells map[string]*Cell
}

func newCellsBase() *cellsBase {
	return &cellsBase{make(map[string]*Cell)}
}

// IsEmpty returns whether a set is empty
func (s *cellsBase) IsEmpty() bool {
	return 0 == len(s.cells)
}

func (s *cellsBase) Size() int {
	return len(s.cells)
}

// Has checks if a Cell is included
func (s *cellsBase) Has(c *Cell) bool {
	return s.HasKey(cellKey(c))
}

// HasKey checks if a Cell is included by its key
func (s *cellsBase) HasKey(key string) bool {
	_, ok := s.cells[key]
	return ok
}

func cellKey(c *Cell) string {
	return fmt.Sprintf("%d;%d", c.X(), c.Y())
}

package data

import (
	"math/rand"
)

// CellsSet represents a set of Cell's with coordinate uniqueness
type CellsSet struct {
	*cellsBase
}

// NewCellsSet creates empty set
func NewCellsSet() *CellsSet {
	return &CellsSet{newCellsBase()}
}

// Add does add a Cell into a Set
func (s *CellsSet) Add(c *Cell) {
	k := cellKey(c)
	if s.HasKey(k) {
		if s.cells[k] != c {
			panic("trying to add different cell with the same coords")
		}
	} else {
		s.cells[k] = c
	}
}

// Remove does remove a Cell from set
func (s *CellsSet) Remove(c *Cell) {
	delete(s.cells, cellKey(c))
}

// Random returns a random Cell from a set
func (s *CellsSet) Random() *Cell {
	if s.IsEmpty() {
		panic("The Set is empty")
	}

	// since this operation is not much offten,
	// and a number of underlying cells shouldn't be so much
	cells := make([]*Cell, len(s.cells))
	i := 0
	for _, c := range s.cells {
		cells[i] = c
		i++
	}
	return cells[rand.Intn(len(cells))]
}

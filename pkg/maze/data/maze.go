package data

import (
	"fmt"
)

// Maze describes a maze
type Maze struct {
	width, height int
	cells         []*Cell
	in, out       *DoorPosition
}

// NewMaze prepares maze as a matrix of cells with all walls set
func NewMaze(width, height int) *Maze {
	length := width * height
	cells := make([]*Cell, length)
	for i := 0; i < length; i++ {
		x := i % width
		y := i / width
		cells[i] = NewCell(x, y)
	}
	return &Maze{width, height, cells, nil, nil}
}

// Width returns a width of maze
func (m *Maze) Width() int {
	return m.width
}

// Height returns a height of maze
func (m *Maze) Height() int {
	return m.height
}

// Entrance returns position of maze entrance or nil
func (m *Maze) Entrance() *DoorPosition {
	return m.in
}

// Exit returns position of maze exit or nil
func (m *Maze) Exit() *DoorPosition {
	return m.out
}

// EntranceCell returns a Cell near maze entrance
func (m *Maze) EntranceCell() *Cell {
	return m.doorCell(m.in)
}

// ExitCell returns a Cell near to maze Exit
func (m *Maze) ExitCell() *Cell {
	return m.doorCell(m.out)
}

// Cell returs a Cell by its coordinates
func (m *Maze) Cell(x, y int) *Cell {
	c, ok := m.cell(x, y)
	if !ok {
		panic("Coords out of range")
	}
	return c
}

// AllCells returns a slice to iterate over all maze cells
func (m *Maze) AllCells() []*Cell {
	return m.cells[:]
}

// AdjacentCell returns a cell adjacent to given coordinates at given position or nil
func (m *Maze) AdjacentCell(x, y int, at Direction) *Cell {
	ax, ay := at.AdjacentCoords(x, y)
	cell, _ := m.cell(ax, ay)
	return cell
}

// SetEntrance sets maze entrance and removes corresponding outer wall
func (m *Maze) SetEntrance(side Direction, offset int) {
	in := NewDoorPosition(side, offset)
	if m.out != nil && *m.out == *in {
		panic("This place is already assigned to Exit")
	}
	x, y := m.doorCoords(side, offset)
	m.RemoveOuterWall(x, y, side)
	m.in = in
}

// SetExit sets maze exit and removes corresponding outer wall
func (m *Maze) SetExit(side Direction, offset int) {
	out := NewDoorPosition(side, offset)
	if m.in != nil && *m.in == *out {
		panic("This place is already assigned to Entrance")
	}
	x, y := m.doorCoords(side, offset)
	m.RemoveOuterWall(x, y, side)
	m.out = out
}

// RemoveWalls - remove walls between adjacent cells
func (m *Maze) RemoveWalls(x, y int, side Direction) {
	current := m.Cell(x, y)
	adjacent := m.AdjacentCell(x, y, side)
	if nil == adjacent {
		panic(fmt.Sprintf("There is no adjacent cell for [%d; %d] at %d", x, y, int(side)))
	}
	current.SetWallAt(side, false)
	adjacent.SetWallAt(side.Opposite(), false)
}

// RemoveOuterWall - remove outer wall
func (m *Maze) RemoveOuterWall(x, y int, side Direction) {
	current := m.Cell(x, y)
	if m.AdjacentCell(x, y, side) != nil {
		panic("Target cell is not edge cell")
	}
	current.SetWallAt(side, false)
}

func (m *Maze) isValidCoords(x, y int) bool {
	return x >= 0 && y >= 0 && x < m.width && y < m.height
}

func (m *Maze) cell(x, y int) (cell *Cell, ok bool) {
	if !m.isValidCoords(x, y) {
		return nil, false
	}
	return m.cells[y*m.width+x], true
}

func (m *Maze) doorCell(door *DoorPosition) *Cell {
	if door == nil {
		return nil
	}
	x, y := m.doorCoords(door.Side(), door.Offset())
	return m.Cell(x, y)
}

func (m *Maze) doorCoords(side Direction, offset int) (x, y int) {
	switch side {
	case TOP:
		return offset, 0
	case RIGHT:
		return m.width - 1, offset
	case BOTTOM:
		return offset, m.height - 1
	case LEFT:
		return 0, offset
	}
	panic("Invalid direction")
}

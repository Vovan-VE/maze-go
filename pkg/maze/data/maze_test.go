package data

import (
	"testing"
)

type point struct {
	x, y int
}

func TestNewMaze(t *testing.T) {
	m := NewMaze(4, 3)
	if m.Width() != 4 {
		t.Error("Wrong width")
	}
	if m.Height() != 3 {
		t.Error("Wrong height")
	}

	expectedCoords := []point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
	}
	for i, cell := range m.AllCells() {
		x, y := cell.XY()
		ec := expectedCoords[i]
		if x != ec.x || y != ec.y {
			t.Errorf("cell[%d]: expected %v, but saw [%d, %d]", i, ec, x, y)
		}
		if cell != m.Cell(x, y) {
			t.Error("a Cell() returned a different cell")
		}
		if !cell.TopWall || !cell.RightWall || !cell.BottomWall || !cell.LeftWall {
			t.Errorf("cell[%d]: not all walls are set", i)
		}
	}
}

func TestCellOutOfRange(t *testing.T) {
	m := NewMaze(4, 3)
	bad := []point{
		{0, -1},
		{1, -1},
		{2, -1},
		{3, -1},
		{4, -1},
		{4, 0},
		{4, 1},
		{4, 2},
		{4, 3},
		{3, 3},
		{2, 3},
		{1, 3},
		{0, 3},
		{-1, 3},
		{-1, 2},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}
	for _, p := range bad {
		func() {
			defer func() {
				e := recover()
				if e == nil {
					t.Errorf("Did not panic with %v", p)
				}
				if e != "Coords out of range" {
					panic(e)
				}
			}()
			_ = m.Cell(p.x, p.y)
		}()
	}
}

func TestSetEntrance(t *testing.T) {
	m := NewMaze(3, 2)
	if m.Entrance() != nil {
		t.Error("Unexpected initial entrance")
	}

	m.SetEntrance(TOP, 0)
	if *m.Entrance() != *NewDoorPosition(TOP, 0) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(0, 0).TopWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(TOP, 2)
	if *m.Entrance() != *NewDoorPosition(TOP, 2) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(2, 0).TopWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(RIGHT, 0)
	if *m.Entrance() != *NewDoorPosition(RIGHT, 0) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(2, 0).RightWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(RIGHT, 1)
	if *m.Entrance() != *NewDoorPosition(RIGHT, 1) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(2, 1).RightWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(BOTTOM, 2)
	if *m.Entrance() != *NewDoorPosition(BOTTOM, 2) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(2, 1).BottomWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(BOTTOM, 0)
	if *m.Entrance() != *NewDoorPosition(BOTTOM, 0) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(0, 1).BottomWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(LEFT, 0)
	if *m.Entrance() != *NewDoorPosition(LEFT, 0) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(0, 0).LeftWall {
		t.Error("Outer wall was not removed")
	}

	m.SetEntrance(LEFT, 1)
	if *m.Entrance() != *NewDoorPosition(LEFT, 1) {
		t.Error("Entrance was set but it's mismatch")
	}
	if m.Cell(0, 1).LeftWall {
		t.Error("Outer wall was not removed")
	}
}

func TestSetEntranceFail(t *testing.T) {
	m := NewMaze(3, 2)
	m.SetExit(TOP, 1)
	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "This place is already assigned to Exit" {
			panic(e)
		}
	}()
	m.SetEntrance(TOP, 1)
}

func TestSetExit(t *testing.T) {
	m := NewMaze(3, 2)
	if m.Exit() != nil {
		t.Error("Unexpected initial entrance")
	}

	m.SetExit(TOP, 0)
	if *m.Exit() != *NewDoorPosition(TOP, 0) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(0, 0).TopWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(TOP, 2)
	if *m.Exit() != *NewDoorPosition(TOP, 2) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(2, 0).TopWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(RIGHT, 0)
	if *m.Exit() != *NewDoorPosition(RIGHT, 0) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(2, 0).RightWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(RIGHT, 1)
	if *m.Exit() != *NewDoorPosition(RIGHT, 1) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(2, 1).RightWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(BOTTOM, 2)
	if *m.Exit() != *NewDoorPosition(BOTTOM, 2) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(2, 1).BottomWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(BOTTOM, 0)
	if *m.Exit() != *NewDoorPosition(BOTTOM, 0) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(0, 1).BottomWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(LEFT, 0)
	if *m.Exit() != *NewDoorPosition(LEFT, 0) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(0, 0).LeftWall {
		t.Error("Outer wall was not removed")
	}

	m.SetExit(LEFT, 1)
	if *m.Exit() != *NewDoorPosition(LEFT, 1) {
		t.Error("Exit was set but it's mismatch")
	}
	if m.Cell(0, 1).LeftWall {
		t.Error("Outer wall was not removed")
	}
}

func TestSetExitFail(t *testing.T) {
	m := NewMaze(3, 2)
	m.SetEntrance(TOP, 1)
	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "This place is already assigned to Entrance" {
			panic(e)
		}
	}()
	m.SetExit(TOP, 1)
}

func TestEntranceCell(t *testing.T) {
	m := NewMaze(3, 2)
	if m.EntranceCell() != nil {
		t.Error("Unexpected entrance cell")
	}
	m.SetEntrance(BOTTOM, 2)
	c := m.EntranceCell()
	if *c != *m.Cell(2, 1) {
		t.Errorf("Wrong cell returned: %v", *c)
	}
}

func TestExitCell(t *testing.T) {
	m := NewMaze(3, 2)
	if m.ExitCell() != nil {
		t.Error("Unexpected exit cell")
	}
	m.SetExit(BOTTOM, 2)
	c := m.ExitCell()
	if *c != *m.Cell(2, 1) {
		t.Errorf("Wrong cell returned: %v", *c)
	}
}

func TestAdjacentCell(t *testing.T) {
	m := NewMaze(4, 3)

	dir := []Direction{TOP, RIGHT, BOTTOM, LEFT}
	cases := []struct {
		at point
		to [4]*point
	}{
		{point{0, 0}, [4]*point{nil, &point{1, 0}, &point{0, 1}, nil}},
		{point{1, 0}, [4]*point{nil, &point{2, 0}, &point{1, 1}, &point{0, 0}}},
		{point{3, 0}, [4]*point{nil, nil, &point{3, 1}, &point{2, 0}}},
		{point{0, 1}, [4]*point{&point{0, 0}, &point{1, 1}, &point{0, 2}, nil}},
		{point{0, 2}, [4]*point{&point{0, 1}, &point{1, 2}, nil, nil}},
		{point{1, 2}, [4]*point{&point{1, 1}, &point{2, 2}, nil, &point{0, 2}}},
		{point{3, 2}, [4]*point{&point{3, 1}, nil, nil, &point{2, 2}}},
	}

	for n, item := range cases {
		at := item.at
		for i, p := range item.to {
			var target *Cell
			if p != nil {
				target = m.Cell(p.x, p.y)
			}
			if c := m.AdjacentCell(at.x, at.y, dir[i]); c != target {
				t.Errorf("case %d: expected %v, but saw %v", n, target, *c)
			}
		}
	}
}

func TestRemoveWalls(t *testing.T) {
	m := NewMaze(3, 2)
	a := m.Cell(0, 0)
	b := m.Cell(0, 1)

	m.RemoveWalls(0, 0, BOTTOM)
	if a.BottomWall || b.TopWall {
		t.Error("Did not remove requested walls")
	}
	if !a.TopWall || !a.RightWall || !a.LeftWall {
		t.Error("Must keep untouched walls")
	}
	if !b.BottomWall || !b.RightWall || !b.LeftWall {
		t.Error("Must keep untouched walls")
	}

	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "There is no adjacent cell for [2; 0] at 1" {
			panic(e)
		}
	}()
	m.RemoveWalls(2, 0, TOP)
}

func TestRemoveOuterWall(t *testing.T) {
	m := NewMaze(3, 2)
	c := m.Cell(2, 0)

	m.RemoveOuterWall(2, 0, RIGHT)
	if c.RightWall {
		t.Error("Did not remove requested wall")
	}
	if !c.TopWall || !c.LeftWall || !c.BottomWall {
		t.Error("Must keep untouched walls")
	}

	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "Target cell is not edge cell" {
			panic(e)
		}
	}()
	m.RemoveOuterWall(2, 0, LEFT)
}

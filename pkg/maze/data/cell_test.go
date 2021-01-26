package data

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	c := NewCell(42, 37)
	if c.X() != 42 {
		t.Error("Wrong X")
	}
	if c.Y() != 37 {
		t.Error("Wrong Y")
	}

	x, y := c.XY()
	if x != 42 || y != 37 {
		t.Error("Wrong XY")
	}

	if !c.TopWall {
		t.Error("TopWall was not set")
	}
	if !c.RightWall {
		t.Error("RightWall was not set")
	}
	if !c.BottomWall {
		t.Error("BottomWall was not set")
	}
	if !c.LeftWall {
		t.Error("LeftWall was not set")
	}
}

func TestNewCellOpened(t *testing.T) {
	c := NewCellOpened(42, 37)
	if c.X() != 42 {
		t.Error("Wrong X")
	}
	if c.Y() != 37 {
		t.Error("Wrong Y")
	}
	if c.TopWall {
		t.Error("TopWall was set")
	}
	if c.RightWall {
		t.Error("RightWall was set")
	}
	if c.BottomWall {
		t.Error("BottomWall was set")
	}
	if c.LeftWall {
		t.Error("LeftWall was set")
	}
}

func TestIsSameCoords(t *testing.T) {
	a := NewCell(0, 0)
	if !a.IsSameCoords(a) {
		t.Error("Same cell for same Cell")
	}
	if !a.IsSameCoords(NewCell(0, 0)) {
		t.Error("Same cell for similar Cell")
	}

	if a.IsSameCoords(NewCell(0, 1)) {
		t.Error("Wrong for different coords")
	}
	if a.IsSameCoords(NewCell(1, 0)) {
		t.Error("Wrong for different coords")
	}
	if a.IsSameCoords(NewCell(1, 1)) {
		t.Error("Wrong for different coords")
	}
}

func TestHasWallAt(t *testing.T) {
	c := NewCell(0, 0)
	if !c.HasWallAt(TOP) {
		t.Error("No wall at TOP")
	}
	if !c.HasWallAt(RIGHT) {
		t.Error("No wall at RIGHT")
	}
	if !c.HasWallAt(BOTTOM) {
		t.Error("No wall at BOTTOM")
	}
	if !c.HasWallAt(LEFT) {
		t.Error("No wall at LEFT")
	}

	c = NewCellOpened(0, 0)
	if c.HasWallAt(TOP) {
		t.Error("Has wall at TOP")
	}
	if c.HasWallAt(RIGHT) {
		t.Error("Has wall at RIGHT")
	}
	if c.HasWallAt(BOTTOM) {
		t.Error("Has wall at BOTTOM")
	}
	if c.HasWallAt(LEFT) {
		t.Error("Has wall at LEFT")
	}
}

func TestSetWallAt(t *testing.T) {
	c := NewCell(0, 0)
	if !c.TopWall || !c.RightWall || !c.BottomWall || !c.LeftWall {
		t.Error("T,R,B,L walls must be set")
	}

	c.SetWallAt(TOP, false)
	if c.TopWall || !c.RightWall || !c.BottomWall || !c.LeftWall {
		t.Error("Expected: no TOP wall and set R,B,L walls")
	}

	c.SetWallAt(RIGHT, false)
	if c.TopWall || c.RightWall || !c.BottomWall || !c.LeftWall {
		t.Error("Expected: no T,R walls and set B,L walls")
	}

	c.SetWallAt(BOTTOM, false)
	if c.TopWall || c.RightWall || c.BottomWall || !c.LeftWall {
		t.Error("Expected: no T,R,B walls and set LEFT wall")
	}

	c.SetWallAt(LEFT, false)
	if c.TopWall || c.RightWall || c.BottomWall || c.LeftWall {
		t.Error("Expected: no T,R,B,L walls")
	}

	c.SetWallAt(TOP, true)
	if !c.TopWall || c.RightWall || c.BottomWall || c.LeftWall {
		t.Error("Expected: set TOP wall and no R,B,L walls")
	}

	c.SetWallAt(RIGHT, true)
	if !c.TopWall || !c.RightWall || c.BottomWall || c.LeftWall {
		t.Error("Expected: set T,R wall and no B,L walls")
	}

	c.SetWallAt(BOTTOM, true)
	if !c.TopWall || !c.RightWall || !c.BottomWall || c.LeftWall {
		t.Error("Expected: set T,R,B wall and no LEFT wall")
	}

	c.SetWallAt(LEFT, true)
	if !c.TopWall || !c.RightWall || !c.BottomWall || !c.LeftWall {
		t.Error("Expected: set T,R,B,L wall")
	}
}

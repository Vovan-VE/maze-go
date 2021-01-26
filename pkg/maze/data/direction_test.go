package data

import (
	"testing"
)

func TestRandomDirection(t *testing.T) {
	all := map[Direction]bool{
		TOP:    true,
		RIGHT:  true,
		BOTTOM: true,
		LEFT:   true,
	}

	for i := 1; i <= 100 && len(all) > 0; i++ {
		d := RandomDirection()
		delete(all, d)
	}

	if len(all) > 0 {
		t.Error("Some directions never returned from `RandomDirection()`")
	}
}

func TestPrev(t *testing.T) {
	if LEFT != TOP.Prev() {
		t.Error("Prev to TOP must be LEFT")
	}
	if TOP != RIGHT.Prev() {
		t.Error("Prev to RIGHT must be TOP")
	}
	if RIGHT != BOTTOM.Prev() {
		t.Error("Prev to BOTTOM must be RIGHT")
	}
	if BOTTOM != LEFT.Prev() {
		t.Error("Prev to LEFT must be BOTTOM")
	}
}

func TestNext(t *testing.T) {
	if RIGHT != TOP.Next() {
		t.Error("Next to TOP must be RIGHT")
	}
	if BOTTOM != RIGHT.Next() {
		t.Error("Next to RIGHT must be BOTTOM")
	}
	if LEFT != BOTTOM.Next() {
		t.Error("Next to BOTTOM must be LEFT")
	}
	if TOP != LEFT.Next() {
		t.Error("Next to LEFT must be TOP")
	}
}

func TestOpposite(t *testing.T) {
	if BOTTOM != TOP.Opposite() {
		t.Error("Opposite to TOP must be BOTTOM")
	}
	if LEFT != RIGHT.Opposite() {
		t.Error("Opposite to RIGHT must be LEFT")
	}
	if TOP != BOTTOM.Opposite() {
		t.Error("Opposite to BOTTOM must be TOP")
	}
	if RIGHT != LEFT.Opposite() {
		t.Error("Opposite to LEFT must be RIGHT")
	}
}

func TestAdjacentCoords(t *testing.T) {
	cases := []struct {
		x, y   int
		d      Direction
		rx, ry int
	}{
		{0, 0, TOP, 0, -1},
		{0, 0, RIGHT, 1, 0},
		{0, 0, BOTTOM, 0, 1},
		{0, 0, LEFT, -1, 0},
		{10, 10, TOP, 10, 9},
		{10, 10, RIGHT, 11, 10},
		{10, 10, BOTTOM, 10, 11},
		{10, 10, LEFT, 9, 10},
	}

	for i, c := range cases {
		ax, ay := c.d.AdjacentCoords(c.x, c.y)
		if ax != c.rx || ay != c.ry {
			t.Errorf("case %d: expected [%d, %d] but saw [%d, %d]", i, c.rx, c.ry, ax, ay)
		}
	}
}

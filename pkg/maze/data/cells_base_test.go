package data

import (
	"testing"
)

func TestBaseIsEmpty(t *testing.T) {
	s := newCellsBase()
	if !s.IsEmpty() {
		t.Error("It must be empty")
	}
}

func TestBaseHas(t *testing.T) {
	s := newCellsBase()
	a := NewCell(0, 0)
	b := NewCell(1, 1)

	if s.Has(a) || s.Has(b) {
		t.Error("did not add neither `a` nor `b` yet")
	}
}

func TestCellKey(t *testing.T) {
	a := cellKey(NewCell(0, 0))
	b := cellKey(NewCell(0, 1))
	c := cellKey(NewCell(1, 0))
	d := cellKey(NewCell(1, 1))
	a1 := cellKey(NewCell(0, 0))

	if a != "0;0" {
		t.Errorf("Unexpected %q", a)
	}
	if b != "0;1" {
		t.Errorf("Unexpected %q", b)
	}
	if c != "1;0" {
		t.Errorf("Unexpected %q", c)
	}
	if d != "1;1" {
		t.Errorf("Unexpected %q", d)
	}

	if a != a1 {
		t.Error("Key for similar cells must be same")
	}
}

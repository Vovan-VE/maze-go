package data

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	s := NewCellsSet()
	if !s.IsEmpty() {
		t.Error("It must be empty")
	}

	c := NewCell(0, 0)
	s.Add(c)
	if s.IsEmpty() {
		t.Error("It must not be empty")
	}

	s.Remove(c)
	if !s.IsEmpty() {
		t.Error("It must be empty")
	}
}

func TestAdd(t *testing.T) {
	s := NewCellsSet()
	c := NewCell(0, 0)
	s.Add(c)
	if s.IsEmpty() {
		t.Error("did not add")
	}
	s.Add(c)
	s.Add(NewCell(10, 20))

	defer func() {
		e := recover()
		if e == nil {
			t.Error("did not panic")
		}
		if e != "trying to add different cell with the same coords" {
			panic(e)
		}
	}()
	s.Add(NewCell(0, 0))
}

func TestRemove(t *testing.T) {
	s := NewCellsSet()
	a := NewCell(0, 0)
	b := NewCell(1, 1)
	s.Add(a)
	s.Add(b)

	s.Remove(a)
	if s.IsEmpty() {
		t.Error("must not be empty yet")
	}
	s.Remove(b)
	if !s.IsEmpty() {
		t.Error("must be empty")
	}
}

func TestHas(t *testing.T) {
	s := NewCellsSet()
	a := NewCell(0, 0)
	b := NewCell(1, 1)
	aCopy := NewCell(0, 0)

	if s.Has(a) || s.Has(b) {
		t.Error("did not add neither `a` nor `b` yet")
	}

	s.Add(a)
	if !s.Has(a) {
		t.Error("must to have `a`")
	}
	if !s.Has(aCopy) {
		t.Error("must to have `aCopy`")
	}
	if s.Has(b) {
		t.Error("did not add `b` yet")
	}

	s.Add(b)
	if !s.Has(b) {
		t.Error("must to have `b`")
	}
	if !s.Has(a) {
		t.Error("must to have `a` still")
	}

	s.Remove(a)
	if s.Has(a) {
		t.Error("must not have `a`")
	}
	if !s.Has(b) {
		t.Error("must to have `b` still")
	}

	s.Remove(b)
	if s.Has(a) || s.Has(b) {
		t.Error("must not have neither `a` nor `b`")
	}
}

func TestRandom(t *testing.T) {
	s := NewCellsSet()
	a := NewCell(0, 0)
	b := NewCell(1, 1)
	s.Add(a)
	s.Add(b)
	rest := map[*Cell]bool{
		a: true,
		b: true,
	}

	for i := 0; i < 100 && len(rest) > 0; i++ {
		c := s.Random()
		delete(rest, c)
	}
	if len(rest) > 0 {
		t.Error("Some cells wasn't returned")
	}
}

func TestRandomEmpty(t *testing.T) {
	s := NewCellsSet()
	defer func() {
		e := recover()
		if e == nil {
			t.Error("Did not panic")
		}
		if e != "The Set is empty" {
			panic(e)
		}
	}()
	s.Random()
}

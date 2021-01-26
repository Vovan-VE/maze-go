package data

import (
	"math/rand"
)

// Direction is a possible direction
type Direction uint8

const (
	// TOP is top/up direction
	TOP Direction = 1
	// RIGHT is right side direction
	RIGHT Direction = 2
	// BOTTOM is bottom/down direction
	BOTTOM Direction = 3
	// LEFT is left side direction
	LEFT Direction = 4
)

// RandomDirection returns a random direction
func RandomDirection() Direction {
	return Direction(1 + rand.Intn(4))
}

// Prev returns sibling anti-clockwise direction
func (d Direction) Prev() Direction {
	return Direction((d+2)%4 + 1)
}

// Next returns next clockwise direction
func (d Direction) Next() Direction {
	return Direction(d%4 + 1)
}

// Opposite returns an opposite direction
func (d Direction) Opposite() Direction {
	return Direction((d+1)%4 + 1)
}

// AdjacentCoords returns an coordinated ajdacent to given coordinates as given direction
func (d Direction) AdjacentCoords(x, y int) (nextX, nextY int) {
	switch d {
	case TOP:
		return x, y - 1

	case RIGHT:
		return x + 1, y

	case BOTTOM:
		return x, y + 1

	case LEFT:
		return x - 1, y
	}
	panic("Invalid direction")
}

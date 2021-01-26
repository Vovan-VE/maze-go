package data

// Cell reprecent one cell of a maze
type Cell struct {
	x, y       int
	TopWall    bool
	RightWall  bool
	BottomWall bool
	LeftWall   bool
}

// NewCell creates a new Cell with all walls set
func NewCell(x, y int) *Cell {
	return &Cell{x, y, true, true, true, true}
}

// NewCellOpened creates a new Cell with all walls removed
func NewCellOpened(x, y int) *Cell {
	return &Cell{x: x, y: y}
}

// X returns an X coordinate
func (c *Cell) X() int {
	return c.x
}

// Y returns an Y coordinate
func (c *Cell) Y() int {
	return c.y
}

// XY returns both X and Y coordinate
func (c *Cell) XY() (x, y int) {
	return c.x, c.y
}

// IsSameCoords checks coords equality with another Cell
func (c *Cell) IsSameCoords(to *Cell) bool {
	return c == to || c.x == to.x && c.y == to.y
}

// HasWallAt returns wall presense at direction
func (c *Cell) HasWallAt(d Direction) bool {
	switch d {
	case TOP:
		return c.TopWall

	case RIGHT:
		return c.RightWall

	case BOTTOM:
		return c.BottomWall

	case LEFT:
		return c.LeftWall
	}
	return false
}

// SetWallAt sets wall presence at direction
func (c *Cell) SetWallAt(d Direction, wall bool) {
	switch d {
	case TOP:
		c.TopWall = wall

	case RIGHT:
		c.RightWall = wall

	case BOTTOM:
		c.BottomWall = wall

	case LEFT:
		c.LeftWall = wall
	}
}

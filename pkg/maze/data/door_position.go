package data

// DoorPosition describes a door position an a maze
type DoorPosition struct {
	side   Direction
	offset int
}

// NewDoorPosition creates new structure
func NewDoorPosition(side Direction, offset int) *DoorPosition {
	return &DoorPosition{side, offset}
}

// Side returns a side
func (p *DoorPosition) Side() Direction {
	return p.side
}

// Offset returns an offset
func (p *DoorPosition) Offset() int {
	return p.offset
}

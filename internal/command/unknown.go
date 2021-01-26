package command

import (
	"fmt"
)

// Unknown is a "gen" cli command
type Unknown struct {
	*command
}

// NewUnknown creates an instance of Generate
func NewUnknown(pool Pooler, name string) *Unknown {
	return &Unknown{newCommand(pool, name)}
}

// Run implements Commander
func (c *Unknown) Run(args []string) (int, error) {
	return 1, fmt.Errorf("Unknown command: %s", c.Name())
}

package command

type command struct {
	pool Pooler
	name string
}

func newCommand(pool Pooler, name string) *command {
	return &command{pool, name}
}

// Name implements Commender Name
func (c *command) Name() string {
	return c.name
}

// Usage implements Commander
func (c *command) Usage() string {
	return ""
}

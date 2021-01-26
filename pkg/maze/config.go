package maze

// Config describes options for maze
type Config struct {
	Width, Height int
	BranchLength  int
	Format        string
	Options       map[string]string
}

// NewConfig creates Config
func NewConfig() *Config {
	return &Config{}
}

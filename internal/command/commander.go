package command

// Commander describes app cli command
type Commander interface {
	Run(args []string) (int, error)
	Name() string
	Usage() string
}

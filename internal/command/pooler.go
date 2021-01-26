package command

// Pooler describes collection of available Comander
type Pooler interface {
	Command(name string) Commander
	Commands() map[string]Commander
}

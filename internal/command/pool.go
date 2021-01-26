package command

import (
	"flag"

	"rsc.io/getopt"
)

const (
	cmdGen     = "gen"
	cmdHelp    = "help"
	defaultCmd = cmdGen
)

var commands = map[string]func(pool Pooler, name string) Commander{
	cmdGen: func(pool Pooler, name string) Commander {
		return NewGenerate(pool, name)
	},
	cmdHelp: func(pool Pooler, name string) Commander {
		return NewHelp(pool, name)
	},
}

type tPool struct {
}

// Run runs appropriate command with given arguments
func Run(args []string) (code int, err error) {
	command, args := (&tPool{}).resolveCommand(args[1:])
	code, err = command.Run(args)
	return
}

// Command implements Pooler
func (pool *tPool) Command(name string) Commander {
	create, ok := commands[name]
	if !ok {
		return nil
	}
	return create(pool, name)
}

// Commands implements Pooler
func (pool *tPool) Commands() (result map[string]Commander) {
	result = make(map[string]Commander, len(commands))
	for name, create := range commands {
		result[name] = create(pool, name)
	}
	return
}

func (pool *tPool) resolveCommand(args []string) (cmd Commander, arguments []string) {
	if len(args) > 0 {
		name := args[0]
		cmd = pool.Command(name)
		if cmd != nil {
			arguments = args[1:]
			return
		}
		if name == "" || name[0] != '-' {
			return NewUnknown(pool, name), nil
		}

		fs := getopt.NewFlagSet(cmdHelp, flag.ExitOnError)
		help := fs.Bool("help", false, "")
		fs.Alias("h", "help")
		fs.SetOutput(stderr)
		if err := fs.Parse(args); err != nil {
			return pool.Command(cmdHelp), nil
		}
		if *help {
			return pool.Command(cmdHelp), args
		}
	}
	return pool.Command(defaultCmd), args
}

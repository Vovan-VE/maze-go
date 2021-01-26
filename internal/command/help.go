package command

import (
	"flag"
	"fmt"
	"sort"

	"rsc.io/getopt"
)

// Help is a "help" cli command and "-h"
type Help struct {
	*command
}

// NewHelp creates an instance of Generate
func NewHelp(pool Pooler, name string) *Help {
	return &Help{newCommand(pool, name)}
}

// Run implements Commander
func (c *Help) Run(args []string) (int, error) {
	var help, list bool
	fs := getopt.NewFlagSet(c.Name(), flag.ExitOnError)
	fs.BoolVar(&help, "help", false, "show help")
	fs.BoolVar(&list, "list", false, "list all available commands")
	fs.Alias("h", "help")
	fs.Alias("l", "list")
	fs.SetOutput(stderr)

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if list {
		cmds := c.pool.Commands()
		names := make([]string, len(cmds))
		for name := range cmds {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			mark := ""
			if name == defaultCmd {
				mark = "        default command"
			}
			fmt.Fprintf(stdout, "  %s%s\n", name, mark)
		}
		fmt.Fprint(
			stdout,
			`
Run 'maze help [command]' to see help about specific command.")
`,
		)

		return 0, nil
	}

	rest := fs.Args()
	if len(rest) > 0 {
		name := rest[0]
		cmd := c.pool.Command(name)
		if cmd != nil {
			fmt.Fprint(stdout, cmd.Usage())
			return 0, nil
		}
		return 1, fmt.Errorf("unknown help topic %q", name)
	}

	fmt.Fprint(
		stdout,
		`maze [command] [options]

Run 'maze help -l' to see available commands.

Run 'maze help [command]' to see help about specific command.

The default command is 'gen'. Run 'maze help gen' too see help for generation.

TBW: ...

`,
	)
	return 0, nil
}

// Usage implements Commander
func (c *Help) Usage() string {
	return `maze help [options] [command]
maze (-h | --help) [options] [command]

Show help either about specified 'command' or common usage help.

Options:

    -l, --list
        List all available commands.

`
}

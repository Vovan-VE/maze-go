package command

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"rsc.io/getopt"

	"github.com/Vovan-VE/maze-go/internal/cli"
	"github.com/Vovan-VE/maze-go/pkg/maze"
	"github.com/Vovan-VE/maze-go/pkg/maze/format"
)

// Generate is a "gen" cli command
type Generate struct {
	*command
}

// NewGenerate creates an instance of Generate
func NewGenerate(pool Pooler, name string) *Generate {
	return &Generate{newCommand(pool, name)}
}

// Run implements Commaner interface
func (g *Generate) Run(args []string) (int, error) {
	config, err := initConfig(g.Name(), args)
	if err != nil {
		return 1, err
	}

	exporter := format.NewExporter(config.Format)
	exporter.ConfigureExport(config.Options)

	maze := maze.NewGenerator(config).Generate()

	result := exporter.ExportMaze(maze)
	fmt.Fprintln(stdout, result)

	return 0, nil
}

// Usage implements Commander interface
func (g *Generate) Usage() string {
	return `maze [gen] [options]

Generate a maze and export it to stdout.

Since 'gen' is the default command, the command name 'gen' is optional.

Options:

    -W <WIDTH>, --width=<WIDTH>
        Maze width in number of CELLs. Default is 30.

    -H <HEIGHT>, --height=<HEIGHT>
        Maze height in number of CELLs. Default is 10.

        Notice about uppercase -H to not mix with -h which is "help".

    -s <SIZE>, --size=<SIZE>
        Alternative way to set both width and height at once. The SIZE
        must be in form <WIDTH>x<HEIGHT>. So, the default size is 30x10.

    -B <BL>, --branch-length=<BL>
        The "branch length" option for generation. BL can be an integer > 1
        (a number of CELLs), string 'max' (which is WIDTH * HEIGHT'), or
        decimal from 0 to 1 as fraction of max (for example, 0.2 is
        round(0.2 * W * H)). Default is 10.

    -f <FORMAT>, --format=<FORMAT>
        Output format. Can be one of 'json' or 'text'. The default is 'text' to
        be human readable.

    -c <NAME>=<VALUE>
        Output format option. The '<NAME>' depends on chosen format in '-f'.
`
}

func initConfig(name string, args []string) (*maze.Config, error) {
	config := maze.NewConfig()
	var bl, size string
	options := cli.NewMapValue()
	fs := getopt.NewFlagSet(name, flag.ExitOnError)
	fs.IntVar(&config.Width, "width", 30, "maze width in number of CELLs")
	fs.IntVar(&config.Height, "height", 10, "maze height in number of CELLs")
	fs.StringVar(&bl, "branch-length", "10", "\"branch length\" option")
	fs.StringVar(&size, "size", "", "both width and height together")
	fs.StringVar(&config.Format, "format", "text", "outout format name")
	fs.Var(options, "c", "output options depending of given format")
	fs.Alias("W", "width")
	fs.Alias("H", "height")
	fs.Alias("B", "branch-length")
	fs.Alias("s", "size")
	fs.Alias("f", "format")
	fs.SetOutput(stderr)

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if size != "" {
		// TODO: check mixing -s with -W -H

		parts := strings.SplitN(size, "x", 3)
		if len(parts) != 2 {
			return nil, errors.New("size must be in form `<WIDTH>x<HEIGHT>`,  like `30x10`")
		}
		widthStr, heightStr := parts[0], parts[1]
		var err error
		if config.Width, err = parsePlusIntValue(widthStr); err != nil {
			return nil, fmt.Errorf("<WIDTH> %s in `-s` (`--size`)", err.Error())
		}
		if config.Height, err = parsePlusIntValue(heightStr); err != nil {
			return nil, fmt.Errorf("<HEIGHT> %s in `-s` (`--size`)", err.Error())
		}
	} else {
		if err := validatePlusIntValue(config.Width); err != nil {
			return nil, fmt.Errorf("<WIDTH> %s in `-W` (`--width`)", err.Error())
		}
		if err := validatePlusIntValue(config.Height); err != nil {
			return nil, fmt.Errorf("<HEIGHT> %s in `-H` (`--height`)", err.Error())
		}
	}

	if bl != "" {
		var err error
		if config.BranchLength, err = parseBranchLength(bl, config.Width*config.Height); err != nil {
			return nil, fmt.Errorf("<BL> %s in `-B` (`--branch-length`)", err.Error())
		}
	}

	if !format.HasExporter(config.Format) {
		return nil, errors.New("unknown format name in `-F` (`--format`)")
	}

	config.Options = options.Values()

	return config, nil
}

func parseBranchLength(input string, max int) (int, error) {
	if input == "max" {
		return max, nil
	}
	if n, err := parsePlusIntValue(input); err == nil {
		if n > max {
			n = max
		}
		return n, nil
	}
	if f, err := strconv.ParseFloat(input, 16); err == nil && f >= 0 && f <= 1 {
		return int(math.Round(f * float64(max))), nil
	}
	return 0, errors.New(
		"must be either integer greater then 1 (number of CELLs), " +
			"a string `max` to set <WIDTH>*<HEIGHT>, " +
			"or decimal from 0 to 1 as fraction of max",
	)
}

func parsePlusIntValue(input string) (int, error) {
	n, err := strconv.ParseUint(input, 10, 16)
	if err == nil {
		i := int(n)
		if err := validatePlusIntValue(i); err != nil {
			return 0, err
		}
		return i, nil
	}
	ne, ok := err.(*strconv.NumError)
	if ok {
		if ne.Err == strconv.ErrSyntax {
			return 0, errors.New("must be positive integer")
		}
		if ne.Err == strconv.ErrRange {
			return 0, errors.New("is out of range")
		}
	}
	return 0, err
}

func validatePlusIntValue(n int) error {
	if n <= 0 {
		return errors.New("must be greater then zero")
	}
	if n > 0xFFFF {
		return errors.New("is too big")
	}
	return nil
}

package cli

import (
	"flag"
)

type fsReader interface {
	Lookup(name string) *flag.Flag
	Visit(fn func(f *flag.Flag))
}

func IsFlagSet(fs fsReader, name string) bool {
	f := fs.Lookup(name)
	if f == nil {
		return false
	}
	name = f.Name
	found := false
	// https://github.com/rsc/getopt/issues/2
	fs.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

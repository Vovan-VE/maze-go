package command

import (
	"io"
	"os"
)

var stdout, stderr io.Writer = os.Stdout, os.Stderr

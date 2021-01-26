package main

import (
	"log"
	"os"

	"github.com/Vovan-VE/maze-go/internal/command"
)

func main() {
	code, err := command.Run(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
	os.Exit(code)
}

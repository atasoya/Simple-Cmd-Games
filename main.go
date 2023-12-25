package main

import (
	"os"
	C"atasoy/cmd-line-casino/commands"
)

func main() {

	args := os.Args[1:]

	switch args[0]{
	case "-h":
		C.Help()
	case "-H":
		C.Help()
	}

}

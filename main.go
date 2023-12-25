package main

import (
	"os"

	C "github.com/atasoya/simple-cmd-games/commands"
)

func main() {

	args := os.Args[1:]

	switch args[0] {
	case "-h":
		C.Help()
	case "-H":
		C.Help()
	case "-hr":
		default_horse_count := 5
		default_duration := 5
		C.Horse_race(default_horse_count, default_duration)
	}

}

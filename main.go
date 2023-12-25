package main

import (
	"fmt"
	"os"

	Help_command "github.com/atasoya/simple-cmd-games/commands/Help"
	Horse_command "github.com/atasoya/simple-cmd-games/commands/Horse_race"
)

func main() {

	args := os.Args[1:]

	switch args[0] {
	case "-h":
		Help_command.Help()
	case "-H":
		Help_command.Help()
	case "-hr":
		duration, horse_count := Horse_command.Validate(args)

		Horse_command.Horse_race(horse_count, duration)
	default:
		fmt.Println(args)
	}

}

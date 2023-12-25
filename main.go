package main

import (
	"fmt"
	"os"

	Help_command "github.com/atasoya/simple-cmd-games/commands/Help"
	Horse_command "github.com/atasoya/simple-cmd-games/commands/Horse_race"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		asciiArt := `
	    _____ _____ __  __ _____  _      ______                    
	   / ____|_   _|  \/  |  __ \| |    |  ____|                   
	  | (___   | | | \  / | |__) | |    | |__                      
	   \___ \  | | | |\/| |  ___/| |    |  __|                     
	   ____) |_| |_| |  | | |    | |____| |____                    
	  |_____/|_____|_|__|_|_| ___|______|______| __ ______  _____  
	   / ____| |    |_   _|  / ____|   /\   |  \/  |  ____|/ ____| 
	  | |    | |      | |   | |  __   /  \  | \  / | |__  | (___   
	  | |    | |      | |   | | |_ | / /\ \ | |\/| |  __|  \___ \  
	  | |____| |____ _| |_  | |__| |/ ____ \| |  | | |____ ____) | 
	   \_____|______|_____|  \_____/_/    \_\_|  |_|______|_____/  
	  |_   _| \ | |  / ____|/ __ \                                 
	    | | |  \| | | |  __| |  | |                                
	    | | | .   | | | |_ | |  | |                                
	   _| |_| |\  | | |__| | |__| |                                
	  |_____|_| \_|  \_____|\____/                                 
																   
	  -h to see all commands															   
	 `
		fmt.Println(asciiArt)
	} else {
		switch args[0] {
		case "-h":
			Help_command.Help()
		case "-H":
			Help_command.Help()
		case "-hr":
			duration, horse_count := Horse_command.Validate(args)

			Horse_command.Horse_race(horse_count, duration)
		}
	}
}

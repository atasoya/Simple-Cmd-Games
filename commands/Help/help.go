package commands

import "fmt"

func Help() {
	fmt.Println("Usage: go run main.go -[code of the game] -[additional arguments if any]")
	fmt.Println("-hr: Horse race , [-d duration] [-hc horse count]")
}

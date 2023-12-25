package commands

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Horse struct {
	Speed    int
	Position []string
}

func (h *Horse) move() {
	for i := 1; i < h.Speed; i++ {
		h.Position = append(h.Position, "-")
	}
}

func Horse_race(horse_count int, duration int) {

	rand.Seed(time.Now().UnixNano())

	horse_array := make([]Horse, horse_count)

	for i := 0; i < horse_count; i++ {
		new_horse := Horse{
			Speed:    rand.Intn(10-1) + 1,
			Position: []string{"-"},
		}
		horse_array[i] = new_horse
	}

	tick := func() {
		for i := 0; i < horse_count; i++ {
			current_horse := &horse_array[i]
			current_horse.move()
		}
	}

	render := func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		for i := 0; i < horse_count; i++ {
			current_horse := horse_array[i]
			fmt.Printf("[%v] ", i+1)
			fmt.Printf("%s", strings.Join(current_horse.Position, ""))
			fmt.Printf("\n")
		}
	}

	game := func() {
		change_speed := func() {
			for i := 0; i < horse_count; i++ {
				current_horse := &horse_array[i]
				current_horse.Speed = rand.Intn(10-1) + 1
			}
		}
		render()

		var input string

		fmt.Println("Press [P] to play")

		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if input == "P" || input == "p" {

			for i := 0; i < duration+1; i++ {
				if i == duration/2 {
					change_speed()
				} else if i == duration/4 {
					change_speed()
				} else if i == 3*(duration/4) {
					change_speed()
				}
				render()
				tick()
				time.Sleep(1 * time.Second)
			}
		}

	}

	game()

}

func Validate(args []string) (int, int) {
	duration := 5    // Default duration
	horse_count := 5 // Default horse count

	for i := 0; i < len(args); i++ {
		if args[i] == "-d" {
			int_value, err := strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Println("Invalid duration argument")
			} else {
				duration = int_value
			}
		}
		if args[i] == "-hc" {
			int_value, err := strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Println("Invalid horse count argument")
			} else {
				horse_count = int_value
			}
		}
	}

	return duration, horse_count
}

package commands

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
			Speed:    rand.Intn(horse_count-1) + 1,
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
					for i := 0; i < horse_count; i++ {
						current_horse := &horse_array[i]
						current_horse.Speed = rand.Intn(horse_count-1) + 1
					}
				}
				render()
				tick()
				time.Sleep(1 * time.Second)
			}
		}

	}

	game()

}

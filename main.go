package main

import (
	"fmt"
	"os"
	"strconv"
)

const statsFileName = "stats.txt"

func main() {
	var p Player
	p.LoadFromFile(statsFileName)

	defer func() {
		p.SaveToFile(statsFileName)
	}()

	for {
		LoadMenu()
		ch, _ := GetPlayerInput()

		switch ch {
		case 'e', 'E':
			p.SaveToFile(statsFileName)
			os.Exit(0)
		case 's', 'S':
			p.ShowStats()
		case 'r', 'R':
			p.ResetStats()
		default:
			i, err := strconv.Atoi(string(ch))

			// player press a key which is not in the menu.
			if err != nil || i < 1 || i > 5 {
				fmt.Println()
				fmt.Println("Please choose one of the options in the main menu!")
			} else {
				PlayRound(&p, i)
			}

		}
		Pause()
	}
}

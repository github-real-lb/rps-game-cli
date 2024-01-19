package main

import (
	"fmt"
	"os"
	"strconv"
)

const statsFileName = "stats.txt"

func main() {
	var p Player
	if err := p.LoadFromFile(statsFileName); err != nil {
		fmt.Println("Error:", err)
		Pause()
	}

	defer func() {
		if err := p.SaveToFile(statsFileName); err != nil {
			fmt.Println("Error:", err)
			Pause()
		}
	}()

	for {
		LoadMenu()
		ch, err := GetPlayerInput()

		if err != nil {
			panic(err)
		}

		switch ch {
		case 'e', 'E':
			if err := p.SaveToFile(statsFileName); err != nil {
				fmt.Println("Error:", err)
				Pause()
			}
			os.Exit(0)
		case 's', 'S':
			p.ShowStats()
		case 'r', 'R':
			if err := p.ResetStats(); err != nil {
				panic(err)
			}
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

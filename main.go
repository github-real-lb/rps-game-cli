package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

// RPS is a map of game's draw options and their int values.
var RPS = map[int]string{
	1: "Rock",
	2: "Paper",
	3: "Scissors",
	4: "Lizard",
	5: "Spock",
}

func main() {
	for {
		welcomeScreen()
		playerDraw := getPlayerDraw()
		computerDraw := getComputerDraw()

		fmt.Println()
		fmt.Println()
		fmt.Printf("You chose %v.\n", RPS[playerDraw])
		fmt.Printf("Computer chose %v.\n", RPS[computerDraw])
		fmt.Println()

		if playerDraw == computerDraw {
			fmt.Println("It's a draw.")
		} else if didPlayerWon(playerDraw, computerDraw) {
			fmt.Println("You've won!")
		} else {
			fmt.Println("You've lost!")
		}
		time.Sleep(time.Second * 6)
	}
}

// clearScreen clears the screen.
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// welcomeScreen prints welcome message and the game's draw options menu.
func welcomeScreen() {
	clearScreen()

	fmt.Println("Welcome to the game Rock, Paper, Scissor, Lizard & Spock")
	fmt.Println("Please choose one of the following:")
	for i := 1; i <= len(RPS); i++ {
		fmt.Println(i, RPS[i])
	}
	fmt.Println("\n0 Exit")
}

// getPlayerDraw gets player draw and returns it's value.
func getPlayerDraw() int {
	for {
		fmt.Print("-> ")

		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		fmt.Print(string(ch))

		i, err := strconv.Atoi(string(ch))

		if i == 0 {
			os.Exit(0)
		}

		if err == nil && (i >= 1 && i <= 5) {
			return i
		}

		fmt.Println()
		fmt.Println("Please press a number between 1 and 5 to indicate your preferred option, or press 0 to exit!")
		fmt.Println()
	}
}

// getComputerDraw gets computer random draw and returns it's value.
func getComputerDraw() int {
	return rand.Intn(len(RPS)-1) + 1
}

// didPlayerWon receives player's and computer's draw and returns true if player has won.
func didPlayerWon(playerDraw, computerDraw int) bool {
	if playerDraw == computerDraw {
		return false
	} else {
		switch playerDraw {
		case 1:
			return computerDraw != 2 && computerDraw != 5
		case 2:
			return computerDraw != 3 && computerDraw != 4
		case 3:
			return computerDraw != 5 && computerDraw != 1
		case 4:
			return computerDraw != 3 && computerDraw != 1
		case 5:
			return computerDraw != 4 && computerDraw != 2
		default:
			return false
		}
	}
}

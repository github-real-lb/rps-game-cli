package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

// stats is a struct to record player's stat during the game.
type stats struct {
	won  int
	lost int
	draw int
}

// RPS is a map of game's draw options and their int values.
var RPS = map[int]string{
	1: "Rock",
	2: "Paper",
	3: "Scissors",
	4: "Lizard",
	5: "Spock",
}

const statsFileName = "stats.txt"

var s = stats{0, 0, 0}

func main() {
	s.loadFromFile(statsFileName)

	defer func() {
		s.saveToFile(statsFileName)
	}()

	for {
		menuLoad()
		ch := getPlayerChoice()

		switch ch {
		case 'e', 'E':
			s.saveToFile(statsFileName)
			os.Exit(0)
		case 's', 'S':
			showStats()
		case 'r', 'R':
			resetStats()
		default:
			i, err := strconv.Atoi(string(ch))

			// player press a key which is not in the menu.
			if err != nil || i < 1 || i > 5 {
				fmt.Println()
				fmt.Println("Please choose one of the options in the main menu!")
			} else {
				playGame(i)
			}

		}
		pause()
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

func menuLoad() {
	clearScreen()

	fmt.Println("Welcome to the game Rock, Paper, Scissor, Lizard & Spock")
	fmt.Println("Please choose one of the following:")
	for i := 1; i <= len(RPS); i++ {
		fmt.Println(i, RPS[i])
	}
	fmt.Println()
	fmt.Println("s Show stats")
	fmt.Println("r Reset stats")
	fmt.Println("e Exit")
	fmt.Println()
	fmt.Print("-> ")
}

// getPlayerChoice gets a single key press and returns it.
func getPlayerChoice() rune {
	ch, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}

	fmt.Print(string(ch))

	return ch
}

// playGame receives player's draw (between 1 to 5), then gets computer's draw and prompts who won.
func playGame(playerDraw int) {
	computerDraw := getComputerDraw()

	fmt.Println()
	fmt.Println()
	fmt.Printf("You chose %v.\n", RPS[playerDraw])
	fmt.Printf("Computer chose %v.\n", RPS[computerDraw])
	fmt.Println()

	if playerDraw == computerDraw {
		s.draw++
		fmt.Println("It's a draw.")
	} else if didPlayerWon(playerDraw, computerDraw) {
		s.won++
		fmt.Println("You've won!")
	} else {
		s.lost++
		fmt.Println("You've lost!")
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

// showStats display the player's statistics of wins, loses and draws.
func showStats() {
	t := s.won + s.lost + s.draw

	fmt.Println()
	fmt.Println()
	fmt.Println("Your Stats are:")
	fmt.Printf("You've won %d times (%.1f%%).\n", s.won, float64(s.won)/float64(t)*100)
	fmt.Printf("You've lost %d times (%.1f%%).\n", s.lost, float64(s.lost)/float64(t)*100)
	fmt.Printf("You've draw %d times (%.1f%%).\n", s.draw, float64(s.draw)/float64(t)*100)
}

// resetStats reset player's stats to 0
func resetStats() {
	fmt.Println()
	fmt.Println()
	fmt.Println("You are about to reset your stats!")
	fmt.Println("Are you sure (y/n)? ")

	ch := getPlayerChoice()

	fmt.Println()
	if ch == 'y' || ch == 'Y' {
		fmt.Println("Reseting stats...")
		s = stats{0, 0, 0}
	} else {
		fmt.Println("Action canceled. Stats are kept.")
	}
}

// saveStats save the players stats to a file.
func (s *stats) saveToFile(fileName string) {
	str := fmt.Sprintf("%d,%d,%d", s.won, s.lost, s.draw)
	err := os.WriteFile(fileName, []byte(str), 0666)

	if err != nil {
		fmt.Println()
		fmt.Println()
		fmt.Println("Cannot save game! Error:", err)
	}
}

// loadStats load the players stats from a file.
func (s *stats) loadFromFile(fileName string) {
	if bs, err := os.ReadFile(fileName); err == nil {
		if ss := strings.Split(string(bs), ","); len(ss) == 3 {
			s.won, _ = strconv.Atoi(ss[0])
			s.lost, _ = strconv.Atoi(ss[1])
			s.draw, _ = strconv.Atoi(ss[2])
		}
	}
}

// pause waits for a key press before continue with game.
func pause() {
	fmt.Println()
	fmt.Println("Press any key to continue...")
	keyboard.GetSingleKey()
}

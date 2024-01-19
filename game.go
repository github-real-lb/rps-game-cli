package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

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

// stats is a struct to record player's stat during the game.
type stats struct {
	won  int
	lost int
	draw int
}

type Player struct {
	Stats stats
}

// clearScreen clears the screen.
func ClearScreen() {
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

// LoadMenu prints the main menu of the game.
func LoadMenu() {
	ClearScreen()

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

// GetPlayerInput gets a single key press and returns it.
// In case of error, it returns ' ' and the error.
func GetPlayerInput() (rune, error) {
	ch, _, err := keyboard.GetSingleKey()
	if err != nil {
		return ' ', err
	} else {
		fmt.Print(string(ch))
		return ch, err
	}
}

// PlayRound receives player's draw (between 1 to 5), then gets computer's draw and prompts who won.
func PlayRound(p *Player, playerDraw int) {
	computerDraw := getComputerDraw()

	fmt.Println()
	fmt.Println()
	fmt.Printf("You chose %v.\n", RPS[playerDraw])
	fmt.Printf("Computer chose %v.\n", RPS[computerDraw])
	fmt.Println()

	if playerDraw == computerDraw {
		p.Stats.draw++
		fmt.Println("It's a draw.")
	} else if didPlayerWon(playerDraw, computerDraw) {
		p.Stats.won++
		fmt.Println("You've won!")
	} else {
		p.Stats.lost++
		fmt.Println("You've lost!")
	}
}

// getComputerDraw gets computer random draw and returns it's value.
func getComputerDraw() int {

	return rand.Intn(len(RPS)*20)%len(RPS) + 1
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

// ShowStats display the player's statistics of wins, loses and draws.
func (p *Player) ShowStats() {
	t := p.Stats.won + p.Stats.lost + p.Stats.draw

	fmt.Println()
	fmt.Println()
	fmt.Println("Your Stats are:")
	fmt.Printf("You've won %d times (%.1f%%).\n", p.Stats.won, float64(p.Stats.won)/float64(t)*100)
	fmt.Printf("You've lost %d times (%.1f%%).\n", p.Stats.lost, float64(p.Stats.lost)/float64(t)*100)
	fmt.Printf("You've draw %d times (%.1f%%).\n", p.Stats.draw, float64(p.Stats.draw)/float64(t)*100)
}

// ResetStats prompts 'Are you sure?' and reset player's stats to 0 if pressed 'y'.
// It returns true is stats were reset and false if not.
// In case of error on recieving input, it returns false and the error.
func (p *Player) ResetStats() error {
	fmt.Println()
	fmt.Println()
	fmt.Println("You are about to reset your stats!")
	fmt.Print("Are you sure (y/n)? ")

	ch, err := GetPlayerInput()

	if err != nil {
		return err
	}

	fmt.Println()
	if ch == 'y' || ch == 'Y' {
		fmt.Println("Reseting stats...")
		p.Stats = stats{0, 0, 0}

	} else {
		fmt.Println()
		fmt.Println("Action canceled. Stats are kept.")
	}

	return nil
}

// SaveToFile save the players stats to a file.
func (p *Player) SaveToFile(fileName string) error {
	str := fmt.Sprintf("%d,%d,%d", p.Stats.won, p.Stats.lost, p.Stats.draw)
	err := os.WriteFile(fileName, []byte(str), 0666)
	return err
}

// LoadFromFile loads the players stats from a file.
func (p *Player) LoadFromFile(fileName string) error {
	p.Stats = stats{0, 0, 0}
	bs, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	ss := strings.Split(string(bs), ",")

	if len(ss) != 3 {
		return errors.New("file isn't in the correct format")
	}

	if p.Stats.won, err = strconv.Atoi(ss[0]); err != nil {
		return errors.New("file isn't in the correct format")
	}

	if p.Stats.lost, err = strconv.Atoi(ss[1]); err != nil {
		return errors.New("file isn't in the correct format")
	}

	if p.Stats.draw, err = strconv.Atoi(ss[2]); err != nil {
		return errors.New("file isn't in the correct format")
	}

	return nil
}

// Pause waits for a key press before continue with game.
func Pause() {
	fmt.Println()
	fmt.Println("Press any key to continue...")
	keyboard.GetSingleKey()
}

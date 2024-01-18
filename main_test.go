package main

import (
	"os"
	"testing"
)

func TestGetComputerDraw(t *testing.T) {
	i := getComputerDraw()
	if i < 1 || i > 5 {
		t.Errorf("Expected computer to draw a number between 1 to 5, but got %d.", i)
	}
}

func TestDidPlayerWon(t *testing.T) {
	if didPlayerWon(1, 2) {
		t.Errorf("Expected Player to lose when player draw 1 (Rock) and computer draw 2 (Paper).")
	}

	if didPlayerWon(1, 5) {
		t.Errorf("Expected Player to lose when player draw 1 (Rock) and computer draw 5 (Spock).")
	}

	if didPlayerWon(2, 3) {
		t.Errorf("Expected Player to lose when player draw 2 (Paper) and computer draw 3 (Scissors).")
	}

	if didPlayerWon(2, 4) {
		t.Errorf("Expected Player to lose when player draw 2 (Paper) and computer draw 4 (Lizard).")
	}

	if didPlayerWon(3, 1) {
		t.Errorf("Expected Player to lose when player draw 3 (Scissors) and computer draw 1 (Rock).")
	}

	if didPlayerWon(3, 5) {
		t.Errorf("Expected Player to lose when player draw 3 (Scissors) and computer draw 5 (Spock).")
	}

	if didPlayerWon(4, 1) {
		t.Errorf("Expected Player to lose when player draw 4 (Lizard) and computer draw 1 (Rock).")
	}

	if didPlayerWon(4, 3) {
		t.Errorf("Expected Player to lose when player draw 4 (Lizard) and computer draw 3 (Scissors).")
	}

	if didPlayerWon(5, 2) {
		t.Errorf("Expected Player to lose when player draw 5 (Spock) and computer draw 2 (Paper).")
	}

	if didPlayerWon(5, 4) {
		t.Error("Expected Player to lose when player draw 5 (Spock) and computer draw 4 (Lizard).")
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	os.Remove("_statstesting.txt")
	var ss = stats{123, 456, 789}
	var sl = stats{0, 0, 0}
	ss.saveToFile("_statstesting.txt")
	sl.loadFromFile("_statstesting.txt")

	if ss.won != sl.won || ss.lost != sl.lost || ss.draw != sl.draw {
		t.Error("Expected to load the same stats that were saved to file.")
	}
	os.Remove("_statstesting.txt")
}

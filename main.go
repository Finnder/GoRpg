package main

import (
	"fmt"
	pm "gorpg/PlayerManager"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	Menu()
}

func ClearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Menu() {

	userInput := ""

	// Entry Menu Text
	color.Yellow("Welcome to Golang RPG!!!")
	fmt.Println("1. Begin")
	fmt.Println("2. Exit")
	fmt.Print("> ")
	fmt.Scan(&userInput)

	if userInput == "1" {
		ClearTerm()

		fmt.Println("1. Select Saved File")
		fmt.Println("2. Create New Character")
		fmt.Print(">")
		fmt.Scan(&userInput)

		if userInput == "1" {
			// Loads Player Data
			pm.LoadPlayer()
		} else {
			// Creates New Player and Saves Data
			pm.NewPlayer()
		}

	} else {
		os.Exit(3)
	}

}

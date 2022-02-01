package PlayerManager

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Assign Global Player Variables
var PlayerName string = ""
var PlayerLevel int = 0
var PlayerXp int = 0
var PlayerMaxXp int = 0
var PlayerAttackDmg int = 0

// Player Object
type Player struct {
	Name      string
	Level     int
	XP        int
	MaxXP     int
	AttackDmg int
}

// File path
var filePathPlayerData string = "./Saved/PlayerData/"
var PlayerChoosenFilePath string = filePathPlayerData

// List out players data
func ListPlayerData() {
	title := color.New(color.FgGreen, color.Bold)

	title.Println("Player Stats")
	fmt.Println("------------")
	fmt.Println("Name: ", PlayerName)
	fmt.Println("Level: ", PlayerLevel)
	fmt.Println("XP:", PlayerXp, "/", PlayerMaxXp)
}

// Save the players info
func SavePlayerData(player Player) {

	// declare local data and covert player struct ints into Strings
	var playerlevel string = strconv.Itoa(player.Level)
	var playerXp string = strconv.Itoa(player.XP)
	var playerMaxXp string = strconv.Itoa(player.MaxXP)
	var playerAttackDmg string = strconv.Itoa(player.AttackDmg)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filePathPlayerData+player.Name+".dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// WRITE TO FILE - SAVE PLAYER DATA
	f.WriteString(player.Name + "\n")
	f.WriteString(playerlevel + "\n")
	f.WriteString(playerXp + "\n")
	f.WriteString(playerMaxXp + "\n")
	f.WriteString(playerAttackDmg + "\n")
	f.Close()
}

// When we want to create a new player
func NewPlayer() {

	// Default variables
	var playerName string = ""
	var playerLevel int = 0
	var playerXP int = 0
	var playerMaxXP int = 5
	var attackDmg int = 1

	// Get user input
	fmt.Println("Please Enter New Your Characters Name: ")
	fmt.Scan(&playerName)

	// Create player object
	var player Player = Player{playerName, playerLevel, playerXP, playerMaxXP, attackDmg}

	// Saves Player Info
	SavePlayerData(player)
}

// Load Player variables back into game
func LoadPlayer() {

	var userInput string = ""

	fmt.Println(" ") // Spacer

	files, _ := ioutil.ReadDir(filePathPlayerData)

	color.Red("    CHARACTER FILES    ")
	color.Red("-----------------------")

	for _, file := range files {
		color.Red(file.Name())
	}

	color.Red("-----------------------")
	fmt.Println(" ")

	fmt.Print("Please Enter A Character To Load: ")
	fmt.Scan(&userInput)

	if userInput != "" {
		for _, file := range files {
			if userInput == file.Name() {

				// Assign file data to global variables
				selectedFile, err := os.Open(filePathPlayerData + file.Name())
				check(err)

				scanner := bufio.NewScanner(selectedFile)

				var FileData []string

				for scanner.Scan() {
					FileData = append(FileData, scanner.Text())
				}

				PlayerName = FileData[0]

				PlayerLevel, err = strconv.Atoi(FileData[1])
				check(err)

				PlayerXp, err = strconv.Atoi(FileData[2])
				check(err)

				PlayerMaxXp, err = strconv.Atoi(FileData[3])
				check(err)

				PlayerAttackDmg, err = strconv.Atoi(FileData[4])
				check(err)
			}
		}
	} else {
		fmt.Print("Please Enter A Valid File Name!")
		LoadPlayer() // Reload function and ask for character to load
	}

}

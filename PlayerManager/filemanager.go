package PlayerManager

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Player struct {
	Name      string
	Level     int
	XP        int
	MaxXP     int
	AttackDmg int
}

var filePathPlayerData string = "./Saved/PlayerData/"

func SavePlayerData(Player_selected Player) {

	// declare local data and covert player struct ints into Strings
	var playerlevel string = strconv.Itoa(Player_selected.Level)
	var playerXp string = strconv.Itoa(Player_selected.XP)
	var playerMaxXp string = strconv.Itoa(Player_selected.MaxXP)
	var playerAttackDmg string = strconv.Itoa(Player_selected.AttackDmg)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filePathPlayerData+Player_selected.Name+".dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// WRITE TO FILE - SAVE PLAYER DATA
	f.WriteString(Player_selected.Name + "\n")
	f.WriteString(playerlevel + "\n")
	f.WriteString(playerXp + "\n")
	f.WriteString(playerMaxXp + "\n")
	f.WriteString(playerAttackDmg + "\n")

	f.Close()
}

func NewPlayer() {

	var playerName string = ""
	var playerLevel int = 0
	var playerXP int = 0
	var playerMaxXP int = 5
	var attackDmg int = 1

	fmt.Println("Please Enter New Your Characters Name: ")
	fmt.Scan(&playerName)

	var player Player = Player{playerName, playerLevel, playerXP, playerMaxXP, attackDmg}

	SavePlayerData(player) // Saves Player Info
}

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

}

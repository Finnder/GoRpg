package EnemyManager

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/fatih/color"
)

type Enemy struct {
	Name      string
	Level     int
	AttackDmg int
	Health    int
}

type TwitchResponse struct {
	TwitchUsername string `json:"viewers"`
	Broadcaster    string `json:"broadcaster"`
	ChatterCount   string `json:"chatter_count"`
}

var enemyNames = []string{"Joe"}

func RandomName() string {
	var rng_enemy_names int = rand.Intn(len(enemyNames))
	return enemyNames[rng_enemy_names]
}

// Begins Enemy Ecounter
func EnemyEncounterBegin(level int) {
	var enemy Enemy = Enemy{RandomName(), level, level * 2, level * 10}

	fmt.Println(enemy)
	color.Red("A level " + strconv.Itoa(level) + " enemy apeared!")

}

func PlayerTurnBegins() {

	fmt.Println("Your Turn Has Began!")

}

func EnemyTurnBegins() {

	fmt.Println("It is now the enemys turn")
}

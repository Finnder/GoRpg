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
	XPGain    int
}

var enemyNames = []string{"Joe"}
var enemyXPMultiplier = 1

func RandomName() string {
	var rng_enemy_names int = rand.Intn(len(enemyNames))
	return enemyNames[rng_enemy_names]
}

// Begins Enemy Ecounter
func EnemyEncounterBegin(level int) {
	var enemy Enemy = Enemy{RandomName(), level, level * 2, level * 10, level * enemyXPMultiplier}

	fmt.Println(enemy)
	color.Red("A level " + strconv.Itoa(level) + " enemy apeared!")

}

func PlayerTurnBegins() {

	fmt.Println("Your Turn Has Began!")

}

func EnemyTurnBegins() {

	fmt.Println("It is now the enemys turn")
}

func EnemyEncounterEnds() {
	fmt.Println("You defeated the enemy!")

}

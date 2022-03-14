package EnemyManager

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

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

	// Create enemy
	var (
		enemyName      string = RandomName()
		enemyLevel     int    = level
		enemyAttackDmg int    = level * 2.0
		enemyHealth    int    = level * 10.0
		enemyXPGain    int    = level * enemyXPMultiplier
	)

	// Instatiate Enemy
	var enemy Enemy = Enemy{enemyName, enemyLevel, enemyAttackDmg, enemyHealth, enemyXPGain}

	fmt.Println(enemy)
	color.Red("A level " + strconv.Itoa(level) + " enemy apeared!")
	PlayerTurnBegins()
}

func PlayerTurnBegins() {

	fmt.Println("Your Turn Has Began!")

}

func SelectMoves() {

}

func EnemyTurnBegins() {
	fmt.Println("It is now the enemys turn")
	time.Sleep(2 * time.Second)
	fmt.Println("Enemy is now deciding what to do...")
	time.Sleep(4 * time.Second)
}

func EnemyEncounterEnds() {
	fmt.Println("You defeated the enemy!")
}

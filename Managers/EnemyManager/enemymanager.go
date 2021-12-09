package EnemyManager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Enemy struct {
	Name      string
	Level     int
	AttackDmg float64
	Health    float64
}

type TwitchResponse struct {
	TwitchUsername string `json:"viewers"`
	Broadcaster    string `json:"broadcaster"`
	ChatterCount   string `json:"chatter_count"`
}

var EnemyNames = []string{""}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// When we want encounter of enemy to being do this
func EnemyEncounterBegin() {

}

func GetTwitchViewers() {

	foo1 := TwitchResponse{}
	getJson("https://tmi.twitch.tv/group/user/finnder7/chatters", &foo1)
	fmt.Println(foo1.ChatterCount)

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Player struct which contains a list of attributes
type Player struct {
	Name      string  `json:"name"`
	Attempts      int     `json:"attempts"`
	AvgHitAngle   float64 `json:"avg_hit_angle"`
	MaxHitSpeed   float64     `json:"max_hit_speed"`
	AvgHitSpeed   float64 `json:"avg_hit_speed"`
	MaxDistance   int     `json:"max_distance"`
	AvgHrDistance int     `json:"avg_hr_distance"`
	BrlPa         float64 `json:"brl_pa"`
}

var Players []Player

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my RESTful API Homepage")
}

func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["name"]
	for _, player := range Players {
		if player.Name == key {
			json.NewEncoder(w).Encode(player)
		}
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var player Player
	json.Unmarshal(reqBody, &player)
	Players = append(Players, player)
	json.NewEncoder(w).Encode(player)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Name := vars["name"]
	for index, player := range Players {
		if player.Name == Name {
			Players = append(Players[:index], Players[index+1:]...)
		}
	}
}

func apiRequests() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", homePage)
	route.HandleFunc("/players", getAllPlayers)
	route.HandleFunc("/players", createPlayer).Methods("POST")
	route.HandleFunc("/players/{name}", deletePlayer).Methods("DELETE")
	route.HandleFunc("/players/{name}", getPlayer)
	log.Fatal(http.ListenAndServe(":3000", route))
}

func main() {
	Players = []Player{
		Player{Name: "Shohei Ohtani",
			Attempts:      133,
			AvgHitAngle:   17.2,
			MaxHitSpeed:   119,
			AvgHitSpeed:   92.7,
			MaxDistance:   451,
			AvgHrDistance: 414,
			BrlPa:         13.5},
		Player{Name: "Ronald Acuña Jr.",
			Attempts:      148,
			AvgHitAngle:   19.2,
			MaxHitSpeed:   115.8,
			AvgHitSpeed:   94.3,
			MaxDistance:   481,
			AvgHrDistance: 411,
			BrlPa:         12},
		Player{Name: "Rafael Devers",
			Attempts:      148,
			AvgHitAngle:   9.3,
			MaxHitSpeed:   112.4,
			AvgHitSpeed:   92.6,
			MaxDistance:   452,
			AvgHrDistance: 413,
			BrlPa:         11.9},
		Player{Name: "Adolis García",
			Attempts:      138,
			AvgHitAngle:   12.3,
			MaxHitSpeed:   112.4,
			AvgHitSpeed:   92.1,
			MaxDistance:   431,
			AvgHrDistance: 396,
			BrlPa:         11.6},
		Player{Name: "Aaron Judge",
			Attempts:      143,
			AvgHitAngle:   11.5,
			MaxHitSpeed:   117.4,
			AvgHitSpeed:   97,
			MaxDistance:   443,
			AvgHrDistance: 398,
			BrlPa:         11.4},
		Player{Name: "Josh Donaldson",
			Attempts:      133,
			AvgHitAngle:   12.6,
			MaxHitSpeed:   113.1,
			AvgHitSpeed:   92.7,
			MaxDistance:   425,
			AvgHrDistance: 409,
			BrlPa:         11.1},
		Player{Name: "Nelson Cruz",
			Attempts:      138,
			AvgHitAngle:   10.1,
			MaxHitSpeed:   116.6,
			AvgHitSpeed:   92.3,
			MaxDistance:   439,
			AvgHrDistance: 406,
			BrlPa:         10.4},
		Player{Name: "Vladimir Guerrero Jr.",
			Attempts:      169,
			AvgHitAngle:   8,
			MaxHitSpeed:   117.4,
			AvgHitSpeed:   94.5,
			MaxDistance:   465,
			AvgHrDistance: 407,
			BrlPa:         10.3},
		Player{Name: "Salvador Perez",
			Attempts:      163,
			AvgHitAngle:   14.5,
			MaxHitSpeed:   114.2,
			AvgHitSpeed:   92.8,
			MaxDistance:   460,
			AvgHrDistance: 407,
			BrlPa:         9.9},
		Player{Name: "Matt Olson",
			Attempts:      163,
			AvgHitAngle:   15.9,
			MaxHitSpeed:   115.3,
			AvgHitSpeed:   92.8,
			MaxDistance:   445,
			AvgHrDistance: 405,
			BrlPa:         9.6},
	}
	apiRequests()
}

package gameplay

import "../gamemap"
import "fmt"
import "os/exec"
import "encoding/json"
import "log"

type Move struct {
	From   int
	To     int
	Troops int
	Type   string
}

type Player struct {
	Name          string
	Team_index    int
	Player_index  int
	// AI_specs
	Code_language string
	Code_path     string
}

func (p Player) PlayTurn(Map gamemap.MapGraph, players []Player, teams [][]Player, time_limit float64) {
	moves := p.GenerateMoves(Map, players, teams, time_limit)
	for _, move := range moves {
		if Map.OwnsBase(move.From, p.Player_index) && move.Type == "attack" {
			outcome := Map.Attack(move.From, move.To, move.Troops)
			fmt.Printf(outcome)
		} else if Map.OwnsBase(move.From, p.Player_index) && move.Type == "support" {
			outcome := Map.Support(move.From, move.To, move.Troops)
			fmt.Printf(outcome)
		} else if Map.OwnsBase(move.From, p.Player_index) && move.Type == "send" {
			if Map.OwnsBase(move.To, p.Player_index) {
				outcome := Map.Support(move.From, move.To, move.Troops)
				fmt.Printf(outcome)
			} else {
				outcome := Map.Attack(move.From, move.To, move.Troops)
				fmt.Printf(outcome)
			}
		}
	}
}

func (p Player) GenerateMoves(Map gamemap.MapGraph, players []Player, teams [][]Player, time_limit float64) []Move {
	map_json, _          := json.Marshal(Map)
	map_json_string      := string(map_json)

	players_json, _      := json.Marshal(players)
	players_json_string  := string(players_json)

	teams_json, _        := json.Marshal(teams)
	teams_json_string    := string(teams_json)

	player_json, _       := json.Marshal(p)
	player_json_string   := string(player_json)

	var moves_json []byte

	if p.Code_language == "python3" {
		moves_json = p.executePython3(map_json_string, player_json_string, players_json_string, teams_json_string)
		// fmt.Printf("%s\n", moves_json)
	} else if p.Code_language == "go" {
		// TODO: create support for GOLANG AI
	} else if p.Code_language == "c++" {
		// TODO: create support for C++ AI
	} else (
		return []Move{}
	)

	var moves []Move
	if err := json.Unmarshal(moves_json, &moves); err != nil {
    log.Fatal(err)
  }
	moves = removeDuplicates(moves, len(Map.Bases))

	return moves
}

func (p Player) executePython3(map_json, player_json, players_json, teams_json string) []byte {
	cmd := exec.Command("python", p.Code_path + "run.py", map_json, player_json, players_json, teams_json)
	returned_result, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return returned_result
}

func removeDuplicates(list []Move, number_of_bases int) []Move {
	base_exists := make([]int, number_of_bases)
	new_list := []Move{}

	for index, item := range list {
		if base_exists[item.From] == 0 {
			base_exists[item.From]++
			new_list = append(new_list, item)
		}
	}

	return new_list
}

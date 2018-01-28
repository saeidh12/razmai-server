package gameplay
// TODO: Add time limit for the AI programm to run
import "../gamemap"
import "os/exec"
import "encoding/json"
import "log"
import "io/ioutil"

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
	Code_name     string
}

func (p Player) GenerateMoves(Map gamemap.MapGraph, players []Player, teams [][]int, time_limit float64, ais_folder string) []Move {
	map_json, _          := json.Marshal(Map)
	map_json_string      := string(map_json)

	players_json, _      := json.Marshal(players)
	players_json_string  := string(players_json)

	teams_json, _        := json.Marshal(teams)
	teams_json_string    := string(teams_json)

	player_json, _       := json.Marshal(p)
	player_json_string   := string(player_json)

	fileBytes, err := ioutil.ReadFile(ais_folder + p.Code_name + "/info.json")
	if err != nil {
		log.Fatal(err)
	}
	ai_info := make(map[string]string)
	if err = json.Unmarshal(fileBytes, &ai_info); err != nil {
		log.Fatal(err)
	}

	var moves_json []byte

	if ai_info["language"] == "python3" || ai_info["language"] == "go" || ai_info["language"] == "c++" {
		moves_json = p.execute(ai_info["command"], ai_info["file"], map_json_string, player_json_string, players_json_string, teams_json_string, ais_folder)
		// fmt.Printf("%s\n", moves_json)
	} else {
		log.Fatal("Language not supported!")
	}

	var moves []Move
	if err := json.Unmarshal(moves_json, &moves); err != nil {
    log.Fatal(err)
  }
	moves = removeDuplicates(moves, len(Map.Bases))

	return moves
}

func (p Player) execute(command, file, map_json, player_json, players_json, teams_json, ais_folder string) []byte {
	var cmd *exec.Cmd
	if command != "" {
		cmd = exec.Command(command, ais_folder + p.Code_name + file, map_json, player_json, players_json, teams_json)
	} else {
		cmd = exec.Command(ais_folder + p.Code_name + file, map_json, player_json, players_json, teams_json)
	}
	returned_result, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return returned_result
}

func removeDuplicates(list []Move, number_of_bases int) []Move {
	base_exists := make([]int, number_of_bases)
	new_list := []Move{}

	for _, item := range list {
		if base_exists[item.From] == 0 {
			base_exists[item.From]++
			new_list = append(new_list, item)
		}
	}

	return new_list
}

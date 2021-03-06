package gameplay

import (
  "testing"
  "reflect"
  "../gamemap"
  "regexp"
)
import "encoding/json"
import "log"
import "io/ioutil"

func CreatePlayerObject(player, team int) Player {
  return Player{
    Name: "Computer-py",
    Team_index: team,
    Player_index: player,
    Code_name: "Computer-py",
  }
}

func CreatePlayerObjectCpp(player, team int) Player {
  return Player{
    Name: "Computer-cpp",
    Team_index: team,
    Player_index: player,
    Code_name: "Computer-c++",
  }
}

func CreatePlayerObjectGolang(player, team int) Player {
  return Player{
    Name: "Computer-go",
    Team_index: team,
    Player_index: player,
    Code_name: "Computer-go",
  }
}

func TestPlayerRemoveDuplicate(t *testing.T) {
  list_of_moves := []Move{
    Move{0, 1, 5, "send"},
    Move{1, 2, 5, "send"},
    Move{2, 3, 5, "send"},
    Move{2, 4, 5, "send"},
    Move{3, 4, 5, "send"},
    Move{0, 2, 5, "send"},
  }

  expected_result := []Move{
    Move{0, 1, 5, "send"},
    Move{1, 2, 5, "send"},
    Move{2, 3, 5, "send"},
    Move{3, 4, 5, "send"},
  }

  actual_result := removeDuplicates(list_of_moves, 5)
	if !reflect.DeepEqual(actual_result, expected_result) {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
}

func TestPlayerExecute4Py(t *testing.T) {
  player_object := CreatePlayerObject(0, 0)

  turns_json      := CreateTurnsJSON()

  player_json   := CreatePlayerJSON("Computer-py")

  players_json  := CreatePlayersJSON("Computer-py")

  teams_json    := CreateTeamsJSON()

  expected_result := "[{\"From\": 0, \"To\": 2, \"Troop\": 0, \"Type\": \"send\"}]"

  fileBytes, err := ioutil.ReadFile(ais_folder + player_object.Code_name + "/info.json")
	if err != nil {
		log.Fatal(err)
	}
	ai_info := make(map[string]string)
	if err = json.Unmarshal(fileBytes, &ai_info); err != nil {
		log.Fatal(err)
	}

  actual_result := string(player_object.execute(ai_info["command"], ai_info["file"], turns_json, player_json, players_json, teams_json, ais_folder))
  if actual_result != expected_result {
		t.Fatalf("Expected %s but got %s", expected_result, actual_result)
	}
}

func TestPlayerExecute4Cpp(t *testing.T) {
  player_object := CreatePlayerObjectCpp(0, 0)

  turns_json      := CreateTurnsJSON()

  player_json   := CreatePlayerJSON("Computer-c++")

  players_json  := CreatePlayersJSON("Computer-c++")

  teams_json    := CreateTeamsJSON()

  expected_result := "[{\"From\":0,\"To\":2,\"Troop\":0,\"Type\":\"send\"}]"

  fileBytes, err := ioutil.ReadFile(ais_folder + player_object.Code_name + "/info.json")
	if err != nil {
		log.Fatal(err)
	}
	ai_info := make(map[string]string)
	if err = json.Unmarshal(fileBytes, &ai_info); err != nil {
		log.Fatal(err)
	}

  actual_result := string(player_object.execute(ai_info["command"], ai_info["file"], turns_json, player_json, players_json, teams_json, ais_folder))
  if actual_result != expected_result {
		t.Fatalf("Expected %s but got %s", expected_result, actual_result)
	}
}

func TestPlayerExecute4Golang(t *testing.T) {
  player_object := CreatePlayerObjectGolang(0, 0)

  turns_json      := CreateTurnsJSON()

  player_json   := CreatePlayerJSON("Computer-go")

  players_json  := CreatePlayersJSON("Computer-go")

  teams_json    := CreateTeamsJSON()

  expected_result := "[{\"From\":0,\"To\":2,\"Troop\":0,\"Type\":\"send\"}]"

  fileBytes, err := ioutil.ReadFile(ais_folder + player_object.Code_name + "/info.json")
	if err != nil {
		log.Fatal(err)
	}
	ai_info := make(map[string]string)
	if err = json.Unmarshal(fileBytes, &ai_info); err != nil {
		log.Fatal(err)
	}

  actual_result := string(player_object.execute(ai_info["command"], ai_info["file"], turns_json, player_json, players_json, teams_json, ais_folder))
  if actual_result != expected_result {
		t.Fatalf("Expected %s but got %s", expected_result, actual_result)
	}
}

func TestPlayerGenerateMoves(t *testing.T) {
  map_object            := gamemap.CreateMapGraphObjectFromFile("../../maps/beginners.json")

  player_object         := CreatePlayerObject(0, 0)

  oponent_player_object := CreatePlayerObject(1, 1)

  players               := []Player{player_object, oponent_player_object}
  teams                 := [][]int{[]int{0},[]int{1}}

  expected_result       := []Move{
    Move{
      From:   0,
      To:     2,
      Troops: 0,
      Type:   "send",
    },
  }
  actual_result := player_object.GenerateMoves(append([]gamemap.MapGraph{}, map_object), players, teams, 20.0, ais_folder)

  if !reflect.DeepEqual(actual_result, expected_result) {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
}

func CreateTurnsJSON() string {
  return "[" + CreateMapJSON() + "]"
}

func CreateMapJSON() string {
  re := regexp.MustCompile("\n[\t| ]*")
	return re.ReplaceAllString(`{
      "Number_of_players":            2,
      "Weak_delimiter":               5,
      "Medium_delimiter":             15,
      "Conquer_bonus":                4,
      "Bases":[
          {
              "Occupying_player":     0,
              "Troop_count":          2,
              "Attack_bonus":         1,
              "Defense_bonus":        1.2,
              "Troop_bonus":          3,
              "Connections":          [2, 3, 4],
              "X":                    -200,
              "Y":                    0
          },
          {
              "Occupying_player":     1,
              "Troop_count":          10,
              "Attack_bonus":         1,
              "Defense_bonus":        1.2,
              "Troop_bonus":          3,
              "Connections":          [8, 9, 10],
              "X":                    200,
              "Y":                    0
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [0, 5, 6],
              "X":                    -100,
              "Y":                    100
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [0, 6],
              "X":                    -100,
              "Y":                    0
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [0, 6, 7],
              "X":                    -100,
              "Y":                    -100
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [2, 8],
              "X":                    0,
              "Y":                    100
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1.1,
              "Defense_bonus":        1.1,
              "Troop_bonus":          2,
              "Connections":          [2, 3, 4, 8, 9, 10],
              "X":                    0,
              "Y":                    0
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [4, 10],
              "X":                    0,
              "Y":                    -100
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [1, 5, 6],
              "X":                    100,
              "Y":                    100
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [1, 6],
              "X":                    100,
              "Y":                    0
          },
          {
              "Occupying_player":     -1,
              "Troop_count":          0,
              "Attack_bonus":         1,
              "Defense_bonus":        1,
              "Troop_bonus":          2,
              "Connections":          [1, 6, 7],
              "X":                    100,
              "Y":                    -100
          }
      ],
      "Edges": [
                                      [0, 2],
                                      [0, 3],
                                      [0, 4],
                                      [2, 5],
                                      [2, 6],
                                      [3, 6],
                                      [4, 6],
                                      [4, 7],
                                      [5, 8],
                                      [6, 8],
                                      [6, 9],
                                      [6, 10],
                                      [7, 10],
                                      [8, 1],
                                      [9, 1],
                                      [10, 1]
      ]
  }`, "")
}

func CreatePlayerJSON(ai_name string) string {
  re := regexp.MustCompile("\n[\t| ]*")
	return re.ReplaceAllString(`{
    "Name":"Computer-py",
    "Team_index":0,
    "Player_index":0,
    "Code_name":"` + ai_name + `"
  }`, "")
}

func CreatePlayersJSON(ai_name string) string {
  re := regexp.MustCompile("\n[\t| ]*")
	return re.ReplaceAllString(`[
    {
      "Name":"Computer-py",
      "Team_index":0,
      "Player_index":0,
      "Code_name":"` + ai_name + `"
    },
    {
      "Name":"Computer-py",
      "Team_index":1,
      "Player_index":1,
      "Code_name":"` + ai_name + `"
    }
  ]`, "")
}

func CreateTeamsJSON() string {
  return "[[0],[1]]"
}

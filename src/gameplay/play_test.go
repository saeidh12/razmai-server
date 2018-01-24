package gameplay

import (
  "testing"
  "reflect"
  // "log"
  "../gamemap"
  // "io/ioutil"
  // "encoding/json"
)

const maps_folder = "../../maps/"
const ais_folder = "../../AIs/"

func CreateGameObject(map_file string) Game {
  map_object            := gamemap.CreateMapGraphObjectFromFile(map_file)
  player_object         := CreatePlayerObject(0, 0)
  oponent_player_object := CreatePlayerObject(1, 1)
  players               := []Player{player_object, oponent_player_object}
  teams                 := [][]int{[]int{0},[]int{1},}

  return Game{
    Map:                 map_object,
    Players:             players,
    Teams:               teams,
    Time_limit:          0.5,
  }
}

func TestPlayInit(t *testing.T) {
  expected_result := CreateGameObject(maps_folder + "beginners.json")

  actual_result   := Game{}
  actual_result.init(
    CreateMapJSON(),
    CreatePlayersJSON("Computer-py"),
    CreateTeamsJSON(),
    0.5,
  )

  if !reflect.DeepEqual(actual_result, expected_result) {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
}

func TestPlayGameEnded(t *testing.T) {
  game          := CreateGameObject(maps_folder + "beginners_ended.json")
  actual_result := game.GameEnded()
  if actual_result != true {
		t.Fatalf("Expected %v but got %v", true, actual_result)
	}

  game = CreateGameObject(maps_folder + "beginners_test.json")
  actual_result = game.GameEnded()
  if actual_result != false {
    t.Fatalf("Expected %v but got %v", false, actual_result)
  }
}

func TestPlayPlayerBaseCount(t *testing.T) {
  game            := CreateGameObject(maps_folder + "beginners_test.json")
  actual_result   := game.PlayerBaseCount(0)
  expected_result := 3
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  actual_result   = game.PlayerBaseCount(1)
  expected_result = 3
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }
}

func TestPlayTeamBaseCount(t *testing.T) {
  game            := CreateGameObject(maps_folder + "beginners_test.json")
  actual_result   := game.TeamBaseCount(0)
  expected_result := 3
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  actual_result   = game.TeamBaseCount(1)
  expected_result = 3
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }
}

func TestPlayPlayerLeaderBoard(t *testing.T) {
  game            := CreateGameObject(maps_folder + "beginners_ended.json")
  actual_result   := game.PlayerLeaderBoard()
  expected_result := []int{1, 0}
  if !reflect.DeepEqual(actual_result, expected_result) {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }
}

func TestPlayTeamLeaderBoard(t *testing.T) {
  game            := CreateGameObject(maps_folder + "beginners_ended.json")
  actual_result   := game.TeamLeaderBoard()
  expected_result := []int{1, 0}
  if !reflect.DeepEqual(actual_result, expected_result) {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }
}

func TestPlayPlayMoves(t *testing.T) {
  game            := CreateGameObject(maps_folder + "beginners.json")
  game.PlayMoves(0, ais_folder)

  actual_result   := game.PlayerBaseCount(0)
  expected_result := 2
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  actual_result   = game.PlayerBaseCount(1)
  expected_result = 1
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  actual_result   = game.PlayerBaseCount(-1)
  expected_result = 8
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  actual_result   = game.Map.Bases[2].Troop_count
  expected_result = 5 // plus conquer bonus
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  actual_result   = game.Map.Bases[0].Troop_count
  expected_result = 1
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }
}

func TestPlayPlayTurn(t *testing.T) {
  game            := CreateGameObject(maps_folder + "beginners.json")
  actual_result   := game.PlayTurn(0, ais_folder)
  expected_result := false
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  game            = CreateGameObject(maps_folder + "beginners_test.json")
  actual_result   = game.PlayTurn(0, ais_folder)
  expected_result = false
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

  game            = CreateGameObject(maps_folder + "beginners_ended.json")
  actual_result   = game.PlayTurn(0, ais_folder)
  expected_result = true
  if actual_result != expected_result {
    t.Fatalf("Expected %v but got %v", expected_result, actual_result)
  }

}

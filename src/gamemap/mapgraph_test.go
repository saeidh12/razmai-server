package gamemap

import "testing"
import "reflect"
import "regexp"

func CreateMapGraphObject() MapGraph {
	return MapGraph{
		Number_of_players: 2,
		Weak_delimiter:    5,
		Medium_delimiter:  10,
		Conquer_bonus:     4,
		Bases: []Base{
			Base{
				Occupying_player: 0,
				Troop_count:      10,
				Attack_bonus:     1.0,
				Defense_bonus:    1.1,
				Troop_bonus:      5,
				Connections:      []int{1, 2},
	      X:                -200,
	      Y:                -100,
			},
			Base{
				Occupying_player: 1,
				Troop_count:      10,
				Attack_bonus:     1.0,
				Defense_bonus:    1.1,
				Troop_bonus:      5,
				Connections:      []int{0, 2},
	      X:                200,
	      Y:                -100,
			},
			Base{
				Occupying_player: -1,
				Troop_count:      2,
				Attack_bonus:     1.0,
				Defense_bonus:    1,
				Troop_bonus:      3,
				Connections:      []int{0, 1},
	      X:                0,
	      Y:                100,
			},
		},
	}
}

func TestMapGraphCopyForPlayers(t *testing.T) {
	map_object := CreateMapGraphObject()

	players := []int{0}

	map_copy := map_object.CopyForPlayers(players)

	if map_copy.Bases[1].Troop_count != 2 {
		t.Fatalf("Expected %v but got %v", 2, map_copy.Bases[1].Troop_count)
	}

	if map_copy.Bases[2].Troop_count != 1 {
		t.Fatalf("Expected %v but got %v", 1, map_copy.Bases[2].Troop_count)
	}
}


func TestMapGraphOwnsBase(t *testing.T) {
	map_object := CreateMapGraphObject()

	player := 1
	actual_result := map_object.OwnsBase(1, player)
	if actual_result != true {
		t.Fatalf("Expected %v but got %v", true, actual_result)
	}

	actual_result = map_object.OwnsBase(0, player)
	if actual_result != false {
		t.Fatalf("Expected %v but got %v", false, actual_result)
	}
}

func TestMapGraphToJSON(t *testing.T) {
	map_object := CreateMapGraphObject()

	expected_result := CreateMapGraphJSONString()
	actual_result := map_object.ToJSON()

	if expected_result != actual_result {
		t.Fatalf("Expected \n%s\n but got \n%s", expected_result, actual_result)
	}
}

func TestMapGraphFromJSON(t *testing.T) {
	expected_result := CreateMapGraphObject()

	map_json := CreateMapGraphJSONString()
	actual_result := MapGraph{}
	actual_result.FromJSON(map_json)

	if !reflect.DeepEqual(actual_result, expected_result) {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
}

func TestMapGraphAddPlayerTroopBonus(t *testing.T) {
	map_object := CreateMapGraphObject()

	player := 1
	map_object.AddPlayerTroopBonus(player)

	if map_object.Bases[1].Troop_count != 15 {
		t.Fatalf("Expected %v but got %v", 15, map_object.Bases[1].Troop_count)
	}
	if map_object.Bases[2].Troop_count != 2 {
		t.Fatalf("Expected %v but got %v", 2, map_object.Bases[2].Troop_count)
	}
}

func TestMapGraphvalidateAttackSupportInput(t *testing.T) {
	map_object := CreateMapGraphObject()

	from, to, troops := 0, 0, 5
	expected_result := "No connection!"
	actual_result := validateAttackSupportInput(from, to, troops, map_object.Bases)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}

	from, to, troops = 2, 1, 5
	expected_result = "" // " F  F  F  F  F "
	actual_result = validateAttackSupportInput(from, to, troops, map_object.Bases)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}

	from, to, troops = 3, 1, 5
	expected_result = "Invalid move!"
	actual_result = validateAttackSupportInput(from, to, troops, map_object.Bases)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
}

func TestMapGraphAttack(t *testing.T) {
	map_object := CreateMapGraphObject()

	from, to, troops := 0, 2, 5
	expected_result := "" // " F  F  F  F  F "
	actual_result := map_object.Attack(map_object.Copy(), from, to, troops)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
	if map_object.Bases[from].Troop_count != 5 {
		t.Fatalf("Expected %v but got %v", 5, map_object.Bases[from].Troop_count)
	}
	if map_object.Bases[to].Troop_count != 7 {
		t.Fatalf("Expected %v but got %v", 7, map_object.Bases[to].Troop_count)
	}
	if map_object.Bases[to].Occupying_player != map_object.Bases[from].Occupying_player {
		t.Fatalf("Expected %v but got %v", map_object.Bases[to].Occupying_player, map_object.Bases[from].Occupying_player)
	}

	from, to, troops = 2, 1, 0
	expected_result = ""
	actual_result = map_object.Attack(map_object.Copy(), from, to, troops)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
	if map_object.Bases[from].Troop_count != 1 {
		t.Fatalf("Expected %v but got %v", 1, map_object.Bases[from].Troop_count)
	}
	if map_object.Bases[to].Troop_count != 5 {
		t.Fatalf("Expected %v but got %v", 5, map_object.Bases[to].Troop_count)
	}
	if map_object.Bases[to].Occupying_player != 1 {
		t.Fatalf("Expected %v but got %v", 1, map_object.Bases[to].Occupying_player)
	}

	from, to, troops = 1, 2, 1
	expected_result = ""
	actual_result = map_object.Attack(map_object.Copy(), from, to, troops)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
	if map_object.Bases[from].Troop_count != 4 {
		t.Fatalf("Expected %v but got %v", 4, map_object.Bases[from].Troop_count)
	}
	if map_object.Bases[to].Troop_count != 0 {
		t.Fatalf("Expected %v but got %v", 0, map_object.Bases[to].Troop_count)
	}
	if map_object.Bases[to].Occupying_player != -1 {
		t.Fatalf("Expected %v but got %v", -1, map_object.Bases[to].Occupying_player)
	}
}

func TestMapGraphSupport(t *testing.T) {
	map_object := CreateMapGraphObject()

	from, to, troops := 0, 2, 5
	expected_result := ""
	actual_result := map_object.Support(map_object.Copy(), from, to, troops)
	if expected_result != actual_result {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
	if map_object.Bases[from].Troop_count != 5 {
		t.Fatalf("Expected %v but got %v", 5, map_object.Bases[from].Troop_count)
	}
	if map_object.Bases[to].Troop_count != 7 {
		t.Fatalf("Expected %v but got %v", 7, map_object.Bases[to].Troop_count)
	}
	if map_object.Bases[to].Occupying_player != -1 {
		t.Fatalf("Expected %v but got %v", -1, map_object.Bases[to].Occupying_player)
	}
}

func TestMapGraphInSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 10}

	value := 5
	actual_value := InSlice(value, list)
	if actual_value != 4 {
		t.Fatalf("Expected %v but got %v", 4, actual_value)
	}

	value = 0
	actual_value = InSlice(value, list)
	if actual_value != -1 {
		t.Fatalf("Expected %v but got %v", -1, actual_value)
	}
}

func TestMapGraphInRange(t *testing.T) {
	a, b := 0, 20

	value := 5
	if !InRange(value, a, b) {
		t.Fatalf("Expected %v but got %v", true, false)
	}

	value = 0
	if InRange(value, a, b) {
		t.Fatalf("Expected %v but got %v", false, true)
	}

	value = 20
	if InRange(value, a, b) {
		t.Fatalf("Expected %v but got %v", false, true)
	}

	value = 25
	if InRange(value, a, b) {
		t.Fatalf("Expected %v but got %v", false, true)
	}
}

func TestMapGraphCreateMapGraphObjectFromFile(t *testing.T) {
	expected_result := CreateMapGraphObject()
	actual_result := CreateMapGraphObjectFromFile("../../maps/test.json")
	if !reflect.DeepEqual(actual_result, expected_result) {
		t.Fatalf("Expected %v but got %v", expected_result, actual_result)
	}
}

func CreateMapGraphJSONString() string {
	re := regexp.MustCompile("\n[\t| ]*")
	return re.ReplaceAllString(`{
		"Number_of_players":2,
		"Weak_delimiter":5,
		"Medium_delimiter":10,
		"Conquer_bonus":4,
		"Bases":[
			{
				"Occupying_player":0,
				"Troop_count":10,
				"Attack_bonus":1,
				"Defense_bonus":1.1,
				"Troop_bonus":5,
				"Connections":[1,2],
	      "X":-200,
	      "Y":-100
			},
			{
				"Occupying_player":1,
				"Troop_count":10,
				"Attack_bonus":1,
				"Defense_bonus":1.1,
				"Troop_bonus":5,
				"Connections":[0,2],
	      "X":200,
	      "Y":-100
			},
			{
				"Occupying_player":-1,
				"Troop_count":2,
				"Attack_bonus":1,
				"Defense_bonus":1,
				"Troop_bonus":3,
				"Connections":[0,1],
	      "X":0,
	      "Y":100
			}
		]
	}`, "")
}

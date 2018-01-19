package gamemap

import (
	"encoding/json"
	"log"
	"io/ioutil"
)

type MapGraph struct {
	Number_of_players int
	Weak_delimiter    int
	Medium_delimiter  int
	Conquer_bonus     int
	Bases             []Base
}

func (mg MapGraph) CopyForPlayers(players []int) MapGraph {
	mg_copy := MapGraph{
		Number_of_players: mg.Number_of_players,
		Weak_delimiter:    mg.Weak_delimiter,
		Medium_delimiter:  mg.Medium_delimiter,
		Conquer_bonus:     mg.Conquer_bonus,
		Bases:             make([]Base, len(mg.Bases)),
	}

	copy(mg_copy.Bases, mg.Bases)


	for index, base := range mg_copy.Bases {
		if InSlice(base.Occupying_player, players) == -1 {
			if mg_copy.Bases[index].Troop_count <= mg_copy.Weak_delimiter {
				mg_copy.Bases[index].Troop_count = 1
			} else if mg_copy.Bases[index].Troop_count <= mg_copy.Medium_delimiter {
				mg_copy.Bases[index].Troop_count = 2
			} else {
				mg_copy.Bases[index].Troop_count = 3
			}
		}
	}

	return mg_copy
}

func (mg MapGraph) Copy() MapGraph {
	mg_copy := MapGraph{
		Number_of_players: mg.Number_of_players,
		Weak_delimiter:    mg.Weak_delimiter,
		Medium_delimiter:  mg.Medium_delimiter,
		Conquer_bonus:     mg.Conquer_bonus,
		Bases:             make([]Base, len(mg.Bases)),
	}

	copy(mg_copy.Bases, mg.Bases)

	return mg_copy
}

func (mg MapGraph) OwnsBase(base int, player int) bool {
	return mg.Bases[base].Occupying_player == player
}

func (mg MapGraph) ToJSON() string {
	map_json, _ := json.Marshal(mg)
	return string(map_json)
}

func (mg *MapGraph) FromJSON(map_json string) {
	bytes := []byte(map_json)
	json.Unmarshal(bytes, mg)
}

func (mg *MapGraph) AddPlayerTroopBonus(player int) {
	for i, base := range mg.Bases {
		if base.Occupying_player == player {
			mg.Bases[i].Troop_count += base.Troop_bonus
		}
	}
}

func (mg *MapGraph) Attack(mg_copy MapGraph, from, to, troops int) string {
	if check := validateAttackSupportInput(from, to, troops, mg.Bases); check != "" {
		return check
	}

	troops = validateTroopCount(troops, mg_copy.Bases[from])
	mg.Bases[from].Troop_count -= troops

	attack_force := mg_copy.Bases[from].get_attack_strength(troops)
	defense_force := mg_copy.Bases[to].get_defense_strength()

	battle_result := attack_force - defense_force

	if battle_result > 0 {
		mg.Bases[to].Troop_count = battle_result + mg.Conquer_bonus
		mg.Bases[to].Occupying_player = mg.Bases[from].Occupying_player
	} else {
		battle_result *= -1
		if battle_result < mg.Bases[to].Troop_count {
			mg.Bases[to].Troop_count = battle_result
			if mg.Bases[to].Troop_count == 0 {
				mg.Bases[to].Occupying_player = -1
			}
		}
	}

	return ""
}

func (mg MapGraph) Support(mg_copy MapGraph, from, to, troops int) string {
	if check := validateAttackSupportInput(from, to, troops, mg.Bases); check != "" {
		return check
	}

	troops = validateTroopCount(troops, mg_copy.Bases[from])

	mg.Bases[from].Troop_count -= troops
	mg.Bases[to].Troop_count += troops

	return ""
}

func validateTroopCount(troops int, base Base) int {
	if troops >= base.Troop_count || troops == 0 {
		return base.Troop_count - 1
	}
	return troops
}

func validateAttackSupportInput(from, to, troops int, bases []Base) string {
	from_base_doesnt_exist := !InRange(from, -1, len(bases))
	if from_base_doesnt_exist {
		return "Invalid move!"
	}

	to_base_doesnt_exist := !InRange(to, -1, len(bases))
	if to_base_doesnt_exist {
		return "Invalid move!"
	}

	from_base_doesnt_have_enough_troops := bases[from].Troop_count < 2
	if from_base_doesnt_have_enough_troops {
		return "Invalid move!"
	}

	troops_negative := troops < 0
	if troops_negative {
		return "Invalid move!"
	}

	to_not_connected_to_from := InSlice(to, bases[from].Connections) < 0

	if to_not_connected_to_from {
		return "No connection!"
	}

	return ""
}

func CreateMapGraphObjectFromFile(map_file string) MapGraph {
  fileBytes, err := ioutil.ReadFile(map_file)
  if err != nil {
    log.Fatal(err)
  }
  map_object := MapGraph{}
  if err = json.Unmarshal(fileBytes, &map_object); err != nil {
    log.Fatal(err)
  }
  return map_object
}

func InSlice(value int, list []int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}

	return -1
}

func InRange(value, a, b int) bool { return value > a && value < b }

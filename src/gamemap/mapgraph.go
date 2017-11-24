package gamemap

import "encoding/json"

type MapGraph struct {
	Number_of_players int
	Weak_delimiter    int
	Medium_delimiter  int
	Conquer_bonus     int
	Bases             []Base
}

func (mg MapGraph) toJSON() string {
	map_json, _ := json.Marshal(mg)
	return string(map_json)
}

func (mg *MapGraph) fromJSON(map_json string) {
	bytes := []byte(map_json)
	json.Unmarshal(bytes, mg)
}

func (mg *MapGraph) attack(from int, to int, troops int) string {
	if mg.Bases[from].Troop_count < 2 || !inRange(from, -1, len(mg.Bases)) || !inRange(to, -1, len(mg.Bases)) || troops < 0 || from == to {
		return "Invalid move!"
	}
	if inSlice(to, mg.Bases[from].Connections) < 0 {
		return "No connection!"
	}

	if troops >= mg.Bases[from].Troop_count {
		troops = mg.Bases[from].Troop_count - 1
	}
	mg.Bases[from].Troop_count -= troops

	attack_force := mg.Bases[from].get_attack_strength(troops)
	defense_force := mg.Bases[to].get_defense_strength()

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

func (mg MapGraph) support(from int, to int, troops int) string {
	if mg.Bases[from].Troop_count < 2 || !inRange(from, -1, len(mg.Bases)) || !inRange(to, -1, len(mg.Bases)) || troops < 0 {
		return "Incorrect input"
	}
	if inSlice(to, mg.Bases[from].Connections) < 0 {
		return "No connection"
	}

	if troops >= mg.Bases[from].Troop_count {
		troops = mg.Bases[from].Troop_count - 1
	}
	mg.Bases[from].Troop_count -= troops
	mg.Bases[to].Troop_count += troops

	return ""
}

func inSlice(value int, list []int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}

	return -1
}

func inRange(value, a, b int) bool { return value > a && value < b }

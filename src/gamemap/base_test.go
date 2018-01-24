package gamemap

import "testing"

func TestBase(t *testing.T) {
	base := Base{Occupying_player: 1, Troop_count: 20, Attack_bonus: 1.3, Defense_bonus: 1.1, Connections: []int{1, 2, 3}}

	attack_strength := base.get_attack_strength(10)

	if attack_strength != 13 {
		t.Fatalf("Expected %v but got %v", 13, attack_strength)
	}

	defense_strength := base.get_defense_strength()

	if defense_strength != 22 {
		t.Fatalf("Expected %v but got %v", 2, defense_strength)
	}
}

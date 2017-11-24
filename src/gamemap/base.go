package gamemap

type Base struct {
	Occupying_player int
	Troop_count      int
	Attack_bonus     float64
	Defense_bonus    float64
	Troop_bonus      int
	Connections      []int
}

func (b Base) get_attack_strength(troops int) int { return int(b.Attack_bonus * float64(troops)) }
func (b Base) get_defense_strength() int          { return int(b.Defense_bonus * float64(b.Troop_count)) }

package gamemap

type Base struct {
	occupying_player int
	troop_count      int
	attack_bonus     float64
	defense_bonus    float64
}

func (b Base) get_occupying_player() int  { return b.occupying_player }
func (b Base) get_troop_count() int       { return b.troop_count }
func (b Base) get_attack_bonus() float64  { return b.attack_bonus }
func (b Base) get_defense_bonus() float64 { return b.defense_bonus }

func (b *Base) set_occupying_player(occupying_player int) { b.occupying_player = occupying_player }
func (b *Base) set_troop_count(troop_count int)           { b.troop_count = troop_count }
func (b *Base) set_attack_bonus(attack_bonus float64)     { b.attack_bonus = attack_bonus }
func (b *Base) set_defense_bonus(defense_bonus float64)   { b.defense_bonus = defense_bonus }

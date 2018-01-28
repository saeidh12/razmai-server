package main

type Player struct {
	Name          string
	Team_index    int
	Player_index  int
	// AI_specs
	Code_name     string
}

type MapGraph struct {
	Number_of_players int
	Weak_delimiter    int
	Medium_delimiter  int
	Conquer_bonus     int
	Bases             []Base
}

type Base struct {
	Occupying_player int
	Troop_count      int
	Attack_bonus     float64
	Defense_bonus    float64
	Troop_bonus      int
	Connections      []int
	X                int
	Y                int
}

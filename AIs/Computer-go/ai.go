package main

func TargetPlayers(Players []Player, my_team int) []int {
  var target_players []int
  target_players = append(target_players, -1)
  for index, player := range Players {
    if player.Team_index != my_team {
      target_players = append(target_players, index)
    }
  }
  return target_players
}

func Commander(mapgraph MapGraph, player Player, players []Player, teams [][]int) []map[string]interface{}  {
  target_players := TargetPlayers(players, player.Team_index)

  var moves []map[string]interface{}
  for index, value := range mapgraph.Bases {
    if value.Occupying_player == player.Player_index {
      path := BFS(index, mapgraph.Bases, target_players)
      if len(path) > 1 {
        var m = map[string]interface{}{
        	"From":    index,
        	"To":      path[1],
          "Troop":   0,
          "Type":    "send",
        }
        moves = append(moves, m)
      }
    }
  }

  return moves
}

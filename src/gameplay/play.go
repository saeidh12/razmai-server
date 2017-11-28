package gameplay

import "../gamemap"
import "encoding/json"
import "log"
import "sort"
import "fmt"

type Game struct {
  Map                 gamemap.MapGraph
  Players             []Player
  Teams               [][]int
  Time_limit          float64
  Player_turn         int
}

func (game *Game) init(
  map_json_string,
  players_json_string,
  teams_json_string string,
  time_limit float64,
  player_turn int) {
  if err := json.Unmarshal([]byte(map_json_string), &game.Map); err != nil {
    log.Fatal(err)
  }

  if err := json.Unmarshal([]byte(players_json_string), &game.Players); err != nil {
    log.Fatal(err)
  }

  if err := json.Unmarshal([]byte(teams_json_string), &game.Teams); err != nil {
    log.Fatal(err)
  }

  game.Time_limit          = time_limit
  game.Player_turn         = player_turn
}

func (game *Game) PlayTurn() bool {
  if !game.GameEnded() && game.PlayerBaseCount(game.Player_turn) > 0 {
    game.Map.AddPlayerTroopBonus(game.Player_turn)
    game.PlayMoves()
    return game.GameEnded()
  }
  return game.GameEnded()
}

func (game *Game) PlayMoves() {
  current_player := game.Players[game.Player_turn]
  team_turn := game.Teams[current_player.Team_index]
  moves := current_player.GenerateMoves(game.Map.CopyForPlayers(team_turn), game.Players, game.Teams, game.Time_limit)

  for _, move := range moves {
		if game.Map.OwnsBase(move.From, game.Player_turn) && move.Type == "attack" {
      if outcome := game.Map.Attack(move.From, move.To, move.Troops); outcome != "" {
        fmt.Println(outcome)
      }
		} else if game.Map.OwnsBase(move.From, game.Player_turn) && move.Type == "support" {
      if outcome := game.Map.Support(move.From, move.To, move.Troops); outcome != "" {
        fmt.Println(outcome)
      }
		} else if game.Map.OwnsBase(move.From, game.Player_turn) && move.Type == "send" {
			if game.Map.OwnsBase(move.To, game.Player_turn) {
        if outcome := game.Map.Support(move.From, move.To, move.Troops); outcome != "" {
          fmt.Println(outcome)
        }
			} else {
				if outcome := game.Map.Attack(move.From, move.To, move.Troops); outcome != "" {
          fmt.Println(outcome)
        }
			}
		}
	}
}

func (game Game) TeamLeaderBoard() []int {
  team_leader_board := []int{}
  for i := 0; i < len(game.Teams); i++ {
    team_leader_board = append(team_leader_board, i)
  }

  sort.Slice(team_leader_board, func(i, j int) bool { return game.TeamBaseCount(team_leader_board[i]) > game.TeamBaseCount(team_leader_board[j]) })
  return team_leader_board
}

func (game Game) PlayerLeaderBoard() []int {
  player_leader_board := []int{}
  for i := 0; i < len(game.Players); i++ {
    player_leader_board = append(player_leader_board, i)
  }

  sort.Slice(player_leader_board, func(i, j int) bool { return game.PlayerBaseCount(player_leader_board[i]) > game.PlayerBaseCount(player_leader_board[j]) })
  return player_leader_board
}

func (game Game) TeamBaseCount(team int) int {
  team_base_count := 0
  for _, player := range game.Teams[team] {
    team_base_count += game.PlayerBaseCount(player)
  }
  return team_base_count
}

func (game Game) PlayerBaseCount(player int) int {
  player_base_count := 0

  for _, base := range game.Map.Bases {
    if base.Occupying_player == player {
      player_base_count++
    }
  }
  return player_base_count
}

func (game Game) GameEnded() bool {
  number_of_active_teams := 0
  for i := 0; i < len(game.Teams); i++ {
    if game.TeamBaseCount(i) > 0 {
      if number_of_active_teams > 0 {
        return false
      } else {
        number_of_active_teams++
      }
    }
  }
  return true
}

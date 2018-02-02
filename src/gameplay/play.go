package gameplay

import "../gamemap"
import "encoding/json"
import "log"
import "sort"
import "fmt"

type Game struct {
  Turns           []gamemap.MapGraph
  Players         []Player
  Teams           [][]int
  Time_limit      float64
  Max_moves       int
}

func (game *Game) init(
  turns_json_string,
  players_json_string,
  teams_json_string string,
  time_limit float64) {
  if err := json.Unmarshal([]byte(turns_json_string), &game.Turns); err != nil {
    log.Fatal(err)
  }

  if err := json.Unmarshal([]byte(players_json_string), &game.Players); err != nil {
    log.Fatal(err)
  }

  if err := json.Unmarshal([]byte(teams_json_string), &game.Teams); err != nil {
    log.Fatal(err)
  }

  game.Time_limit          = time_limit
}

func (game *Game) PlayTurn(Player_turn int, ais_folder string) bool {
  if !game.GameEnded() && game.PlayerBaseCount(Player_turn) > 0 {
    game.PlayMoves(Player_turn, ais_folder)
    return game.GameEnded()
  }
  return game.GameEnded()
}

func (game *Game) PlayMoves(Player_turn int, ais_folder string) {
  turn           := len(game.Turns)
  current_player := game.Players[Player_turn]
  moves          := current_player.GenerateMoves(game.Turns4player(current_player), game.Players, game.Teams, game.Time_limit, ais_folder)
  mg_copy        := game.Turns[turn - 1].Copy()
  game.Turns      = append(game.Turns, mg_copy)
  game.Turns[turn].AddPlayerTroopBonus(Player_turn)

  for _, move := range moves {
    if game.Turns[turn].OwnsBase(move.From, Player_turn) {
      if move.Type == "attack" {
        if outcome := game.Turns[turn].Attack(mg_copy, move.From, move.To, move.Troops); outcome != "" {
          fmt.Println(outcome)
        }
  		} else if move.Type == "support" {
        if outcome := game.Turns[turn].Support(mg_copy, move.From, move.To, move.Troops); outcome != "" {
          fmt.Println(outcome)
        }
  		} else if move.Type == "send" {
  			if game.Turns[turn].OwnsBase(move.To, Player_turn) {
          if outcome := game.Turns[turn].Support(mg_copy, move.From, move.To, move.Troops); outcome != "" {
            fmt.Println(outcome)
          }
  			} else {
  				if outcome := game.Turns[turn].Attack(mg_copy, move.From, move.To, move.Troops); outcome != "" {
            fmt.Println(outcome)
          }
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
  turn           := len(game.Turns) - 1

  for _, base := range game.Turns[turn].Bases {
    if base.Occupying_player == player {
      player_base_count++
    }
  }
  return player_base_count
}

func (game Game) GameEnded() bool {
  if len(game.Turns) == game.Max_moves {
    return true
  }
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

func (game Game) Turns4player(current_player Player) []gamemap.MapGraph {
  team_turn      := game.Teams[current_player.Team_index]
  var newTurns []gamemap.MapGraph
  for _, Map := range game.Turns {
    newTurns = append(newTurns, Map.CopyForPlayers(team_turn))
  }
  return newTurns
}

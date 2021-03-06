package main

import (
  "encoding/json"
  "log"
  "net/http"
  "io/ioutil"
  "fmt"
  "github.com/rs/cors"
  "./gamemap"
  "./gameplay"
  "os"
)

var maps_folder = "./maps/"
var ais_folder  = "./AIs/"

func testPost(w http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var game gameplay.Game
  err := decoder.Decode(&game)
  if err != nil {
      panic(err)
  }
  defer req.Body.Close()
  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(game)
}

func MapsList(w http.ResponseWriter, req *http.Request) {
  response := make(map[string]gamemap.MapGraph)

  files, err := ioutil.ReadDir(maps_folder)
  if err != nil {
    log.Fatal(err)
  }
  for _, f := range files {
    name_length := len(f.Name()) - 5
    response[f.Name()[:name_length]] = gamemap.CreateMapGraphObjectFromFile(maps_folder + f.Name())
  }

  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(response)
}

func AIsList(w http.ResponseWriter, req *http.Request) {
  files, err := ioutil.ReadDir(ais_folder)
  if err != nil {
    log.Fatal(err)
  }
  AIs := []string{}
  for _, f := range files {
    AIs = append(AIs, f.Name())
  }

  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(AIs)
}

func TestConnection(w http.ResponseWriter, req *http.Request) {
  response := true
  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(response)
}

func HomePage(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "<h1 style=\"text-align: center; width: 100%%; color: cornflowerblue;\">This is the RAZMAI game server</h1>")
}

func PlayTurn(w http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var data struct {
    Game        gameplay.Game `json:"game"`
    Player_turn int
  }
  err := decoder.Decode(&data)
  if err != nil {
      panic(err)
  }
  defer req.Body.Close()

  gameEnded                      := data.Game.PlayTurn(data.Player_turn, ais_folder)
  response                       := make(map[string]interface{})
  response["newturn"]             = data.Game.Turns[len(data.Game.Turns) - 1]
  response["game_ended"]          = gameEnded
  response["player_leader_board"] = data.Game.PlayerLeaderBoard()
  response["team_leader_board"]   = data.Game.TeamLeaderBoard()


  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(response)
}

func main() {
  args := os.Args

  if len(args) == 3 {
    maps_folder = args[1]
    ais_folder  = args[2]
  }
  mux := http.NewServeMux()
  mux.HandleFunc("/",                HomePage)
  mux.HandleFunc("/maps",            MapsList)
  mux.HandleFunc("/ais",             AIsList)
  mux.HandleFunc("/play-turn",       PlayTurn)
  mux.HandleFunc("/test-connection", TestConnection)
  mux.HandleFunc("/test-post",       testPost)

  handler := cors.Default().Handler(mux)
  fmt.Println("Server running on \"http://localhost:8012/\"")
  log.Fatal(http.ListenAndServe(":8012", handler))
}


// func save() error {
//     filename := p.Title + ".txt"
//     return ioutil.WriteFile(filename, p.Body, 0600)
// }

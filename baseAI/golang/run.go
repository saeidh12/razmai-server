package main

import (
  "encoding/json"
  "io/ioutil"
  "os"
)

func main() {
  args := os.Args
  var mapgraph MapGraph
  var player Player
  var players []Player
  var teams [][]int

  if len(args) == 5 {
    err := json.Unmarshal([]byte(args[1]), &mapgraph)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal([]byte(args[2]), &player)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal([]byte(args[3]), &players)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal([]byte(args[4]), &teams)
    if err != nil {
        panic(err)
    }

    j, err := json.Marshal(Commander(mapgraph, player, players, teams))
    if err != nil {
        panic(err)
    }
    os.Stdout.Write(j)
  }
}

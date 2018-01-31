package main

import (
  "encoding/json"
  "io/ioutil"
  "os"
)

func main() {
  args := os.Args
  var turns []MapGraph
  var player Player
  var players []Player
  var teams [][]int

  if len(args) == 5 {
    err := json.Unmarshal([]byte(args[1]), &turns)
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

  } else {
    var mapgraph MapGraph
    fileBytes, err := ioutil.ReadFile(args[1])
    if err != nil {
      panic(err)
    }
    if err = json.Unmarshal(fileBytes, &mapgraph); err != nil {
      panic(err)
    }

    turns = append([]MapGraph{}, mapgraph)

    err = json.Unmarshal(
      []byte(`{"Name":"Computer-go","Team_index":0,"Player_index":0,"Code_path":"AIs/Computer-go/"}`),
      &player)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(
      []byte(`[{"Name":"Computer-go","Team_index":0,"Player_index":0,"Code_path":"AIs/Computer-go/"},{"Name":"Computer-go","Team_index":1,"Player_index":1,"Code_path":"AIs/Computer-go/"}]`),
      &players)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(
      []byte(`[[0],[1]]`),
      &teams)
    if err != nil {
        panic(err)
    }
  }
  j, err := json.Marshal(Commander(turns, player, players, teams))
  if err != nil {
      panic(err)
  }
  os.Stdout.Write(j)
}

package main

import (
  "encoding/json"
  "log"
  "net/http"
  "io/ioutil"
  "fmt"
  "github.com/rs/cors"
  "./gamemap"
)

// func test(rw http.ResponseWriter, req *http.Request) {
//   req.ParseForm()
//   log.Println(req.Form)
//   //LOG: map[{"test": "that"}:[]]
//   var t test_struct
//   for key, _ := range req.Form {
//     log.Println(key)
//     //LOG: {"test": "that"}
//     err := json.Unmarshal([]byte(key), &t)
//     if err != nil {
//       log.Println(err.Error())
//     }
//   }
//   log.Println(t.Test)
//   //LOG: that
// }

func MapsList(w http.ResponseWriter, req *http.Request) {
  response := make(map[string]gamemap.MapGraph)

  files, err := ioutil.ReadDir("./maps")
  if err != nil {
    log.Fatal(err)
  }
  for _, f := range files {
    name_length := len(f.Name()) - 5
    response[f.Name()[:name_length]] = gamemap.CreateMapGraphObjectFromFile("./maps/" + f.Name())
  }

  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(response)
}

func AIsList(w http.ResponseWriter, req *http.Request) {
  files, err := ioutil.ReadDir("./AIs")
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

func main() {
  // http.HandleFunc("/", MapsandAIsList)
  // http.HandleFunc("/test-connection", TestConnection)
  // // http.HandleFunc("/play", test)
  // // http.HandleFunc("/add-ai/file", test)
  // // http.HandleFunc("/add-map/file", test)
  // // http.HandleFunc("/add-map/json", test)
  mux := http.NewServeMux()
  // TODO: put home link to say this is the server
  mux.HandleFunc("/maps", MapsList)
  mux.HandleFunc("/ais", AIsList)
  mux.HandleFunc("/test-connection", TestConnection)

  handler := cors.Default().Handler(mux)
  fmt.Println("Server running on \"http://localhost:8012/\"")
  log.Fatal(http.ListenAndServe(":8012", handler))
}


// func save() error {
//     filename := p.Title + ".txt"
//     return ioutil.WriteFile(filename, p.Body, 0600)
// }

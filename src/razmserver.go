package main

import (
  "encoding/json"
  "log"
  "net/http"
  "io/ioutil"
  "fmt"
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

func MapsandAIsList(w http.ResponseWriter, req *http.Request) {
  response := make(map[string][]string)


  files, err := ioutil.ReadDir("./AIs")
  if err != nil {
    log.Fatal(err)
  }
  AIs := []string{}
  for _, f := range files {
    AIs = append(AIs, f.Name())
  }
  response["AIs"] = AIs

  files, err = ioutil.ReadDir("./maps")
  if err != nil {
    log.Fatal(err)
  }
  maps := []string{}
  for _, f := range files {
    name_length := len(f.Name()) - 5
    maps = append(maps, f.Name()[:name_length])
  }
  response["maps"] = maps

  json_encoder := json.NewEncoder(w)
  json_encoder.Encode(response)
}

func main() {
  http.HandleFunc("/", MapsandAIsList)
  // http.HandleFunc("/play", test)
  // http.HandleFunc("/add-ai/file", test)
  // http.HandleFunc("/add-map/file", test)
  // http.HandleFunc("/add-map/json", test)
  fmt.Println("Server running on \"localhost:8012\"")
  log.Fatal(http.ListenAndServe(":8012", nil))
}


// func save() error {
//     filename := p.Title + ".txt"
//     return ioutil.WriteFile(filename, p.Body, 0600)
// }

# razmai-server
This is the server which plays the RazmAI game. There is a gui to be used with this [here](https://github.com/saeidh12/razmai-gui).

## Prerequesets

* You must have golang installed. You can find instuctions [here](https://golang.org/).
* [Go CORS handler](https://github.com/rs/cors) `go get github.com/rs/cors`

## Installation and Running

1. Clone the repository
2. Go into the project directory
3. Run the razmserver: `go run src/razmserver.go`

you should see:

`Server running on "http://localhost:8012/"`

## Endpoints

### /
This route just displays a simple welcome page if the server is running.

### /maps
This route sends back a json object with the keys being the map names available in the server and the values being the map json itself.

### /ais
This route returns the list of the names of AIs available by the server.

### /play-turn
This route receives a json containing the the game object and the index of the player whose turn it is.
The game object is described in the [Wiki](https://github.com/saeidh12/razmai-server/wiki/Game-Object).

The server plays the turn and returns;
* the new game object
* game_ended: a boolean indicating if the game has finished
* player_leader_board: a list of players ordered by the number of bases they own
* team_leader_board: a list of teams ordered by the number of bases they own

### /test-connection
This route just returns a boolean to indicate the server is working properly.

## Adding AIs
You can add your AIs to the server by moving their file to the ais directory in the project.
Read the [wiki](https://github.com/saeidh12/razmai-server/wiki/Creating-Custom-AI) on how to create your own AI.

Currently supported languages:
* Python3
* C++
* Go

## Adding Maps
You can add your Maps to the server by moving their json file to the maps directory in the project.
Read the [wiki](https://github.com/saeidh12/razmai-server/wiki/Creating-Custom-Map) on how to create your own map.


## TODOs
* Tournament endpoints
* Implement time limit
* Add max number of moves to determine draw
* Add option to save and replay a game

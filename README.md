This is the server which plays the RazmAI game.

## Prerequesets

You must have golang installed. You can find instuctions [here](https://golang.org/).

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
This route sends back a json object with the keays being the map names available in the server and the values being the map json itself.

### /ais
This route returns the list of the names of AIs available by the server.

### /play-turn
This route recieves a json containing the the game object and the index of the player whose turn it is.
The game object is described in the [Wiki](https://github.com/saeidh12/razmai-server/wiki).

The server plays the turn and returns;
* the new game object
* game_ended: a boolean indecating if the game has finished
* player_leader_board: a list of players ordered by the number of bases they own
* team_leader_board: a list of teams ordered by the number of bases they own

### /test-connection
This route just returns a boolean to endicate the server is working properly.






















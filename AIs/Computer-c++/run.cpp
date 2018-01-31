#include <iostream>
#include <fstream>
#include "json.hpp"
using namespace std;
using json = nlohmann::json;
#include "ai.hpp"

int main(int argc, char** argv) {
  json Turns, Player, Players, Teams;

  if (argc == 5) {
    Turns   = json::parse(argv[1]);
    Player  = json::parse(argv[2]);
    Players = json::parse(argv[3]);
    Teams   = json::parse(argv[4]);
  }
  else {
    ifstream ifs (argv[1]);
    string Map_string( (istreambuf_iterator<char>(ifs) ),
                    (istreambuf_iterator<char>()    ) );
    // cout << Map_string << endl;
    json Map = json::parse(Map_string);
    Turns.push_back(Map);
    Player   = R"({"Name":"Computer-cpp","Team_index":0,"Player_index":0,"Code_path":"AIs/Computer-c++/"})"_json;
    Players  = R"([{"Name":"Computer-cpp","Team_index":0,"Player_index":0,"Code_path":"AIs/Computer-c++/"},{"Name":"Computer-cpp","Team_index":1,"Player_index":1,"Code_path":"AIs/Computer-c++/"}])"_json;
    Teams    = R"([[0],[1]])"_json;
  }

  json moves = Commander(Turns, Player, Players, Teams);

  cout << moves;

  return 0;
}

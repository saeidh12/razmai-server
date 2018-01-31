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

    json moves = Commander(Turns, Player, Players, Teams);

    cout << moves;
  }
  return 0;
}

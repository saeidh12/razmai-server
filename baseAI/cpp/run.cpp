#include <iostream>
#include <fstream>
#include "json.hpp"
using namespace std;
using json = nlohmann::json;
#include "ai.hpp"


int main(int argc, char** argv) {
  json Map, Player, Players, Teams;


  if (argc == 5) {
    Map     = json::parse(argv[1]);
    Player  = json::parse(argv[2]);
    Players = json::parse(argv[3]);
    Teams   = json::parse(argv[4]);

    json moves = Commander(Map, Player, Players, Teams);

    cout << moves;
  }
  return 0;
}

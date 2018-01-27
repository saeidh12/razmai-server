#include <iostream>
#include <fstream>
#include "json.hpp"
using namespace std;
using json = nlohmann::json;
#include "ai.hpp"
// #include <stdio.h>
// #ifdef WINDOWS
// #include <direct.h>
// #define GetCurrentDir _getcwd
// #else
// #include <unistd.h>
// #define GetCurrentDir getcwd
// #endif

int main(int argc, char** argv) {
  json Map, Player, Players, Teams;
  // char buff[FILENAME_MAX];
  // GetCurrentDir( buff, FILENAME_MAX );
  // string current_working_dir(buff);

  // cout << current_working_dir << endl;

  if (argc == 5) {
    Map     = json::parse(argv[1]);
    Player  = json::parse(argv[2]);
    Players = json::parse(argv[3]);
    Teams   = json::parse(argv[4]);
  }
  else {
    ifstream ifs (argv[1]);
    string Map_string( (istreambuf_iterator<char>(ifs) ),
                    (istreambuf_iterator<char>()    ) );
    // cout << Map_string << endl;
    Map = json::parse(Map_string);
    Player  = R"({"Name":"Computer-py","Team_index":0,"Player_index":0,"Code_language":"python3","Code_path":"AIs/Computer-py/"})"_json;
    Players = R"([{"Name":"Computer-py","Team_index":0,"Player_index":0,"Code_language":"python3","Code_path":"AIs/Computer-py/"},{"Name":"Computer-py","Team_index":1,"Player_index":1,"Code_language":"python3","Code_path":"AIs/Computer-py/"}])"_json;
    Teams   = R"([[0],[1]])"_json;
  }

  json moves = Commander(Map, Player, Players, Teams);

  cout << moves;

  return 0;
}

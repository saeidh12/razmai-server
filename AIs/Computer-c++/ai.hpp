#ifndef AI_HPP
#define AI_HPP

#include "mysearches.hpp"

json TargetPlayers(json players_json, json my_team) {
  json target_players;
  target_players.push_back(-1);
  json::iterator i;
  for (i = players_json.begin(); i != players_json.end(); i++) {
    json player = *i;
    if (player["Team_index"] != my_team)
      target_players.push_back(player["Team_index"]);
  }
  return target_players;
}

json Commander(json map_json, json player_json, json players_json, json teams_json) {
  int my_team = player_json["Team_index"];
  json target_players = TargetPlayers(players_json, my_team);

  json moves;
  json bases = map_json["Bases"];
  for (int i = 0; i < bases.size(); i++)
    if (bases[i]["Occupying_player"] == player_json["Player_index"]) {
      json path = BFS(i, bases, target_players);
      if (path.size() > 1) {
        int from_base_index    = i;
        int to_base_index      = path[1];
        int troops             = 0;
        string type_of_command = "send";
        json j = {
          {"From", from_base_index},
          {"To", to_base_index},
          {"Troop", troops},
          {"Type", type_of_command}
        };
        moves.push_back(j);
      }
    }

  return moves;
}

#endif

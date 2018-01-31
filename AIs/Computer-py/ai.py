from mysearches import bfs

def Commander(turns_list, player_dict, players_list, teams_list):
    current_turn   = len(turns_list) - 1
    target_players = [-1,] + [player["Player_index"] for player in players_list if player["Team_index"] != player_dict["Team_index"]]

    bases = turns_list[current_turn]["Bases"]
    moves = []
    for index, base in enumerate(bases):
        if base["Occupying_player"] == player_dict["Player_index"]:
            path = bfs(base, bases, target_players)
            if path:
                from_base_index = index
                to_base_index   = path[1]
                troops          = 0 # base["Troop_count"]
                type_of_command = "send"# if map_dict["Bases"][to_base_index]["Occupying_player"] in target_players else "support"
                moves.append({"From": from_base_index, "To": to_base_index, "Troop": troops, "Type": type_of_command})
    return moves

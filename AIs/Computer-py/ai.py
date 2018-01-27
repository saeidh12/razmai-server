from mysearches import bfs

def Commander(map_dict, player_dict, players_list, teams_list):
    target_players = [-1,] + [player["Player_index"] for player in players_list if player["Team_index"] != player_dict["Team_index"]]

    moves = []
    for index, base in enumerate(map_dict["Bases"]):
        if base["Occupying_player"] == player_dict["Player_index"]:
            path = bfs(base, map_dict["Bases"], target_players)
            if path:
                from_base_index = index
                to_base_index   = path[1]
                troops          = 0 # base["Troop_count"]
                type_of_command = "send"# if map_dict["Bases"][to_base_index]["Occupying_player"] in target_players else "support"
                moves.append({"From": from_base_index, "To": to_base_index, "Troop": troops, "Type": type_of_command})
    return moves

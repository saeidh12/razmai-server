import json
import sys
from ai import Commander

def run(argv):
    if len(argv) == 5:
        Map      = json.loads(argv[1])
        Player   = json.loads(argv[2])
        Players  = json.loads(argv[3])
        Teams    = json.loads(argv[4])

        moves = Commander(Map, Player, Players, Teams)
        sys.stdout.write(json.dumps(moves))
    else:
        with open("./maps/beginners_test.json", encoding='utf-8-sig') as json_file:
            text = json_file.read()
            Map = json.loads(text)
        Player   = json.loads('{"Name":"Computer-py","Team_index":0,"Player_index":0,"Code_language":"python3","Code_path":"AIs/Computer-py/"}')
        Players  = json.loads('[{"Name":"Computer-py","Team_index":0,"Player_index":0,"Code_language":"python3","Code_path":"AIs/Computer-py/"},{"Name":"Computer-py","Team_index":1,"Player_index":1,"Code_language":"python3","Code_path":"AIs/Computer-py/"}]')
        Teams    = json.loads('[[0],[1]]')

        moves = Commander(Map, Player, Players, Teams)
        print(moves)

run(sys.argv)

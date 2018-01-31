import json
import sys
from ai import Commander

def run(argv):
    if len(argv) == 5:
        Turns    = json.loads(argv[1])
        Player   = json.loads(argv[2])
        Players  = json.loads(argv[3])
        Teams    = json.loads(argv[4])

        moves = Commander(Turns, Player, Players, Teams)
        sys.stdout.write(json.dumps(moves))
    else:
        pass

run(sys.argv)

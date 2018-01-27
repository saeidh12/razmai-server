from myqueue import CheckableQueue

def bfs(start, bases, players):
    # a FIFO open_set
    open_set = CheckableQueue()
    # an empty set to maintain visited nodes
    closed_set = set()
    # a dictionary to maintain meta information (used for path formation)
    meta = dict()  # key -> (parent state, action to reach child)

    # initialize
    start = bases.index(start)
    meta[start] = None
    open_set.put(start)

    while not open_set.empty():

        parent_state = open_set.get()

        if is_goal(bases[parent_state], players):
              return construct_path(parent_state, meta)

        for child_state in bases[parent_state]["Connections"]:

            if child_state in closed_set:
                continue

            if child_state not in open_set:
                meta[child_state] = parent_state
                open_set.put(child_state)

        closed_set.add(parent_state)

def construct_path(state, meta):
    action_list = list()

    while True:
        action_list.insert(0, state)
        row = meta.get(state)
        if row != None:
            state = row
        else:
            return action_list

def is_goal(state, players):
    for player in players:
        if state["Occupying_player"] == player:
            return True
    return False

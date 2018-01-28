package main


func IndexOf(array []int, v int) int {
  for index, value := range array {
    if value == v {
      return index
    }
  }
  return -1
}

func ConstructPath(meta []int, state int) []int {
  var path []int

  for true {
    path = append(path, state)
    parent := meta[state]
    if parent != -1 {
      state = parent
    } else {
      for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
        path[i], path[j] = path[j], path[i]
      }
      return path
    }
  }
  return []int{}
}

func BFS(start int, graph []Base, target_players []int) []int {
  // The array of nodes showing what the parent of
  // each node is in the BST tree with root at start.
  meta := make([]int, len(graph))
  meta[start] = -1;

  // Mark all the vertices as not visited
  visited := make([]bool, len(graph))

  // Create a queue for BFS
  var queue []int

  // Mark the current node as visited and enqueue it
  visited[start] = true
  queue = append(queue, start)

  for len(queue) > 0 {
    start, queue = queue[0], queue[1:]

    connections := graph[start].Connections

    // Get all adjacent vertices of the dequeued
    // vertex s. If a adjacent has not been visited,
    // then mark it visited and enqueue it .get<int>()
    for _, value := range connections {
      if (!visited[value]) {
        meta[value]    = start;
        if IndexOf(target_players, graph[value].Occupying_player) >= 0 {
          return ConstructPath(meta, value)
        }
        visited[value] = true
        queue = append(queue, value)
      }
    }
  }

  var temp []int
  return temp
}

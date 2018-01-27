#ifndef MYSEARCHES_HPP
#define MYSEARCHES_HPP

#include <list>

int IndexOf(json array, int var) {
  for (int i = 0; i < array.size(); i++)
    if (array[i] == var)
      return i;
  return -1;
}

json ConstructPath(int* meta, int state) {
  list<int> path;

  while (true) {
    path.push_front(state);
    int parent = meta[state];
    if (parent != -1)
      state = parent;
    else {
      json j(path);
      return j;
    }

  }
}

json BFS(int start, json graph, json target_players) {
  // The array of nodes showing what the parent of
  // each node is in the BST tree with root at start.
  int *meta   = new int[graph.size()];
  meta[start] = -1;

  // Mark all the vertices as not visited
  bool *visited = new bool[graph.size()];
  for(int i = 0; i < graph.size(); i++)
    visited[i] = false;

  // Create a queue for BFS
  list<int> queue;

  // Mark the current node as visited and enqueue it
  visited[start] = true;
  queue.push_back(start);

  // 'i' will be used to get all adjacent
  // vertices of a vertex
  json::iterator i;

  while(!queue.empty()) {
    // Dequeue a vertex from queue and print it
    start = queue.front();
    queue.pop_front();

    json connections = graph[start]["Connections"];

    // Get all adjacent vertices of the dequeued
    // vertex s. If a adjacent has not been visited,
    // then mark it visited and enqueue it .get<int>()
    for (int i = 0; i < connections.size(); i++) {
      int current_node_index = connections[i];
      if (!visited[current_node_index]) {
        meta[current_node_index]    = start;
        if (IndexOf(target_players, graph[current_node_index]["Occupying_player"]) >= 0)
          return ConstructPath(meta, current_node_index);
        visited[current_node_index] = true;
        queue.push_back(current_node_index);
      }
    }
  }
  delete visited;
  delete meta;

  json temp;
  return temp;
}

#endif

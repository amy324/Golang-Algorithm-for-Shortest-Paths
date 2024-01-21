

# Dijkstra's Algorithm for Shortest Paths

This project is part of a series of short CLI-based projects inspired by classic computer science problems. I have been reading up on computer science theories and thought there is no better way to test my understanding than with mini golang CLI projects.

## Dijkstra's Algorithm

This Go program implements Dijkstra's algorithm to find the shortest paths in a weighted graph. The program provides a textual representation of the graph's adjacency matrix and outputs the shortest paths from a specified source vertex.

[Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) is a graph search algorithm that efficiently finds the shortest paths in a weighted graph. It starts from a source vertex and explores the graph to determine the shortest paths to all other vertices.

The main steps of Dijkstra's algorithm are:

1\. **Initialize distances:** Set the distance from the source vertex to itself to 0, and all other distances to infinity.\
2\. **Visit vertices:** For each vertex, update its distance if a shorter path is found.\
3\. **Mark as visited:** Mark the visited vertex as "processed" to ensure its distance remains final.\
4\. **Repeat:** Repeat the process until all vertices are processed.

## Program Overview

### 1. Graph Representation

- The `Graph` struct represents a weighted graph using a 2D adjacency matrix.\
- The matrix is initialized with maximum integer values (`math.MaxInt32`) indicating no direct edges between vertices.

### 2. Initialization

- The `Initialize` method sets up a new graph with a specified number of vertices.\
- It initializes the adjacency matrix with maximum integer values.

### 3. Adding Edges

- The `AddEdge` method adds a weighted edge between two vertices in the graph.\
- Both the source-to-destination and destination-to-source entries in the adjacency matrix are updated.

### 4. Dijkstra's Algorithm

- The `Dijkstra` method finds the shortest paths from a given source vertex using Dijkstra's algorithm.\
- It maintains an array `dist` to store the current shortest distances from the source vertex to each vertex in the graph.\
- The algorithm iteratively selects the vertex with the minimum distance and updates the distances of its neighboring vertices.

### 5. Utility Functions

- `minDistance`: Finds the vertex with the minimum distance value from the set of vertices not yet processed.\
- `PrintGraph`: Prints the adjacency matrix of the graph, displaying edge weights or "∞" for unconnected vertices.\
- `PrintShortestPaths`: Prints the shortest paths from the source vertex to all other vertices.

### 6. Main Function

- The `main` function demonstrates the usage of the implemented graph and Dijkstra's algorithm.\
- It creates a graph with 6 vertices, adds weighted edges, and prints the adjacency matrix.\
- Dijkstra's algorithm is applied to find the shortest paths from a specified source vertex to all other vertices, and the result is printed.

## Graph Representation

The program uses an adjacency matrix to represent the weighted graph. The graph is initialized with a given number of vertices, and edges are added with corresponding weights.

```go
// Example graph initialization:
g := Graph{}
g.Initialize(6)
g.AddEdge(0, 1, 2)
// Add more edges...
```
## Getting started

1.  Clone this repository to your local machine.
2.  Ensure you have GoLang installed.
3.  Navigate to the project directory in your terminal.


## Running the Program

To run the program, execute the following command:

```bash
go run .
```

The program will output the graph's adjacency matrix and the shortest paths from the specified source vertex.

**Example Output:**

```plaintext\
Graph Adjacency Matrix:
∞ 2 4 ∞ ∞ ∞
2 ∞ 1 7 ∞ ∞
4 1 ∞ ∞ 3 ∞
∞ 7 ∞ ∞ 1 5
∞ ∞ 3 1 ∞ 2
∞ ∞ ∞ 5 2 ∞

Shortest Paths from Vertex 0:
Vertex 0: 0
Vertex 1: 2
Vertex 2: 3
Vertex 3: 7
Vertex 4: 6
Vertex 5: 8
```

Feel free to explore the code in `main.go` and customize it for your specific use case.
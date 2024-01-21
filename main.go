package main

import (
	"fmt"
	"math"
)

const infinity = "∞"

// Graph represents the weighted graph using an adjacency matrix.
type Graph struct {
	vertices int
	matrix   [][]int
}

// Initialize initializes a new graph with the given number of vertices.
func (g *Graph) Initialize(vertices int) {
	g.vertices = vertices
	g.matrix = make([][]int, vertices)
	for i := range g.matrix {
		g.matrix[i] = make([]int, vertices)
		for j := range g.matrix[i] {
			g.matrix[i][j] = math.MaxInt32 // Initialize with maximum integer value
		}
	}
}

// AddEdge adds a weighted edge between two vertices in the graph.
func (g *Graph) AddEdge(src, dest, weight int) {
	g.matrix[src][dest] = weight
	g.matrix[dest][src] = weight
}

// Dijkstra implements Dijkstra's algorithm to find the shortest paths from a source vertex.
func (g *Graph) Dijkstra(source int) []int {
	dist := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	// Initialize distances with infinity and set the source vertex distance to 0.
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	for count := 0; count < g.vertices-1; count++ {
		// Find the minimum distance vertex from the set of vertices not yet processed.
		u := minDistance(dist, visited)
		visited[u] = true

		// Update the distance value of the neighboring vertices.
		for v := 0; v < g.vertices; v++ {
			if !visited[v] && g.matrix[u][v] != math.MaxInt32 && dist[u] != math.MaxInt32 &&
				dist[u]+g.matrix[u][v] < dist[v] {
				dist[v] = dist[u] + g.matrix[u][v]
			}
		}
	}

	return dist
}

// minDistance finds the vertex with the minimum distance value from the set of vertices not yet processed.
func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := -1

	for v := range dist {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

// PrintGraph prints the adjacency matrix of the graph.
func PrintGraph(g *Graph) {
	fmt.Println("Graph Adjacency Matrix:")
	for _, row := range g.matrix {
		for _, val := range row {
			if val == math.MaxInt32 {
				fmt.Print(infinity, " ")
			} else {
				fmt.Print(val, " ")
			}
		}
		fmt.Println()
	}
}

// PrintShortestPaths prints the shortest paths from the source vertex.
func PrintShortestPaths(source int, shortestPaths []int) {
	fmt.Printf("\nShortest Paths from Vertex %d:\n", source)
	for i, distance := range shortestPaths {
		fmt.Printf("Vertex %d: %d\n", i, distance)
	}
}

func main() {
	g := Graph{}
	g.Initialize(6)

	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 7)
	g.AddEdge(2, 4, 3)
	g.AddEdge(3, 4, 1)
	g.AddEdge(3, 5, 5)
	g.AddEdge(4, 5, 2)

	source := 0

	PrintGraph(&g)

	shortestPaths := g.Dijkstra(source)
	PrintShortestPaths(source, shortestPaths)
}


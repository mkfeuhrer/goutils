package graph

// Graph represents an adjacency list graph
type Graph struct {
	vertices map[int][]int
	directed bool // Indicates if the graph is directed
}

// NewGraph creates a new Graph
func NewGraph(directed bool) *Graph {
	return &Graph{vertices: make(map[int][]int), directed: directed}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
	if !g.directed {
		g.vertices[w] = append(g.vertices[w], v) // For undirected graph, add both edges
	}
}

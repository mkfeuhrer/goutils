package graph

// BFS performs Breadth-First Search starting from the source vertex
func (g *Graph) BFS(start int) []int {
	// Handle empty graph case
	if len(g.vertices) == 0 {
		return []int{}
	}

	// Create a queue for BFS
	queue := []int{start}

	// Create a map to keep track of visited vertices
	visited := make(map[int]bool)
	visited[start] = true

	// Slice to store the order of traversal
	traversalOrder := []int{}

	for len(queue) > 0 {
		// Dequeue a vertex from queue
		vertex := queue[0]
		queue = queue[1:]

		// Add this vertex to the traversal order
		traversalOrder = append(traversalOrder, vertex)

		// Get all adjacent vertices of the dequeued vertex
		for _, adjacent := range g.vertices[vertex] {
			// If an adjacent has not been visited, mark it visited and enqueue it
			if !visited[adjacent] {
				visited[adjacent] = true
				queue = append(queue, adjacent)
			}
		}
	}

	return traversalOrder
}

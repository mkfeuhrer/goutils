package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	t.Run("Basic graph", func(t *testing.T) {
		g1 := NewGraph(false)
		g1.AddEdge(0, 1)
		g1.AddEdge(0, 2)
		g1.AddEdge(1, 2)
		g1.AddEdge(2, 3)
		g1.AddEdge(3, 4)

		result1 := g1.BFS(0)
		expected1 := []int{0, 1, 2, 3, 4}
		if len(result1) != len(expected1) {
			t.Errorf("BFS traversal order length is incorrect. Got: %v, Expected: %v", result1, expected1)
			return
		}
		for i, v := range expected1 {
			if result1[i] != v {
				t.Errorf("BFS traversal order is incorrect. Got: %v, Expected: %v", result1, expected1)
				return
			}
		}
	})

	t.Run("Empty graph", func(t *testing.T) {
		g2 := NewGraph(false)
		result2 := g2.BFS(0)
		expected2 := []int{}
		if len(result2) != len(expected2) {
			t.Errorf("BFS traversal order length is incorrect for empty graph. Got: %v, Expected: %v", result2, expected2)
			return
		}

	})

	t.Run("Single vertex", func(t *testing.T) {
		g3 := NewGraph(false)
		g3.AddEdge(0, 0)
		result3 := g3.BFS(0)
		expected3 := []int{0}
		if len(result3) != len(expected3) {
			t.Errorf("BFS traversal order length is incorrect for single vertex. Got: %v, Expected: %v", result3, expected3)
			return
		}
		for i, v := range expected3 {
			if result3[i] != v {
				t.Errorf("BFS traversal order is incorrect for single vertex. Got: %v, Expected: %v", result3, expected3)
				return
			}
		}
	})

	t.Run("Disconnected graph", func(t *testing.T) {
		g4 := NewGraph(false)
		g4.AddEdge(0, 1)
		g4.AddEdge(2, 3)
		result4 := g4.BFS(0)
		expected4 := []int{0, 1}
		if len(result4) != len(expected4) {
			t.Errorf("BFS traversal order length is incorrect for disconnected graph. Got: %v, Expected: %v", result4, expected4)
			return
		}
		for i, v := range expected4 {
			if result4[i] != v {
				t.Errorf("BFS traversal order is incorrect for disconnected graph. Got: %v, Expected: %v", result4, expected4)
				return
			}
		}
	})
}

func TestAddEdgeAndBFS(t *testing.T) {
	// Test for undirected graph
	t.Run("undirected graph edge case", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		result5 := g.BFS(0)
		expected := []int{0, 1, 2}
		if len(result5) != len(expected) {
			t.Errorf("BFS traversal order length is incorrect after adding edges. Got: %v, Expected: %v", result5, expected)
			return
		}
		for i, v := range expected {
			if result5[i] != v {
				t.Errorf("BFS traversal order is incorrect after adding edges. Got: %v, Expected: %v", result5, expected)
				return
			}
		}
	})

	// Test for directed graph
	t.Run("directed graph add edge", func(t *testing.T) {
		gDirected := NewGraph(true)
		gDirected.AddEdge(0, 1)
		gDirected.AddEdge(1, 2)
		resultDirected := gDirected.BFS(0)
		expectedDirected := []int{0, 1, 2}
		if len(resultDirected) != len(expectedDirected) {
			t.Errorf("BFS traversal order length is incorrect for directed graph. Got: %v, Expected: %v", resultDirected, expectedDirected)
			return
		}
		for i, v := range expectedDirected {
			if resultDirected[i] != v {
				t.Errorf("BFS traversal order is incorrect for directed graph. Got: %v, Expected: %v", resultDirected, expectedDirected)
				return
			}
		}
	})
}

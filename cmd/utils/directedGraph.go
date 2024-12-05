package utils

import (
	"fmt"
)

type DirectedGraph struct {
	matrix [][]int
}

func CreateGraph(size int) *DirectedGraph {
	matrix := make([][]int, size)

	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	return &DirectedGraph{
		matrix: matrix,
	}
}

func (dg *DirectedGraph) AddEdge(A int, B int) {
	dg.matrix[A][B] = 1
}

func (dg *DirectedGraph) RemoveEdge(A int, B int) {
	dg.matrix[A][B] = 0
}

func (dg *DirectedGraph) RemoveAllEdges() {
	for i := 0; i < len(dg.matrix); i++ {
		for j := 0; j < len(dg.matrix); j++ {
			dg.matrix[i][j] = 0
		}
	}
}

func (dg *DirectedGraph) HasEdge(A int, B int) bool {
	return dg.matrix[A][B] == 1
}

func (dg *DirectedGraph) String() string {
	result := ""

	for _, val := range dg.matrix {
		result += fmt.Sprintln(val)
	}

	return result
}

func (dg *DirectedGraph) GetMatrix() [][]int {
	return dg.matrix
}

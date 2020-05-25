package graph

import "github.com/illfate2/graph-api/pkg/model"

type (
	IncidenceMatrix map[model.Edge]map[model.Node]int
	AdjacencyMatrix map[model.Node]map[model.Node]int
)

type Methods interface {
	IncidenceMatrix(graph model.Graph) IncidenceMatrix
	AdjacencyMatrix(graph model.Graph) AdjacencyMatrix
}

type Graph struct {
}

func (g Graph) IncidenceMatrix(graph model.Graph) IncidenceMatrix {
	nodes := make(map[model.Node]struct{})
	for _, e := range graph.Edges {
		nodes[e.From] = struct{}{}
		nodes[e.To] = struct{}{}
	}

	edges := make(IncidenceMatrix)
	for _, e := range graph.Edges {
		mNodes := make(map[model.Node]int)
		for n := range nodes {
			mNodes[n] = 0
			if e.From == n {
				mNodes[n] = 1
			}
			if e.To == n {
				mNodes[n] = 1
			}
		}
		edges[e] = mNodes
	}
	return edges
}

func (g Graph) AdjacencyMatrix(graph model.Graph) AdjacencyMatrix {
	nodes := make(map[model.Node][]model.Node)
	for _, e := range graph.Edges {
		nodes[e.From] = []model.Node{}
		nodes[e.To] = []model.Node{}
	}
	for _, e := range graph.Edges {
		nodes[e.From] = append(nodes[e.From], e.To)
		nodes[e.To] = append(nodes[e.To], e.From)
	}
	matrix := make(AdjacencyMatrix)
	for n, s := range nodes {
		nodeCount := make(map[model.Node]int)
		for n2 := range nodes {
			nodeCount[n2] = 0
			for _, nInS := range s {
				if nInS == n2 {
					nodeCount[n2] = 1
					break
				}
			}
		}
		matrix[n] = nodeCount
	}
	return matrix
}

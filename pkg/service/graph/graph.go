package graph

import (
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"

	"github.com/illfate2/graph-api/pkg/model"
)

type (
	IncidenceMatrix map[model.Edge]map[model.Node]int
	AdjacencyMatrix map[model.Node]map[model.Node]int
)

type Methods interface {
	IncidenceMatrix(graph model.Graph) IncidenceMatrix
	AdjacencyMatrix(graph model.Graph) AdjacencyMatrix
	ShortestPath(graph model.Graph, fromNode, toNode uint64) []model.Node
	AllShortestPaths(graph model.Graph, fromNode, toNode uint64) [][]model.Node
	HamiltonianPath(graph model.Graph, orig uint64) ([]model.Node, bool)
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

func (g Graph) ShortestPath(graph model.Graph, fromNode, toNode uint64) []model.Node {
	shortest := path.DijkstraAllPaths(g.toUndirectedGraph(graph))
	p, _, _ := shortest.Between(int64(fromNode), int64(toNode))
	resPath := make([]model.Node, 0, len(p))

	nodes := graphToNodes(graph)
	for _, n := range p {
		resNode := nodes[uint64(n.ID())]
		resPath = append(resPath, resNode)
	}
	return resPath
}

func (g Graph) AllShortestPaths(graph model.Graph, fromNode, toNode uint64) [][]model.Node {
	shortest := path.DijkstraAllPaths(g.toUndirectedGraph(graph))
	p, _ := shortest.AllBetween(int64(fromNode), int64(toNode))
	resPaths := make([][]model.Node, 0, len(p))

	nodes := graphToNodes(graph)
	for i := range p {
		resPath := make([]model.Node, 0)
		for _, n := range p[i] {
			resNode := nodes[uint64(n.ID())]
			resPath = append(resPath, resNode)
		}
		resPaths = append(resPaths, resPath)
	}
	return resPaths
}

func (g Graph) HamiltonianPath(graph model.Graph, orig uint64) ([]model.Node, bool) {
	visited := make(map[uint64]bool)
	path := []uint64{orig}
	nodeToEdges := make(map[uint64]map[uint64]struct{})
	for _, e := range graph.Edges {
		nodeToEdges[e.From.ID] = make(map[uint64]struct{})
		nodeToEdges[e.To.ID] = make(map[uint64]struct{})
	}
	for _, e := range graph.Edges {
		nodeToEdges[e.From.ID][e.To.ID] = struct{}{}
		nodeToEdges[e.To.ID][e.From.ID] = struct{}{}
	}

	hamiltonPath, find := g.hamiltonianPath(orig, orig, visited, path, nodeToEdges)
	if !find {
		return nil, false
	}
	nodes := graphToNodes(graph)
	resPath := make([]model.Node, 0, len(hamiltonPath))
	for _, n := range path {
		resNode := nodes[n]
		resPath = append(resPath, resNode)
	}
	return resPath, true
}

func (g Graph) hamiltonianPath(
	orig, dest uint64,
	visited map[uint64]bool,
	path []uint64,
	nodeToEdges map[uint64]map[uint64]struct{},
) ([]uint64, bool) {
	if len(visited) == len(nodeToEdges) {
		if path[len(path)-1] == dest {
			return path, true
		}

		return nil, false
	}

	for tv := range nodeToEdges[orig] {
		if _, ok := visited[tv]; !ok && (dest != tv || len(visited) == len(nodeToEdges)-1) {
			visited[tv] = true
			path = append(path, tv)
			if path, found := g.hamiltonianPath(tv, dest, visited, path, nodeToEdges); found {
				return path, true
			}
			path = path[:len(path)-1]
			delete(visited, tv)
		}
	}

	return nil, false
}

func (g Graph) toUndirectedGraph(graph model.Graph) *simple.UndirectedGraph {
	undirGraph := simple.NewUndirectedGraph()
	for _, e := range graph.Edges {
		undirGraph.Edge(int64(e.From.ID), int64(e.To.ID))
	}
	return undirGraph
}

func graphToNodes(graph model.Graph) map[uint64]model.Node {
	nodes := make(map[uint64]model.Node)
	for _, e := range graph.Edges {
		nodes[e.From.ID] = e.From
		nodes[e.To.ID] = e.To
	}
	return nodes
}

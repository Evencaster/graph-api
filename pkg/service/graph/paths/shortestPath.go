package paths

import (
	"github.com/illfate2/graph-api/pkg/model"
)

var allPaths [][]model.Node
var shortestCost uint64

func rec(nodes map[model.Node]struct{}, matrix map[model.Node]map[model.Node]int, currentNode model.Node, toNode, cost uint64, visitedEdges map[model.Node]uint64, path []model.Node) {

	if currentNode.ID == toNode {

		if shortestCost >= cost {
			allPaths = append(allPaths, path)
			shortestCost = cost

		}
		return
	}

	cost++

	for nextNode := range nodes {

		if matrix[currentNode][nextNode] == 0 {
			continue
		}

		if visitedEdges[nextNode] >= cost || visitedEdges[nextNode] == 0 {

			visitedEdges[nextNode] = cost
			path := append(path, nextNode)
			rec(nodes, matrix, nextNode, toNode, cost, visitedEdges, path)

		}
	}
}

func AllShortestPathsFind(nodes map[model.Node]struct{}, matrix map[model.Node]map[model.Node]int, currentNode model.Node, toNode uint64) [][]model.Node {

	allPaths = nil
	visitedEdges := make(map[model.Node]uint64)

	cost := uint64(0)
	shortestCost = uint64(len(nodes) + 2)

	var path []model.Node
	var shortestPaths [][]model.Node

	path = append(path, currentNode)
	visitedEdges[currentNode] = uint64(len(nodes)) + 2

	rec(nodes, matrix, currentNode, toNode, cost, visitedEdges, path)

	if allPaths == nil {
		return nil
	}

	for _, path := range allPaths {
		if len(path) == len(allPaths[len(allPaths)-1]) {
			shortestPaths = append(shortestPaths, path)
		}
	}

	return shortestPaths
}

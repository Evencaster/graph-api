package service

import (
	"github.com/illfate2/graph-api/pkg/model"
	"github.com/illfate2/graph-api/pkg/repository"
	"github.com/illfate2/graph-api/pkg/service/graph"
)

type Service interface {
	repository.Repository
	IncidenceMatrix(id uint64) (graph.IncidenceMatrix, error)
	AdjacencyMatrix(id uint64) (graph.AdjacencyMatrix, error)
	PlanarCheck(id uint64) (bool, error)
	FindDiameter(id uint64) (uint64, error)
	FindRadius(id uint64) (uint64, error)
	FindCenter(id uint64) ([]model.Node, error)
	ShortestPath(graphID, fromNode, toNode uint64) ([]model.Node, error)
	AllShortestPaths(graphID, fromNode, toNode uint64) ([][]model.Node, error)
	HamiltonianPath(graphID, startedNode uint64) ([]model.Node, error)
	EulerianCycle(graphID, startedNode uint64) ([]model.Node, error)
}

type Graph struct {
	repository repository.Repository
	graph      graph.Methods
}

func NewGraph(repo repository.Repository) *Graph {
	return &Graph{
		repository: repo,
		graph:      graph.Graph{},
	}
}

func (g *Graph) PlanarCheck(id uint64) (bool, error) {
	foundGraph, err := g.Graph(id)
	if err != nil {
		return false, err
	}
	return g.graph.PlanarCheck(foundGraph), nil
}

func (g *Graph) FindDiameter(id uint64) (uint64, error) {
	foundGraph, err := g.Graph(id)
	if err != nil {
		return 0, err
	}
	return g.graph.FindDiameter(foundGraph), nil
}

func (g *Graph) FindCenter(id uint64) ([]model.Node, error) {
	foundGraph, err := g.Graph(id)
	if err != nil {
		return nil, err
	}
	return g.graph.FindCenter(foundGraph), nil
}

func (g *Graph) FindRadius(id uint64) (uint64, error) {
	foundGraph, err := g.Graph(id)
	if err != nil {
		return 0, err
	}
	return g.graph.FindRadius(foundGraph), nil
}

func (g *Graph) IncidenceMatrix(id uint64) (graph.IncidenceMatrix, error) {
	foundGraph, err := g.Graph(id)
	if err != nil {
		return nil, err
	}
	return g.graph.IncidenceMatrix(foundGraph), nil
}

func (g *Graph) AdjacencyMatrix(id uint64) (graph.AdjacencyMatrix, error) {
	foundGraph, err := g.Graph(id)
	if err != nil {
		return nil, err
	}
	return g.graph.AdjacencyMatrix(foundGraph), nil
}

func (g *Graph) ShortestPath(graphID, fromNode, toNode uint64) ([]model.Node, error) {
	foundGraph, err := g.Graph(graphID)
	if err != nil {
		return nil, err
	}
	return g.graph.ShortestPath(foundGraph, fromNode, toNode), nil
}

func (g *Graph) AllShortestPaths(graphID, fromNode, toNode uint64) ([][]model.Node, error) {
	foundGraph, err := g.Graph(graphID)
	if err != nil {
		return nil, err
	}
	return g.graph.AllShortestPaths(foundGraph, fromNode, toNode), nil
}

func (g *Graph) HamiltonianPath(graphID, startedNode uint64) ([]model.Node, error) {
	return g.path(graphID, startedNode, g.graph.HamiltonianPath)
}

func (g *Graph) EulerianCycle(graphID, startedNode uint64) ([]model.Node, error) {
	return g.path(graphID, startedNode, g.graph.EulerianCycle)
}

type findPathF func(graph model.Graph, startedNode uint64) ([]model.Node, bool)

func (g *Graph) path(graphID, startedNode uint64, f findPathF) ([]model.Node, error) {
	foundGraph, err := g.Graph(graphID)
	if err != nil {
		return nil, err
	}
	path, found := f(foundGraph, startedNode)
	if !found {
		return nil, repository.ErrNotFound
	}
	return path, nil
}

func (g *Graph) List() ([]model.Graph, error) {
	return g.repository.List()
}

func (g *Graph) CreateGraph(graph model.Graph) (uint64, error) {
	return g.repository.CreateGraph(graph)
}

func (g *Graph) Graph(id uint64) (model.Graph, error) {
	return g.repository.Graph(id)
}

func (g *Graph) UpdateGraph(graph model.Graph) error {
	return g.repository.UpdateGraph(graph)
}

func (g *Graph) DeleteGraph(id uint64) error {
	return g.repository.DeleteGraph(id)
}

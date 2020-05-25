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

package service

import (
	"github.com/illfate2/graph-api/pkg/model"
)

type Service interface {
	CreateGraph(graph model.Graph) (uint64, error)
	Graph(id uint64) (model.Graph, error)
	UpdateGraph(graph model.Graph) error
	DeleteGraph(id uint64) error
}

type Graph struct {
	repository Service
}

func NewGraph(repo Service) *Graph {
	return &Graph{repository: repo}
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

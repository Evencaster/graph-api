package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	_ "github.com/google/uuid"

	"github.com/illfate2/graph-api/pkg/model"
)

var ErrNotFound = errors.New("not found")

type DB struct {
	data sync.Map
	mx   sync.Mutex
	uuid uuid.UUID
}

func New() *DB {
	return &DB{
		data: sync.Map{},
		uuid: uuid.New(),
		mx:   sync.Mutex{},
	}
}

func (d *DB) CreateGraph(graph model.Graph) (uint64, error) {
	d.mx.Lock()
	id := uint64(d.uuid.ID())
	d.mx.Unlock()
	graph.ID = id
	d.data.Store(id, graph)
	return id, nil
}

func (d *DB) Graph(id uint64) (model.Graph, error) {
	graph, find := d.data.Load(id)
	if !find {
		return model.Graph{}, ErrNotFound
	}
	g, ok := graph.(model.Graph)
	if !ok {
		return model.Graph{}, errors.New("can't cast")
	}
	return g, nil
}

func (d *DB) UpdateGraph(graph model.Graph) error {
	d.data.Store(graph.ID, graph)
	return nil
}

func (d *DB) DeleteGraph(id uint64) error {
	d.data.Delete(id)
	return nil
}

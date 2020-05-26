package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/illfate2/graph-api/pkg/model"
)

func TestGraph_ShortestPath(t *testing.T) {
	type args struct {
		graph    model.Graph
		fromNode uint64
		toNode   uint64
	}
	tests := []struct {
		name string
		args args
		want []model.Node
	}{
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								Name:  "First",
								Color: "BLue",
								ID:    1,
							},
							To: model.Node{
								ID: 2,
							},
						},
						{
							From: model.Node{
								ID: 2,
							},
							To: model.Node{
								ID: 3,
							},
						},
						{
							From: model.Node{
								Name:  "First",
								Color: "BLue",
								ID:    1,
							},
							To: model.Node{
								ID: 3,
							},
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								ID: 5,
							},
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 3,
							},
						},
					},
				},
				fromNode: 1,
				toNode:   4,
			},
			want: []model.Node{
				{
					Name:  "First",
					Color: "BLue",
					ID:    1,
				},
				{
					ID: 3,
				},
				{
					ID: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Graph{}
			got := g.ShortestPath(tt.args.graph, tt.args.fromNode, tt.args.toNode)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGraph_HamiltonianPath(t *testing.T) {
	type args struct {
		graph model.Graph
		orig  uint64
	}
	tests := []struct {
		name  string
		args  args
		want  []model.Node
		want1 bool
	}{
		{
			name: "",
			args: args{
				graph: model.Graph{
					ID:   0,
					Name: "",
					Edges: []model.Edge{
						{
							From: model.Node{
								Name:  "First",
								Color: "BLue",
								ID:    1,
							},
							To: model.Node{
								ID: 2,
							},
						},
						{
							From: model.Node{
								ID: 2,
							},
							To: model.Node{
								ID: 3,
							},
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								Name:  "First",
								Color: "BLue",
								ID:    1,
							},
						},
					},
				},
				orig: 1,
			},
			want: []model.Node{
				{
					Name:  "First",
					Color: "BLue",
					ID:    1,
				},
				{
					ID: 2,
				},
				{
					ID: 3,
				},
				{
					Name:  "First",
					Color: "BLue",
					ID:    1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Graph{}
			got, _ := g.HamiltonianPath(tt.args.graph, tt.args.orig)
			assert.Equal(t, tt.want, got)
		})
	}
}

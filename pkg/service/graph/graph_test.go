package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/illfate2/graph-api/pkg/model"
)


func TestGraph_PlanarCheck(t *testing.T) {
	type args struct{
		graph model.Graph
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								ID: 1,
							},
							To: model.Node{
								ID: 2,
							},
							IsDirected: true,
							Name:       "1",
						},
						{
							From: model.Node{
								ID: 2,
							},
							To: model.Node{
								ID: 3,
							},
							IsDirected: true,
							Name:       "2",
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								ID: 4,
							},
							IsDirected: true,
							Name:       "3",
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 1,
							},
							IsDirected: true,
							Name:       "4",
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 5,
							},
							IsDirected: true,
							Name:       "5",
						},
						{
							From: model.Node{
								ID: 5,
							},
							To: model.Node{
								ID: 3,
							},
							IsDirected: true,
							Name:       "6",
						},
					},
				},
			},
			want: true,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			g := Graph{}
			got := g.PlanarCheck(test.args.graph)
			assert.Equal(t, test.want, got)
		})
	}

}

func TestGraph_IncidenceMatrix(t *testing.T) {
	type args struct{
		graph model.Graph
	}
	answerOriented := make(map[model.Edge]map[model.Node]int)
	answerUnOriented := make(map[model.Edge]map[model.Node]int)

	matrixOriented := [6][5]int{
		{1, -1, 0, 0, 0},
		{0, 1, -1, 0, 0},
		{0, 0, 1, -1, 0},
		{-1, 0, 0, 1, 0},
		{0, 0, 0, 1, -1},
		{0, 0, -1, 0, 1},
	}

	matrixUnOriented := [6][5]int{
		{1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1},
	}

	edgeList := []model.Edge{
		{
		Name: "1",
		From: model.Node{ID: 1},
		To:   model.Node{ID: 2},
		},
		{
			Name: "2",
			From: model.Node{ID: 2},
			To:   model.Node{ID: 3},
		},
		{
			Name: "3",
			From: model.Node{ID: 3},
			To:   model.Node{ID: 4},
		},
		{
			Name: "4",
			From: model.Node{ID: 4},
			To:   model.Node{ID: 1},
		},
		{
			Name: "5",
			From: model.Node{ID: 4},
			To:   model.Node{ID: 5},
		},
		{
			Name: "6",
			From: model.Node{ID: 5},
			To:   model.Node{ID: 3},
		},
	}

	for i:=0; i<6; i++{
		sub := make(map[model.Node]int)
		for j := 0; j<5; j++{
			sub[model.Node{ID: uint64(j+1)}] = matrixOriented[i][j]
		}
		answerOriented[model.Edge{Name: edgeList[i].Name, From: edgeList[i].From, To: edgeList[i].To, IsDirected: true}] = sub
	}

	for i:=0; i<6; i++{
		sub := make(map[model.Node]int)
		for j := 0; j<5; j++{
			sub[model.Node{ID: uint64(j+1)}] = matrixUnOriented[i][j]
		}
		answerUnOriented[edgeList[i]] = sub
	}

	tests := []struct {
		name string
		args args
		want IncidenceMatrix
	}{
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								ID: 1,
							},
							To: model.Node{
								ID: 2,
							},
							IsDirected: true,
							Name: "1",
						},
						{
							From: model.Node{
								ID: 2,
							},
							To: model.Node{
								ID: 3,
							},
							IsDirected: true,
							Name: "2",
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								ID: 4,
							},
							IsDirected: true,
							Name: "3",
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 1,
							},
							IsDirected: true,
							Name: "4",
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 5,
							},
							IsDirected: true,
							Name: "5",
						},
						{
							From: model.Node{
								ID: 5,
							},
							To: model.Node{
								ID: 3,
							},
							IsDirected: true,
							Name: "6",
						},
					},
				},
			},
			want: answerOriented,
		},
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								ID: 1,
							},
							To: model.Node{
								ID: 2,
							},
							Name: "1",
						},
						{
							From: model.Node{
								ID: 2,
							},
							To: model.Node{
								ID: 3,
							},
							Name: "2",
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								ID: 4,
							},
							Name: "3",
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 1,
							},
							Name: "4",
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 5,
							},
							Name: "5",
						},
						{
							From: model.Node{
								ID: 5,
							},
							To: model.Node{
								ID: 3,
							},
							Name: "6",
						},
					},
				},
			},
			want: answerUnOriented,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			g := Graph{}
			got := g.IncidenceMatrix(test.args.graph)
			assert.Equal(t, test.want, got)
		})
	}

}

func TestGraph_AdjacencyMatrix(t *testing.T) {
	type args struct{
		graph model.Graph
	}
	answerOriented := make(map[model.Node]map[model.Node]int)
	answerUnOriented := make(map[model.Node]map[model.Node]int)

	matrixOriented := [5][5]int{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 0, 1, 0, 0},
	}

	matrixUnOriented := [5][5]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 0, 1, 1, 0},
	}

	for i:=0; i<5; i++{
		sub := make(map[model.Node]int)
		for j := 0; j<5; j++{
			sub[model.Node{ID: uint64(j+1)}] = matrixOriented[i][j]
		}
		answerOriented[model.Node{ID: uint64(i+1)}] = sub
	}

	for i:=0; i<5; i++{
		sub := make(map[model.Node]int)
		for j := 0; j<5; j++{
			sub[model.Node{ID: uint64(j+1)}] = matrixUnOriented[i][j]
		}
		answerUnOriented[model.Node{ID: uint64(i+1)}] = sub
	}

	tests := []struct {
		name string
		args args
		want AdjacencyMatrix
	}{
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								ID: 1,
							},
							To: model.Node{
								ID: 2,
							},
							IsDirected: true,
						},
						{
							From: model.Node{
								ID: 2,
							},
							To: model.Node{
								ID: 3,
							},
							IsDirected: true,
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								ID: 4,
							},
							IsDirected: true,
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 1,
							},
							IsDirected: true,
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 5,
							},
							IsDirected: true,
						},
						{
							From: model.Node{
								ID: 5,
							},
							To: model.Node{
								ID: 3,
							},
							IsDirected: true,
						},
					},
				},
			},
			want: answerOriented,
		},
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								ID: 1,
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
								ID: 4,
							},
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 1,
							},
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 5,
							},
						},
						{
							From: model.Node{
								ID: 5,
							},
							To: model.Node{
								ID: 3,
							},
						},
					},
				},
			},
			want: answerUnOriented,
		},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			g := Graph{}
			got := g.AdjacencyMatrix(test.args.graph)
			assert.Equal(t, test.want, got)
		})
	}
}

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

func TestGraph_AllShortestPaths(t *testing.T) {
	type args struct {
		graph    model.Graph
		fromNode uint64
		toNode   uint64
	}
	tests := []struct {
		name string
		args args
		want [][]model.Node
	}{
		{
			args: args{
				graph: model.Graph{
					Edges: []model.Edge{
						{
							From: model.Node{
								Name:  "First",
								Color: "BLue",
								ID:    2,
							},
							To: model.Node{
								ID: 1,
							},
						},
						{
							From: model.Node{
								ID: 3,
							},
							To: model.Node{
								ID: 2,
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
						{
							From: model.Node{
								ID: 5,
							},
							To: model.Node{
								ID: 1,
								Name:  "First",
								Color: "BLue",
							},
						},
						{
							From: model.Node{
								ID: 4,
							},
							To: model.Node{
								ID: 5,
							},
						},
					},
				},
				fromNode: 1,
				toNode:   4,
			},
			want: [][]model.Node{
				{
					{
						Name:  "First",
						Color: "BLue",
						ID:    1,
					},
					{
						ID: 5,
					},
					{
						ID: 4,
					},
				},
				{
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
		},
		//{
		//	args: args{
		//		graph: model.Graph{
		//			Edges: []model.Edge{
		//				{
		//					From: model.Node{
		//						Name:  "First",
		//						Color: "BLue",
		//						ID:    2,
		//					},
		//					To: model.Node{
		//						ID: 1,
		//					},
		//					IsDirected: true,
		//				},
		//				{
		//					From: model.Node{
		//						ID: 3,
		//					},
		//					To: model.Node{
		//						ID: 2,
		//					},
		//					IsDirected: true,
		//				},
		//				{
		//					From: model.Node{
		//						Name:  "First",
		//						Color: "BLue",
		//						ID:    1,
		//					},
		//					To: model.Node{
		//						ID: 3,
		//					},
		//					IsDirected: true,
		//				},
		//				{
		//					From: model.Node{
		//						ID: 3,
		//					},
		//					To: model.Node{
		//						ID: 5,
		//					},
		//					IsDirected: true,
		//				},
		//				{
		//					From: model.Node{
		//						ID: 4,
		//					},
		//					To: model.Node{
		//						ID: 3,
		//					},
		//					IsDirected: true,
		//				},
		//				{
		//					From: model.Node{
		//						ID: 5,
		//					},
		//					To: model.Node{
		//						ID: 1,
		//						Name:  "First",
		//						Color: "BLue",
		//					},
		//					IsDirected: true,
		//				},
		//				{
		//					From: model.Node{
		//						ID: 4,
		//					},
		//					To: model.Node{
		//						ID: 5,
		//					},
		//					IsDirected: true,
		//				},
		//			},
		//		},
		//		fromNode: 4,
		//		toNode:   1,
		//	},
		//	want: [][]model.Node{
		//		{
		//			{
		//				ID:    4,
		//			},
		//			{
		//				ID: 5,
		//			},
		//			{
		//				ID: 1,
		//				Name:  "First",
		//				Color: "BLue",
		//			},
		//		},
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Graph{}
			got := g.AllShortestPaths(tt.args.graph, tt.args.fromNode, tt.args.toNode)
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
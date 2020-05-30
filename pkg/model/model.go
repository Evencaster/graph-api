package model

type Graph struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Edges []Edge `json:"arcs"`
	Nodes []Node `json:"vertexes"`
}

type NodeShape string

type Node struct {
	ID    uint64    `json:"id"`
	X 	  uint64 	`json:"x"`
	Y 	  uint64   	`json:"y"`
	Name  string    `json:"name"`
	Shape NodeShape `json:"shape"`
	Color string    `json:"color"`
}

type Edge struct {
	ID 		   uint64 `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	From       Node   `json:"vertex1"`
	To         Node   `json:"vertex2"`
	Angle12    Angle  `json:"angle12"`
	Angle21    Angle  `json:"angle21"`
	IsDirected bool   `json:"isDirected"`
}

type Angle struct {
	Sin	float64 `json:"sin"`
	Cos float64 `json:"cos"`
}

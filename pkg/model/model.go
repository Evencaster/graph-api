package model

type Graph struct {
	ID uint64
	Name  string `json:"name"`
	Edges []Edge `json:"edges"`
}

type NodeShape string

type Node struct {
	Name  string    `json:"name"`
	Shape NodeShape `json:"shape"`
	Color string    `json:"color"`
}

type Edge struct {
	Name       string `json:"name"`
	Color      string `json:"color"`
	From       Node   `json:"from"`
	To         Node   `json:"to"`
	IsDirected bool   `json:"isDirected"`
}

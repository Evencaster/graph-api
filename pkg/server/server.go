package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/illfate2/graph-api/pkg/model"
	"github.com/illfate2/graph-api/pkg/service"
	"github.com/illfate2/graph-api/pkg/service/graph"
)

type Server struct {
	http.Handler
	service service.Service
}

func New(service service.Service) *Server {
	r := mux.NewRouter()
	s := Server{
		service: service,
		Handler: r,
	}
	r.Use(CORS)
	r.HandleFunc("/api/v1/graph", s.CreateGraph).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}", s.Graph).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}", s.UpdateGraph).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}", s.DeleteGraph).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}/adjacencyMatrix", s.AdjacencyMatrix).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}/incidenceMatrix", s.IncidenceMatrix).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}/shortestPath", s.ShortestPath).
		Queries("fromNode", "{fromNode}", "toNode", "{toNode}").Methods(http.MethodGet)

	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}/allShortestPath", s.AllShortestPaths).
		Queries("fromNode", "{fromNode}", "toNode", "{toNode}").Methods(http.MethodGet)

	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}/hamiltonianPath", s.HamiltonianPath).
		Queries("startNode", "{startNode}").Methods(http.MethodGet)

	r.HandleFunc("/api/v1/graph/{id:[1-9]+[0-9]*}/planarCheck", s.PlanarCheck).Methods(http.MethodGet)
	return &s
}

func (s *Server) CreateGraph(w http.ResponseWriter, req *http.Request) {
	var g model.Graph
	err := json.NewDecoder(req.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error when decoding JSON: ", err)
		return
	}
	id, err := s.service.CreateGraph(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error when creating graph: ", err)
		return
	}
	resp := struct {
		ID uint64 `json:"id"`
	}{
		ID: id,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) Graph(w http.ResponseWriter, req *http.Request) {
	id, err := getID(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g, err := s.service.Graph(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(g)
}

func (s *Server) UpdateGraph(w http.ResponseWriter, req *http.Request) {
	id, err := getID(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var g model.Graph
	err = json.NewDecoder(req.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	g.ID = id

	err = s.service.UpdateGraph(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteGraph(w http.ResponseWriter, req *http.Request) {
	id, err := getID(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.service.DeleteGraph(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) IncidenceMatrix(w http.ResponseWriter, req *http.Request) {
	id, err := getID(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m, err := s.service.IncidenceMatrix(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := struct {
		Matrix graph.IncidenceMatrix `json:"matrix"`
	}{
		Matrix: m,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) AdjacencyMatrix(w http.ResponseWriter, req *http.Request) {
	id, err := getID(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m, err := s.service.AdjacencyMatrix(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := struct {
		Matrix graph.AdjacencyMatrix `json:"matrix"`
	}{
		Matrix: m,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) ShortestPath(w http.ResponseWriter, req *http.Request) {
	args, err := getShortestPathArgs(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path, err := s.service.ShortestPath(args.graphID, args.fromNode, args.toNode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := struct {
		Path []model.Node `json:"path"`
	}{
		Path: path,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) PlanarCheck(w http.ResponseWriter, req *http.Request){
	id, err := getID(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := s.service.PlanarCheck(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := struct {
		IsPlanar bool `json:"isPlanar"`
	}{
		IsPlanar: res,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) HamiltonianPath(w http.ResponseWriter, req *http.Request) {
	s.path(w, req, s.service.HamiltonianPath)
}

func (s *Server) EulerianCycle(w http.ResponseWriter, req *http.Request) {
	s.path(w, req, s.service.EulerianCycle)
}

type pathF func(graphID, startedNode uint64) ([]model.Node, error)

func (s *Server) path(w http.ResponseWriter, req *http.Request, f pathF) {
	args, err := getPathArgs(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path, err := f(args.graphID, args.startedNode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := struct {
		Path []model.Node `json:"path"`
	}{
		Path: path,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) AllShortestPaths(w http.ResponseWriter, req *http.Request) {
	args, err := getShortestPathArgs(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path, err := s.service.AllShortestPaths(args.graphID, args.fromNode, args.toNode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := struct {
		Path [][]model.Node `json:"paths"`
	}{
		Path: path,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func getID(req *http.Request) (uint64, error) {
	return getSpecificID(req, "id")
}

type shortestPathArgs struct {
	graphID  uint64
	fromNode uint64
	toNode   uint64
}

func getShortestPathArgs(req *http.Request) (shortestPathArgs, error) {
	id, err := getID(req)
	if err != nil {
		return shortestPathArgs{}, err
	}
	fromNode, err := getSpecificID(req, "fromNode")
	if err != nil {
		return shortestPathArgs{}, err
	}
	toNode, err := getSpecificID(req, "toNode")
	if err != nil {
		return shortestPathArgs{}, err
	}
	return shortestPathArgs{
		graphID:  id,
		fromNode: fromNode,
		toNode:   toNode,
	}, nil
}

type pathArgs struct {
	graphID     uint64
	startedNode uint64
}

func getPathArgs(req *http.Request) (pathArgs, error) {
	id, err := getID(req)
	if err != nil {
		return pathArgs{}, err
	}
	startedNode, err := getSpecificID(req, "startedNode")
	if err != nil {
		return pathArgs{}, err
	}
	return pathArgs{
		graphID:     id,
		startedNode: startedNode,
	}, nil
}

func getSpecificID(req *http.Request, idName string) (uint64, error) {
	vars := mux.Vars(req)
	id, err := strconv.ParseUint(vars[idName], 10, 64)
	return id, err
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/illfate2/graph-api/pkg/repository"
	"github.com/illfate2/graph-api/pkg/server"
	"github.com/illfate2/graph-api/pkg/service"
)

const port = "PORT"

func main() {
	port := os.Getenv(port)
	if port == "" {
		log.Fatal("empty port")
	}

	repo := repository.New()
	graph := service.NewGraph(repo)
	s := server.New(graph)
	log.Print("Running on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, s))
}

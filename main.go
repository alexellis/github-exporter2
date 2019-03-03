package main

import (
	"log"

	"github.com/alexellis/github-exporter2/pkg/exporter"
	"github.com/alexellis/github-exporter2/pkg/server"
)

func main() {
	// repo := types.CreateRepo(nil)

	exporter1 := exporter.NewExporter()
	exporter.RegisterExporter(exporter1)

	s := server.Create()
	log.Fatal(s.ListenAndServe())
}

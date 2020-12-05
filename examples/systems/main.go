package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	log.Printf("Loading systems from: %#v", gbfs.SystemsNABSA)

	systems, err := c.LoadSystems(gbfs.SystemsNABSA)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	log.Printf("Success! Got %d systems: %#v", len(systems), systems)
}

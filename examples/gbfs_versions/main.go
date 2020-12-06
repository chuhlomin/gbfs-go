package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for gbfs_version
	url := "no example, sorry"

	log.Printf("Loading GTFS versions from URL %s", url)

	resp, err := c.LoadGBFSVersions(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, v := range resp.Data.Versions {
		log.Printf("- Version %s", v.Version)
		log.Printf("  URL: %s", v.URL)
	}
}

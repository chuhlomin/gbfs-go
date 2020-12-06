package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for system_regions
	url := "https://gbfs.nextbike.net/maps/gbfs/v1/nextbike_cs/hr/system_regions.json"

	log.Printf("Loading station regions from URL %s", url)

	resp, err := c.LoadSystemRegions(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, r := range resp.Data.Regions {
		log.Printf("- ID: %s", r.ID)
		log.Printf("  Name: %s", r.Name)
	}
}

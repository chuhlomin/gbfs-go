package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for vehicle_types
	url := "no example, sorry"

	log.Printf("Loading GBFS vehicle types from URL %s", url)

	resp, err := c.LoadVehicleTypes(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, t := range resp.Data.VehicleTypes {
		log.Printf("- Form factor: %s", t.FormFactor)
		log.Printf("  Propulsion type: %s", t.PropulsionType)
		log.Printf("  Max range meters: %f", t.MaxRangeMeters)
		log.Printf("  Name: %s", t.Name)
	}
}

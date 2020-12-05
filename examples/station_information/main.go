package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for station_information
	url := "https://gbfs.citibikenyc.com/gbfs/en/station_information.json"

	log.Printf("Loading station information from URL %s", url)

	resp, err := c.LoadStationInformation(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	if len(resp.Data.Stations) > 2 {
		log.Println("Showing only first 2 stations")
		resp.Data.Stations = resp.Data.Stations[0:2]
	}

	for _, s := range resp.Data.Stations {
		log.Printf("- ID: %s", s.ID)
		log.Printf("  Name: %s", s.Name)
		log.Printf("  Short name: %s", s.ShortName)
		log.Printf("  Location: [%.2f,%.2f]", s.Lat, s.Lon)
		log.Printf("  Capacity: %d", s.Capacity)
	}
}

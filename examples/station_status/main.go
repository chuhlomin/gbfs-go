package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for station_status
	url := "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

	log.Printf("Loading station status from URL %s", url)

	resp, err := c.LoadStationStatus(url)
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
		log.Printf("  Num bikes available: %d", s.NumBikesAvailable)
		log.Printf("  Num bikes disabled: %d", s.NumBikesDisabled)
		log.Printf("  Num docks available: %d", s.NumDocksAvailable)
		log.Printf("  Num docks disabled: %d", s.NumBikesDisabled)
		log.Printf("  Last reported: %s", s.LastReported.Time().In(loc).Format(time.RFC3339))
		log.Printf("  Is installed: %v", s.IsInstalled)
		log.Printf("  Is renting: %v", s.IsRenting)
		log.Printf("  Is returning: %v", s.IsReturning)
	}
}

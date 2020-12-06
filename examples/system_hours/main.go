package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for station_status
	url := "https://gbfs.nextbike.net/maps/gbfs/v1/nextbike_bi/en/system_hours.json"

	log.Printf("Loading station hours from URL %s", url)

	resp, err := c.LoadSystemHours(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, rh := range resp.Data.RentalHours {
		log.Printf("- User types: %v", rh.UserTypes)
		log.Printf("  Days: %v", rh.Days)
		log.Printf("  Start time: %v", rh.StartTime.String())
		log.Printf("  Stop time: %v", rh.EndTime.String())
	}
}

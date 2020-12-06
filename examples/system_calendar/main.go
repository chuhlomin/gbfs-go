package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for system_calendar
	url := "https://gbfs.citibikenyc.com/gbfs/en/system_calendar.json"

	log.Printf("Loading system calendar from URL %s", url)

	resp, err := c.LoadSystemCalendar(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, c := range resp.Data.Calendars {
		log.Printf("- Start: %d/%d/%d", c.StartMonth, c.StartDay, c.StartYear)
		log.Printf("  End: %d/%d/%d", c.EndMonth, c.EndDay, c.EndYear)
	}
}

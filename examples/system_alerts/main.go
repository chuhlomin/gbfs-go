package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for system_alerts
	url := "https://s3.amazonaws.com/lyft-lastmile-production-iad/lbs/dca/system_alerts.json"

	log.Printf("Loading system alerts from URL %s", url)

	resp, err := c.LoadSystemAlerts(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, a := range resp.Data.Alerts {
		log.Printf("- Description: %s", a.Description)
		log.Printf("  Summary: %s", a.Summary)
		log.Printf("  ID: %s", a.ID)
		log.Printf("  Type: %s", a.Type)
		log.Printf("  Last updated: %s", a.LastUpdated.Time().In(loc).Format(time.RFC3339))

		if len(a.Times) > 0 {
			log.Println("  Times:")
			for _, t := range a.Times {
				log.Printf("    - Time start: %s", t.Start.String())
				log.Printf("      Time end: %s", t.End.String())
			}
		}
	}
}

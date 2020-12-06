package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for free_bike_status
	url := "https://gbfs.nextbike.net/maps/gbfs/v1/nextbike_lv/lv/free_bike_status.json"

	log.Printf("Loading free bike status from URL %s", url)

	resp, err := c.LoadFreeBikeStatus(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	if len(resp.Data.Bikes) > 2 {
		log.Println("Showing only first 2 bikes")
		resp.Data.Bikes = resp.Data.Bikes[0:2]
	}

	for _, b := range resp.Data.Bikes {
		log.Printf("- ID: %s", b.BikeID)
		log.Printf("  Location: [%.2f, %.2f]", b.Lat, b.Lon)
		log.Printf("  Is reserved: %v", b.IsReserved)
		log.Printf("  Is disabled: %v", b.IsDisabled)
	}
}

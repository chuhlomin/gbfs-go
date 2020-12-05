package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystems to discover GBFS URLs
	url := "https://gbfs.citibikenyc.com/gbfs/gbfs.json"

	log.Printf("Loading GBFS from URL %s", url)

	gbfs, err := c.LoadGBFS(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated %s", gbfs.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL %d", gbfs.TTL)

	for language, data := range gbfs.Data {
		log.Printf("Feeds [%s]:", language)

		for _, feed := range data.Feeds {
			log.Printf(" - %s\t%s", feed.Name, feed.URL)
		}
	}
}

package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for system_information
	url := "https://gbfs.citibikenyc.com/gbfs/en/system_information.json"

	log.Printf("Loading system information from URL %s", url)

	resp, err := c.LoadSystemInformation(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	log.Printf("System ID: %s", resp.Data.URL)
	log.Printf("Name: %s", resp.Data.Name)
	log.Printf("Short Name: %s", resp.Data.ShortName)
	log.Printf("Email: %s", resp.Data.Email)
	log.Printf("Phone number: %s", resp.Data.PhoneNumber)
	log.Printf("Operator: %s", resp.Data.Operator)
	log.Printf("Start date: %s", resp.Data.StartDate.String())
}

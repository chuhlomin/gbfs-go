package main

import (
	"log"
	"time"

	"github.com/chuhlomin/gbfs-go"
)

func main() {
	c := gbfs.NewClient("github.com/chuhlomin/gbfs-go", 30*time.Second)

	// See c.LoadSystem to discover GBFS URL for system_pricing_plans
	url := "https://tucson-us.publicbikesystem.net/ube/gbfs/v1/en/system_pricing_plans"

	log.Printf("Loading system pricing plans from URL %s", url)

	resp, err := c.LoadSystemPricingPlans(url)
	if err != nil {
		log.Fatalf("ERROR: Failed to %v", err)
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("ERROR: Failed to load time location: %v", err)
	}

	log.Printf("Last updated: %s", resp.LastUpdated.Time().In(loc).Format(time.RFC3339))
	log.Printf("TTL: %d", resp.TTL)

	for _, p := range resp.Data.Plans {
		log.Printf("- ID: %s", p.ID)
		log.Printf("  URL: %s", p.URL)
		log.Printf("  Name: %s", p.Name)
		log.Printf("  Currency: %s", p.Currency)
		log.Printf("  Price: %s", p.Price)
		log.Printf("  Is taxable: %v", p.IsTaxable)
		log.Printf("  Description: %s", p.Description)

		if len(p.PerKmPricing) > 0 {
			log.Printf("  Per km pricing:")
			for _, pr := range p.PerKmPricing {
				log.Printf("    - Start: %d", pr.Start)
				log.Printf("      Rate: %f", pr.Rate)
				log.Printf("      Interval: %d", pr.Interval)
				log.Printf("      End: %d", pr.End)
			}
		}

		if len(p.PerMinPricing) > 0 {
			log.Printf("  Per min pricing:")
			for _, pr := range p.PerMinPricing {
				log.Printf("    - Start: %d", pr.Start)
				log.Printf("      Rate: %f", pr.Rate)
				log.Printf("      Interval: %d", pr.Interval)
				log.Printf("      End: %d", pr.End)
			}
		}
	}
}

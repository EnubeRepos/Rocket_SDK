package main

import (
	"log"
	"net/url"
	"time"

	rocketsdkgo "github.com/enuberepos/Rocket_SDK_Go"
)

func main() {
	e, _ := url.Parse("https://analytics.rocket.enube.me")

	c, err := rocketsdkgo.Login(e, "username", "password")
	if err != nil {
		log.Fatalf("Failed to log into client\n%s", err)
	}

	// Get usage of current month
	usage, err := c.GetCurrentUsage()
	if err != nil {
		log.Fatal("Failed to get current usage\n%s", err)
	}

	log.Print(usage)

	// Get usage of specific month
	usage, err = c.GetUsageMonth(2024, time.February)
	if err != nil {
		log.Fatalf("Failed to get monthly usage\n%s", err)
	}

	log.Print(usage)

	// Get usage of specific period
	start, _ := time.Parse(time.DateOnly, "2024-01-01")
	end, _ := time.Parse(time.DateOnly, "2024-02-30")

	usage, err = c.GetUsagePeriod(start, end)
	if err != nil {
		log.Fatalf("Failed to get period usage\n%s", err)
	}

	log.Print(usage)
}

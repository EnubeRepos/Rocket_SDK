package main

import (
	"flag"
	"log"
	"net/url"
	"os"

	"github.com/enuberepos/Rocket_SDK_Go/client"
)

var (
	user     = flag.String("user", os.Getenv("ROCKET_API_USER"), "The username used for login, defaults to ROCKET_API_USER environment variable")
	passwd   = flag.String("passwd", os.Getenv("ROCKET_API_PASSWD"), "The password for the user login, defaults to ROCKET_API_PASSWD environment variable")
	endpoint = flag.String("api", os.Getenv("ROCKET_API_URL"), "The endpoint of the Rocket API, must be a valid URL, defaults to ROCKET_API_URL environment variable")
)

func init() {
	flag.Parse()
}

func main() {
	e, err := url.Parse(*endpoint)
	if err != nil {
		log.Fatalf("\"api\" is not a valid url!\n%s", err)
	}

	c := client.NewClient(e)

	if err := c.Login(*user, *passwd); err != nil {
		log.Fatalf("Failed to log into client\n%s", err)
	}

	log.Print("Succefully logged into the Rocket API")
}

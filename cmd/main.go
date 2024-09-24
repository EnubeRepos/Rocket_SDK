package main

import (
	"errors"
	"log"
	"net/url"
	"os"

	"github.com/leaanthony/clir"
)

const version = "v0.1.0"

type rootArgs struct {
	User     string `name:"user" description:"The username to log into. Defaults to ROCKET_API_USER"`
	Password string `name:"passwd" description:"The password of the user to log into. Defaults to ROCKET_API_PASSWD"`
	Endpoint string `name:"endpoint" description:"The endpoint of the Rocket v4 API. Defaults to ROCKET_API_ENDPOINT"`
}

var root = clir.NewCli("rocket", "the Rocket CLI", version)
var rootA = &rootArgs{
	User:     os.Getenv("ROCKET_API_USER"),
	Password: os.Getenv("ROCKET_API_PASSWD"),
	Endpoint: os.Getenv("ROCKET_API_ENDPOINT"),
}

func init() {
	root.AddFlags(rootA)
	root.PreRun(func(c *clir.Cli) error {
		if rootA.User == "" {
			return errors.New("Argument \"user\" or environment ROCKET_API_USER is not set")
		}
		if rootA.Password == "" {
			return errors.New("Argument \"passwd\" or environment ROCKET_API_PASSWD is not set")
		}
		if rootA.Endpoint == "" {
			return errors.New("Argument \"endpoint\" or environment ROCKET_API_ENDPOINT is not set")
		}

		if _, err := url.Parse(rootA.Endpoint); err != nil {
			return errors.Join(errors.New("Argument \"endpoint\" or environment ROCKET_API_ENDPOINT are not a valid URL"), err)
		}

		return nil
	})
}

func main() {
	if err := root.Run(); err != nil {
		log.Fatal(err)
	}
}

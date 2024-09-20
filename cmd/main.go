package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"time"

	rocketsdkgo "github.com/enuberepos/Rocket_SDK_Go"
	"github.com/enuberepos/Rocket_SDK_Go/api"
)

// TODO: THIS NEEDS A WHOLE REFACTOR

var (
	user     = flag.String("user", os.Getenv("ROCKET_API_USER"), "The username used for login, defaults to ROCKET_API_USER environment variable")
	passwd   = flag.String("passwd", os.Getenv("ROCKET_API_PASSWD"), "The password for the user login, defaults to ROCKET_API_PASSWD environment variable")
	endpoint = flag.String("api", os.Getenv("ROCKET_API_ENDPOINT"), "The endpoint of the Rocket API, must be a valid URL, defaults to ROCKET_API_URL environment variable")
)

var (
	getUsagePeriod  = flag.NewFlagSet("get-usage-period", flag.ExitOnError)
	getUsageMonth   = flag.NewFlagSet("get-usage-month", flag.ExitOnError)
	getUsageCurrent = flag.NewFlagSet("get-usage-current", flag.ExitOnError)
)

func init() {
	flag.Parse()
}

func main() {
	e, err := url.Parse(*endpoint)
	if err != nil {
		log.Fatalf("\"api\" is not a valid url!\n%s", err)
	}

	c, err := rocketsdkgo.LoginAzure(e, *user, *passwd)
	if err != nil {
		log.Fatalf("Failed to log into client\n%s", err)
	}

	var usage api.Indicator

	args := flag.Args()

	if len(args) == 0 {
		log.Fatalf("No subcommand was provided")
	}

	switch args[0] {
	case "get-usage-period":
		start := getUsagePeriod.String("start", time.Now().Format(time.DateOnly), "The start date of the query, must be a date in format YYYY-MM-DD")
		end := getUsagePeriod.String("end", time.Now().Format(time.DateOnly), "The end date of the query, must be a date in format YYYY-MM-DD")

		err := getUsagePeriod.Parse(args[1:])
		if err != nil {
			log.Fatalf("Failed to parse flags\n%s", err)
		}

		s, err := time.Parse(time.DateOnly, *start)
		if err != nil {
			log.Fatalf("Failed to parse flag \"start\", not in valid date format\n%s", err)
		}

		e, err := time.Parse(time.DateOnly, *end)
		if err != nil {
			log.Fatalf("Failed to parse flag \"end\", not in valid date format\n%s", err)
		}

		usage, err = c.GetIndicatorsPeriod(s, e)
		if err != nil {
			log.Fatalf("Failed to get usage data due to\n%s", err)
		}
	case "get-usage-month":
		month := getUsageMonth.String("month", time.Now().Format("2006-01"), "The month to be queried, must be in format YYYY-MM")

		err := getUsageMonth.Parse(args[1:])
		if err != nil {
			log.Fatalf("Failed to parse flags\n%s", err)
		}

		m, err := time.Parse("2006-01", *month)
		if err != nil {
			log.Fatalf("Failed to parse flag \"month\", not in valid date format\n%s", err)
		}

		_, err = c.GetUsageMonth(m.Year(), m.Month())
		if err != nil {
			log.Fatalf("Failed to get usage data due to\n%s", err)
		}
	case "get-usage-current":
		_ = getUsageCurrent

		_, err = c.GetUsageCurrent()
		if err != nil {
			log.Fatalf("Failed to get usage data due to\n%s", err)
		}
	default:
		log.Fatalf("Not a valid subcommand was provided")
	}

	b, err := json.Marshal(usage)
	if err != nil {
		log.Fatalf("Failed to convert data to json due to\n%s", err)
	}

	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Fatalf("Failed to stdout due to\n%s", err)
	}
}

package main

import (
	"encoding/json"
	e "errors"
	"fmt"
	"net/url"
	"os"
	"time"

	rocketsdkgo "github.com/enuberepos/Rocket_SDK_Go"
	"github.com/enuberepos/Rocket_SDK_Go/api"
	"github.com/enuberepos/Rocket_SDK_Go/client"
	"github.com/leaanthony/clir"
)

type provider string

var (
	AWS   provider = "aws"
	AZURE provider = "azure"
	GCP   provider = "gcp"
)

func wrapAction(action func(client.Client) error, provider provider) func() error {
	return func() error {
		var client client.Client
		var err error

		// We can ignore the error since the string is checked on PreRun
		endpoint, _ := url.Parse(rootA.Endpoint)

		switch provider {
		case AWS:
			client, err = rocketsdkgo.LoginAws(endpoint, rootA.User, rootA.Password)
		case AZURE:
			client, err = rocketsdkgo.LoginAzure(endpoint, rootA.User, rootA.Password)
		case GCP:
			client, err = rocketsdkgo.LoginGcp(endpoint, rootA.User, rootA.Password)
		}

		if err != nil {
			return e.Join(fmt.Errorf("Failed to login into provider %s API", provider), err)
		}

		return action(client)
	}
}

func setCommonCommands(c *clir.Command, provider provider) {
	getUsage := c.NewSubCommand("usage", "")
	getUsage.AddFlags(getUsageFlags)
	getUsage.Action(wrapAction(getUsageAction, provider))
}

type getUsageFlagsType struct {
	Month string `name:"month" description:"The month to get the usage from, must be in format YYYY-MM"`
	Start string `name:"start" description:"The start day to get the usage from, must be in format YYYY-MM-DD and defined alongside \"end\""`
	End   string `name:"end" description:"The end day to get the usage from, must be in format YYYY-MM-DD and defined alongside \"start\""`
}

var getUsageFlags = &getUsageFlagsType{}

func getUsageAction(c client.Client) error {
	var usage []api.Usage
	var err error

	if getUsageFlags.Start != "" && getUsageFlags.End != "" {
		var startTime, endTime time.Time

		startTime, err = time.Parse(time.DateOnly, getUsageFlags.Start)
		if err != nil {
			return e.Join(e.New("\"start\" is not in valid format YYYY-MM-DD"), err)
		}
		endTime, err = time.Parse(time.DateOnly, getUsageFlags.End)
		if err != nil {
			return e.Join(e.New("\"end\" is not in valid format YYYY-MM-DD"), err)
		}

		usage, err = c.GetUsagePeriod(startTime, endTime)

	} else if getUsageFlags.Month != "" {
		var t time.Time
		t, err = time.Parse("2006-01", getUsageFlags.Month)
		if err != nil {
			return e.Join(e.New("\"month\" is not in valid format YYYY-MM"), err)
		}
		usage, err = c.GetUsageMonth(t.Year(), t.Month())

	} else {
		usage, err = c.GetUsageCurrent()
	}

	if err != nil {
		return e.Join(e.New("Failed to get usage data"), err)
	}

	b, err := json.Marshal(usage)
	if err != nil {
		return e.Join(e.New("Failed parse usage data to JSON"), err)
	}

	if _, err := os.Stdout.Write(b); err != nil {
		return e.Join(e.New("Failed to write data to STDOUT"), err)
	}

	return nil
}

package main

import "github.com/enuberepos/Rocket_SDK_Go/client"

var gcpClient *client.GcpClient

var gcp = root.NewSubCommandInheritFlags("gcp", "commands related to GCP")

func init() {
	setCommonCommands(gcp, GCP)
}

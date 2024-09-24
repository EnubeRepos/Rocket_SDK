package main

import "github.com/enuberepos/Rocket_SDK_Go/client"

var awsClient *client.AwsClient

var aws = root.NewSubCommandInheritFlags("aws", "commands related to Aws")

func init() {
	setCommonCommands(aws, AWS)
}

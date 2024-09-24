package main

import "github.com/enuberepos/Rocket_SDK_Go/client"

var azureClient *client.AzureClient

var azure = root.NewSubCommandInheritFlags("azure", "commands related to Azure")

func init() {
	setCommonCommands(azure, AZURE)
}

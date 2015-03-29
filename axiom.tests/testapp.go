package main

import (
	"axiom"
	"axiom.tests/config"
)

func main() {

	// Start the app
	axiom.Init(config.AppConfig, config.RouteConfig)
}

package main

import (
	"axiom"
	"axiom.tests/config"
)

// Entry point of the app.
// routes are defined under config/routes.go
// controllers are defined under the controllers directory
// e.g controllers/homecontroller.go
func main() {

	// Register route table
	mux := axiom.Bind(config.RouteConfig)

	// Run the app
	axiom.Serve(mux)
}

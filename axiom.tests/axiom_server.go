package main

import (
	"axiom"
	"axiom.tests/controllers"
)

// Test 3/25/15 2:12 PM CST

// Controllers should be defined in their appropriate sub directory
func main() {

	// Create a route and assign a controller to it
	route := axiom.NewRoute("Default", controllers.HomeController)

	// Add to the route table
	routes := []axiom.Route{*route}

	// Register route table
	mux := axiom.Bind(routes)

	// Run the app
	axiom.Serve(mux) // OK
}

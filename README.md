# Axiom: A web MVC framework for Go.

### Progress
- - 
This is a brand new pet project, only a few days old.

So far, we are able to:
1. Define routes, controllers, handlers
2. Render views
3. Render JSON


### Example
- - - 

#### App main
```
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
```
#### Example controller
```
package controllers

import (
	"axiom"
)

// Set up a controller called 'home' with three actions
// 1. home/welcome -> View
// 2. home/fetch -> JSON
// 3. home/hello -> View
var HomeController = &axiom.Controller{
	Name: "home",
	Actions: map[string]axiom.Action{
		"welcome": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) axiom.ActionResult {
				return axiom.View("index") // Should match view/index.html
			},
		},
		"fetch": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) axiom.ActionResult {
				var query = c.Params.Query
				return axiom.Json(query) // Echo back the query as JSON
			},
		},
		"hello": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) axiom.ActionResult {
				return axiom.View("hello") // Should match view/hello.html
			},
		},
	},
	// Redirect to this action by default
	ActionDefault: "welcome",
}
```
#### Route configuration
```
package config

import (
	"axiom"
	"axiom.tests/controllers"
)

// Route configuration should be defined here.
var RouteConfig = []axiom.Route{
	axiom.Route{
		Name:       "Default",
		Url:        "/",
		Controller: controllers.HomeController,
	},
}
```

### Contributing
- - -

Welcome, but not expected =P





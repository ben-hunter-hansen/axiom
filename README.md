# Axiom: A web MVC framework for Go.

### Disclaimer
- - -

This is a brand new pet project, only a few days old.

The only existing functionality is a very basic "route table",
and ability to apply custom action handlers to each controller.

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
// Set up a controller called 'home' with three actions
// 1. home/welcome
// 2. home/query
// 3. home/hello
var HomeController = &axiom.Controller{
	Name: "home",
	Actions: map[string]axiom.Action{
		"welcome": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) (int, error) {
				http.ServeFile(c.Out, c.Request, "views/index.html")
				return http.StatusOK, nil
			},
		},
		"query": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) (int, error) {
				fmt.Fprintf(c.Out, "id: %s", c.Params.Query.Get("id"))
				return http.StatusOK, nil
			},
		},
		"hello": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) (int, error) {
				http.ServeFile(c.Out, c.Request, "views/hello.html")
				return http.StatusOK, nil
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





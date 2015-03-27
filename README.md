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
package controllers

import (
	"axiom"
	"net/http"
)

// Each controller should be defined in it's own file within the
// controllers package.
var HomeController = &axiom.Controller{
	Name: "home",
	Actions: map[string]axiom.Action{
		"welcome": axiom.Action{
			Name:   "welcome",
			Method: "GET",
			Handler: func(c *axiom.Controller) (int, error) {
				http.ServeFile(c.Out, c.Request, "views/index.html")
				return http.StatusOK, nil
			},
		},
		"hello": axiom.Action{
			Name:   "hello",
			Method: "GET",
			Handler: func(c *axiom.Controller) (int, error) {
				http.ServeFile(c.Out, c.Request, "views/hello.html")
				return http.StatusOK, nil
			},
		},
		"hello/again": axiom.Action{
			Name:   "hello/again",
			Method: "GET",
			Handler: func(c *axiom.Controller) (int, error) {
				http.ServeFile(c.Out, c.Request, "views/hello2.html")
				return http.StatusOK, nil
			},
		},
	},
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





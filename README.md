# Axiom: A web MVC framework for Go.

### Progress
- - -

This is a brand new pet project, only a few days old.

What is included so far:

* Define routes, controllers, handlers
* Actions: View, JSON
* Mapped resource directory
* Logging system in the works

### Example
- - - 

#### App main
```
package main

import (
	"axiom"
	"axiom.tests/config"
)

func main() {

	// Start the app
	axiom.Init(config.AppConfig, config.RouteConfig)
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
			Handler: func(c *axiom.Controller) axiom.ActionType {
				return axiom.View("index") // Should match view/index.html
			},
		},
		"fetch": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) axiom.ActionType {
				var query = c.Params.Query
				return axiom.Json(query) // Echo back the query as JSON
			},
		},
		"hello": axiom.Action{
			Method: "GET",
			Handler: func(c *axiom.Controller) axiom.ActionType {
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
#### App configuration
```
package config

import "axiom"

// Default application configuration
var AppConfig = &axiom.AppConfiguration{
	AppDir: axiom.AppDirectory{
		ViewRoot:     "views",
		ResourceRoot: "resources",
		Extensions: []string{
			".htm", ".html", ".xhtml",
		},
	},
}
```

### Contributing
- - -

Welcome, but not expected =P





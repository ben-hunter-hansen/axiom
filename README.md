# Axiom: A web MVC framework for Go.

### Disclaimer
- - -

This is a brand new pet project, only a few days old.

The only existing functionality is a very basic "route table",
and ability to apply custom action handlers to each controller.

### Example
- - - 

```

package main

import (
	//"fmt"
	"net/http"

	"axiom"
)

// Test 3/24/15 9:45 PM CST

// Should be able to define a handler anywhere.
func index(c *axiom.Controller) (int, error) {
	http.ServeFile(c.Out, c.Request, "views/index.html") // OK
	return http.StatusOK, nil                            // OK
}

// Should set up a route with a controller with a couple actions
func main() {
	indexAction := axiom.NewAction("welcome", "GET") // mapped to: GET /welcome
	indexAction.Handler = index

	// Create a couple actions
	helloAction := axiom.NewAction("hello", "GET")         //mapped to: GET /hello
	anotherAction := axiom.NewAction("hello/again", "GET") //mapped to: GET /hello/again

	// Assign a couple anonymous function handlers
	helloAction.Handler = func(c *axiom.Controller) (int, error) {
		http.ServeFile(c.Out, c.Request, "views/hello.html")
		return http.StatusOK, nil
	}
	anotherAction.Handler = func(c *axiom.Controller) (int, error) {
		http.ServeFile(c.Out, c.Request, "views/hello2.html")
		return http.StatusOK, nil
	}

	// Create a controller and assign the actions
	// mapped to: /home/{action}
	homeCtrl := axiom.NewController("home", []axiom.Action{
		*indexAction,
		*helloAction,
		*anotherAction,
	})

	// Create a route and assign a controller to it, then add to route table
	route := axiom.NewRoute("Default", homeCtrl)
	routes := []axiom.Route{*route}

	// Register route table and controllers
	mux := axiom.Bind(routes)
	axiom.Serve(mux) // OK
}

```

### Contributing
- - -

Welcome, but not expected =P





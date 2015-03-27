package controllers

import (
	"axiom"
	"fmt"
	"net/http"
)

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

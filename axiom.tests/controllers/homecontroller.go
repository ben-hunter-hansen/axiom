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

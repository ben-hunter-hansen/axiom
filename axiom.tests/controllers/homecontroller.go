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

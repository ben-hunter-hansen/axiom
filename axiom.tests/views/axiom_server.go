package main

import (
	//"fmt"
	"net/http"

	"axiom"
)

func hello(c *axiom.Controller) (int, error) {
	http.ServeFile(c.Out, c.Request, "hello.html")
	return http.StatusOK, nil
}

func main() {
	helloAction := axiom.NewAction("hello", "GET")
	helloAction.Handler = hello

	homeCtrl := axiom.NewController("home", []axiom.Action{
		*helloAction,
	})
	route := axiom.NewRoute("Default", homeCtrl)
	routes := []axiom.Route{*route}
	mux := axiom.Bind(routes)
	axiom.Serve(mux)
}

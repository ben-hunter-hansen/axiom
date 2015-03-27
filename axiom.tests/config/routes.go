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

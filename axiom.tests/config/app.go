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

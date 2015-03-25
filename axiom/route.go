package axiom

import (
//"net/http"
)

//TODO: DOCUMENT

// A Route has a base URI, and a controller to handle all requests
// under the base UI.
type Route struct {
	Name       string
	Url        string
	Children   []string
	Controller *Controller
}

func NewRoute(name string, c *Controller) *Route {
	url := "/" + c.Name + "/"
	var children []string
	for k, _ := range c.Actions {
		children = append(children, k)
	}
	return &Route{
		Name:       name,
		Url:        url,
		Children:   children,
		Controller: c,
	}
}

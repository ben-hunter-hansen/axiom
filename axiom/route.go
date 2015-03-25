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
	Controller *Controller
}

func NewRoute(name string, c *Controller) *Route {
	url := "/" + c.Name + "/"
	return &Route{
		Name:       name,
		Url:        url,
		Controller: c,
	}
}

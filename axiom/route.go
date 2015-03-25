package axiom

import (
	"net/http"
)

// A route has a name, base URL, child routes, and a controller.
type Route struct {
	Name       string
	Url        string
	Children   []string
	Controller *Controller
}

// Creates a new route for a given controller
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

// The ServeHTTP method of a Route will receive all requests that
// fall under it's base URI component and will determine the appropriate
// action to invoke.
func (r Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqUrl := req.URL.Path
	r.Controller.Request = req
	r.Controller.Out = w
	for name, action := range r.Controller.Actions {
		if reqUrl == r.Url+name {
			status, err := action.Handler(r.Controller)
			if err != nil {
				switch status {
				default:
					http.ServeFile(w, req, "500.html")
				}
			}
			return
		}
	}
	http.ServeFile(w, req, "views/404.html")
}

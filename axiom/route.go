package axiom

import (
	"net/http"
	"strings"
)

// A route has a name, base URL, and a controller
type Route struct {
	Name       string
	Url        string
	Controller *Controller
}

// The ServeHTTP method of a Route will receive all requests that
// fall under it's base URI component and will determine the appropriate
// action to invoke.
//
// TOOD: Need to find a smarter way to handle this mechanism
func (r Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Controller.Request = req
	r.Controller.Params = RequestParams{
		Query: req.URL.Query(),
	}
	r.Controller.Out = w

	act := findAction(req.URL.Path, r)
	status, err := act.Handler(r.Controller)

	if err != nil {
		switch status {
		case http.StatusNotFound:
			http.ServeFile(w, req, "views/404.html")
		case http.StatusInternalServerError:
			http.ServeFile(w, req, "views/500.html")
		}
	}
}

// One-timer for figuring out which action to take.
// Hopefully this will be gone soon
func findAction(url string, r Route) Action {
	if url == r.Url {
		def := r.Controller.ActionDefault
		return r.Controller.Actions[def]
	}
	for name, action := range r.Controller.Actions {
		actionPath := r.Url + r.Controller.Name + "/" + name
		if url == strings.ToLower(actionPath) {
			return action
		}
	}
	return Action{
		Handler: func(c *Controller) (int, error) {
			return http.StatusNotFound, nil
		},
	}
}

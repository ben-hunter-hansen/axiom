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
	r.Controller.Apply(w, req, RequestParams{
		Query: req.URL.Query(),
	})
	var act = Action{}
	if strings.ToLower(r.Url) == strings.ToLower(req.URL.Path) {
		act = r.Controller.Actions[r.Controller.ActionDefault]
	} else {
		act = r.Controller.GetAction(r.Url, req.URL.Path)
	}
	result := act.Handler(r.Controller)

	// Need to define custom resource types ASAP
	switch res := result.Resource.(type) {
	case string: // Should be a view
		http.ServeFile(w, req, res)
	case []byte: // Should be json
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

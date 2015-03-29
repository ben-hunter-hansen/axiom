package axiom

import (
	"net/http"
	"strings"
)

// A Controller receives requests and determine which
// action should be invoked.
//
// A Controller has:
// 1. A name which will identify the controller within the route table.
// 2. the request in question
// 3. A writer for responding to the requestee
// 4. A map containing the controllers actions
//    k: Action URL value (name)
//    v: Handler function
// 5. Default action name
type Controller struct {
	Name          string
	Request       *http.Request
	Params        RequestParams
	Out           http.ResponseWriter
	Actions       map[string]Action
	ActionDefault string
}

// Provides the controller with request context and a corresponding writer
func (c *Controller) Apply(w http.ResponseWriter, r *http.Request, p RequestParams) {
	c.Out = w
	c.Request = r
	c.Params = p
}

// Checks for an action matching a URL.
// Returns the action if found, else 404
func (c *Controller) GetAction(routeUrl string, reqUrl string) Action {
	routeUrl = strings.ToLower(routeUrl)
	reqUrl = strings.ToLower(reqUrl)
	for name, _ := range c.Actions {
		fullPath := routeUrl + c.Name + "/" + name
		if fullPath == reqUrl {
			return c.Actions[name]
		}
	}
	return Action{
		Handler: func(c *Controller) ActionType {
			return View("404")
		},
	}
}

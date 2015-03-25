package axiom

import "net/http"

// A Controller receives requests and determine which
// action should be invoked.
//
// A Controller has:
// 1. A name which will identify the controller within the route table.
// 2. the request in question
// 3. A writer for responding to the requestee
// 4. A list of actions that the controller knows how to do.
type Controller struct {
	Name    string
	Request *http.Request
	Out     http.ResponseWriter
	Actions []Action
}

// Returns a named controller with an action list
func NewController(name string, a []Action) *Controller {
	return &Controller{
		Name:    name,
		Actions: a,
	}
}

// An Action is something a Controller is configured to handle
//
// 1. A name which the controller can identify
// 2. The actions request method.
// 3. User defined handler.
type Action struct {
	Name    string
	Method  string
	Handler func(*Controller) (int, error)
}

func NewAction(name string, method string) *Action {
	return &Action{
		Name:   name,
		Method: method,
	}
}

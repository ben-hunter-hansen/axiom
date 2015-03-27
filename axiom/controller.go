package axiom

import (
	"net/http"
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

// An Action is something a Controller is configured to handle
//
// 2. The actions request method.
// 3. User defined handler.
type Action struct {
	Method  string
	Handler func(*Controller) (int, error)
}

package axiom

// An Action is something a Controller is configured to handle
//
// 2. The actions request method.
// 3. User defined handler.
type Action struct {
	Method  string
	Handler func(*Controller) ActionResult
}

// The result of an actions handler function
// 1. Error, obvious
// 2. HTTP status code
// 3. Resource, e.g a view, some json, etc
type ActionResult struct {
	Error    error
	Status   int
	Resource interface{}
}

func NewActionResult(status int, res interface{}, err error) ActionResult {
	return ActionResult{
		Error:    err,
		Status:   status,
		Resource: res,
	}
}

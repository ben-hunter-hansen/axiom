package axiom

import (
	"encoding/json"
)

type ActionType interface{}

// An Action is something a Controller is configured to handle
//
// 2. The actions request method.
// 3. User defined handler.
type Action struct {
	Method  string
	Handler func(*Controller) ActionType
}

// An action associated with a view
type ViewAction struct {
	FullPath string
	Name     string
}

// An action associated with json
type JsonAction struct {
	Data []byte
}

// Returns a new action, of type view
func View(v string) ActionType {
	return ViewAction{
		Name: v,
	}
}

// Returns a new action, of type json
func Json(v interface{}) ActionType {
	dat, err := json.Marshal(v)
	if err != nil {
		Logger.Error.Printf("%s\n", err.Error())
	}
	return JsonAction{Data: dat}
}

package axiom

import (
	"encoding/json"
	"net/http"
	"os"
)

const ViewNotFound = "views/404.html"
const ViewInternalError = "views/500.html"

// Try to find the requested view resource
// Indicate success or failure in the action result
func View(v string) ActionResult {
	path := "views/" + v + ".html" // For testing this module
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return NewActionResult(http.StatusNotFound, ViewNotFound, err)
	}
	return NewActionResult(http.StatusOK, path, nil)
}

// Encode json and include it within the action result
func Json(v interface{}) ActionResult {
	dat, err := json.Marshal(v)
	if err != nil {
		return NewActionResult(http.StatusInternalServerError, ViewInternalError, err)
	}
	return NewActionResult(http.StatusOK, dat, nil)
}

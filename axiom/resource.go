package axiom

import (
	"net/http"
	"strings"
)

// Provides mapping and a ServeHTTP method for
// a set of resources located under a common directory
type ResourceMapping struct {
	Url       string
	Resources []string
}

// Returns a new resource map
func NewResourceMap(path string, res []string) ResourceMapping {
	return ResourceMapping{
		Url:       "/" + path + "/",
		Resources: res,
	}
}

// Serves up the requested resource
func (rm ResourceMapping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, file := range rm.Resources {
		if strings.HasSuffix(file, r.URL.Path) {
			http.ServeFile(w, r, file)
		}
	}
}

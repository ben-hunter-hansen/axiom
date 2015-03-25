// Axiom MVC framework core library
// Author: Ben Hansen
// http://www.bensdevblog.com
package axiom

import (
	"net/http"
)

func Bind(r []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range r {
		// TODO: handle default "/"
		mux.Handle(route.Url, route)
	}
	return mux
}

func Serve(mx *http.ServeMux) {
	http.ListenAndServe(":8080", mx)
}

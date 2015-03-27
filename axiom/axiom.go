// Axiom MVC framework core library
// Author: Ben Hansen
// http://www.bensdevblog.com

package axiom

import "net/http"

// Goes through the route table and
// bind action handlers.
func Bind(r []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range r {
		mux.Handle(route.Url, route)
	}
	return mux
}

// (> o.o )>
func Serve(mx *http.ServeMux) {
	http.ListenAndServe(":8080", mx)
}

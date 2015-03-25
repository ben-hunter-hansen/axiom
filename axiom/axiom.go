// Axiom MVC framework core library
// Author: Ben Hansen
// http://www.bensdevblog.com
package axiom

import (
	"net/http"
)

// Internal handler for injecting application context to the user-
// defined handler function.
func (r Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqUrl := req.URL.Path
	r.Controller.Request = req
	r.Controller.Out = w
	for name, action := range r.Controller.Actions {
		if reqUrl == r.Url+name {
			status, err := action.Handler(r.Controller)
			if err != nil {
				switch status {
				default:
					http.ServeFile(w, req, "500.html")
				}
			}
			return
		}
	}
	http.ServeFile(w, req, "views/404.html")
}

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

// Axiom MVC framework core library
// Author: Ben Hansen
// http://www.bensdevblog.com

package axiom

import (
	"net/http"
)

func Init(a *AppConfiguration, r []Route) {
	mux := http.NewServeMux()
	a.AppDir.Update()

	Logger.Info.Println("Configuring HTTP handlers..")

	// Bind routes
	for _, route := range r {
		route.AppDir = &a.AppDir
		mux.Handle(route.Url, route)
	}

	// Bind resources
	for path, files := range a.AppDir.Resources {
		resMap := NewResourceMap(path, files)
		mux.Handle(resMap.Url, resMap)
	}

	Logger.Info.Printf("Server starting on port: %d\n", 8080)
	// ( > o.o)> [app]
	http.ListenAndServe(":8080", mux)
}

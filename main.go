package main

import (
	"github.com/gorilla/mux"
	"github.com/juanfgs/react-js-test/routes"
	"net/http"


)

func main() {
	mux := mux.NewRouter()

	/* Handle CSS and JS paths */
	jsHandler := http.FileServer(http.Dir("./static/js/"))
	cssHandler := http.FileServer(http.Dir("./static/css/"))
	mux.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsHandler))
	mux.PathPrefix("/css/").Handler(http.StripPrefix("/css/",cssHandler))
	/* Load Routes */
	for _, route := range routes.Get() {
		mux.HandleFunc(route.Destination, route.Handler)
	}



	http.Handle("/",mux)
	http.ListenAndServe("localhost:6060", mux)

}

package main

import(
	"github.com/juanfgs/baseline/routes"
	"goji.io"
	"goji.io/pat"
	"net/http"
)



func main(){
	mux := goji.NewMux()

	/* Load Routes */
	for _,route := range routes.Get(){
		mux.HandleFuncC(pat.New(route.Destination),route.Handler)
	}

	/* Handle CSS and JS paths */
	mux.Handle(pat.New("/css/*"), http.FileServer(http.Dir("./static/css/")))
	mux.Handle(pat.New("/js/*"),  http.FileServer(http.Dir("./static/js/")))

	http.ListenAndServe("localhost:6060", mux)


}

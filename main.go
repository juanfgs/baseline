package main

import(
	"github.com/juanfgs/encweb/routes"
	"goji.io"
	"goji.io/pat"
	"net/http"
)



func main(){
	mux := goji.NewMux()
	for _,route := range routes.Get(){
		mux.HandleFuncC(pat.New(route.Destination),route.Handler)
	}
	http.ListenAndServe("localhost:6060", mux)


}

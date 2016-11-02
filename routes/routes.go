package routes

import(

	"github.com/juanfgs/baseline/handlers"
	"net/http"
)

type Route struct {
	Destination string
	Handler func( w http.ResponseWriter, r *http.Request) ()
}


func Get() []Route{
	return []Route{
		{ "/", handlers.FormHandler{}.Index },

	}
}

package routes

import(

	"github.com/juanfgs/cv/handlers"
	"net/http"
)

type Route struct {
	Destination string
	Handler func( w http.ResponseWriter, r *http.Request) ()
}


func Get() []Route{
	return []Route{
		{ "/", handlers.FormHandler{}.Index },
		{ "/tictactoe", handlers.FormHandler{}.Index },
		{ "/create", handlers.FormHandler{}.Post },
	}
}

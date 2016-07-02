package routes

import(
	"golang.org/x/net/context"
	"github.com/juanfgs/baseline/handlers"
	"net/http"
)

type Route struct {
	Destination string
	Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) ()
}


func Get() []Route{
	return []Route{
		{ "/", handlers.FormHandler{}.Index },

	}
}

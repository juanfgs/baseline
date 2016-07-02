package routes

import(
	"golang.org/x/net/context"
	"github.com/juanfgs/encweb/handlers"
	"net/http"
)

type Route struct {
	Destination string
	Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) ()
}


func Get() []Route{
	return []Route{
		{ "/:name", handlers.FormHandler{}.Index },
	}
}

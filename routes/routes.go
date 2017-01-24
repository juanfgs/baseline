package routes

import(

	"github.com/juanfgs/cv/handlers/admin"
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
		{ "/admin", admin.AdminHandler{}.Index },
		{ "/admin/sections", admin.SectionsHandler{}.Index },
	}
}

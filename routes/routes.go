package routes

import(

	"github.com/juanfgs/cv/handlers/admin"
	"github.com/juanfgs/cv/handlers"
	"net/http"
)

type Route struct {
	Destination string
	Handler func( w http.ResponseWriter, r *http.Request) ()
	Method string
}


func Get() []Route{
	return []Route{
		{ "/", handlers.FormHandler{}.Index, "GET" },
		{ "/admin", admin.AdminHandler{}.Index , "GET"},
		{ "/admin/sections", admin.SectionsHandler{}.Index, "GET" },
		{ "/admin/sections/create", admin.SectionsHandler{}.CreateForm, "GET" },
		{ "/admin/sections/create", admin.SectionsHandler{}.Create, "POST"},
		{ "/admin/sections/edit/{id}", admin.SectionsHandler{}.EditForm, "GET"},
		{ "/admin/sections/edit/{id}", admin.SectionsHandler{}.Edit, "POST"},
		{ "/admin/sections/delete/{id}", admin.SectionsHandler{}.Delete, "GET"},
		{ "/admin/items", admin.ItemsHandler{}.Index, "GET" },
		{ "/admin/items/create", admin.ItemsHandler{}.CreateForm, "GET" },
		{ "/admin/items/create", admin.ItemsHandler{}.Create, "POST"},
		{ "/admin/items/create", admin.ItemsHandler{}.Create, "POST"},
		{ "/data.json", handlers.JsonHandler{}.Index, "GET"},
	}
}

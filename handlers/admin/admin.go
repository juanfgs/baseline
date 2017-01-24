package admin

import (
	"net/http"
	"github.com/juanfgs/cv/handlers"
)


type AdminHandler struct {
	handlers.BaseHandler
}


func (t AdminHandler) Index(w http.ResponseWriter, r *http.Request){
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/index.tpl")
	t.Render(w,"")
}

package handlers

import (
	"net/http"
	_ "github.com/juanfgs/cv/models"
)

type FormHandler struct{
	BaseHandler
}

func (t FormHandler) Index(w http.ResponseWriter, r *http.Request) {
	t.AddTemplate( "layout.tpl")
	t.AddTemplate( "index.tpl")
	t.Render(w, "")
}

func (t FormHandler) Post(w http.ResponseWriter, r *http.Request) {

}

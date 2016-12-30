package handlers

import (
	"net/http"

)

type FormHandler struct{
	BaseHandler
}

func (t FormHandler) Index(w http.ResponseWriter, r *http.Request) {
	t.AddTemplate( "layout.tpl")
	t.AddTemplate( "index.tpl")
	t.Render(w, "")
}

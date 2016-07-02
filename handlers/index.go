package handlers

import (
	// "fmt"
        // "goji.io/pat"
	"net/http"
	"golang.org/x/net/context"
)

type FormHandler struct{
	BaseHandler
}

func (t FormHandler) Index(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	t.Render(w, "foo")
}

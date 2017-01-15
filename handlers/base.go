package handlers

import (
	"html/template"
	"net/http"
	"github.com/juanfgs/cv/config"
)




type BaseHandler struct {
	templates []string
	tpl *template.Template  /* Compiled template */
}



func (self *BaseHandler)  AddTemplate(templateName string){
	self.templates = append(self.templates, config.Values.AppPath + "views/" + templateName)
}

func (self BaseHandler) Render(w http.ResponseWriter, data interface{}) {
	var err error

	self.tpl, err = template.New("layout").ParseFiles(self.templates...)
	if err != nil{
		panic(err)
	}

	err = self.tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}


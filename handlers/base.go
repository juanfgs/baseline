package handlers

import (
	"html/template"
	"net/http"
	"log"
)




type BaseHandler struct {
	templates []string
	tpl *template.Template  /* Compiled template */
}



func (self *BaseHandler)  AddTemplate(templateName string){
	self.templates = append(self.templates, "views/" + templateName)
}

func (self BaseHandler) Render(w http.ResponseWriter, data interface{}) {
	var err error
	log.Println(self.templates)

	self.tpl, err = template.New("layout").ParseFiles(self.templates...)
	if err != nil{
		panic(err)
	}

	err = self.tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}


package handlers

import (
	"html/template"
	"net/http"
	"github.com/juanfgs/cv/config"
)




type BaseHandler struct {
	templates []string
	tpl *template.Template  /* Compiled template */
	Data map[string]interface{}
	Funcs map[string]interface{}
}

func (self *BaseHandler) AddFunc(name string, f interface{}){
	self.initializeFuncs()
	self.Funcs[name] = f
}

func (self *BaseHandler) initializeFuncs(){
	if len(self.Funcs) == 0 {
		self.Funcs = make(map[string]interface{})
	}
}

func (self *BaseHandler)  AddTemplate(templateName string){
	self.templates = append(self.templates, config.Values.AppPath + "views/" + templateName)
}

func (self BaseHandler) Render(w http.ResponseWriter, data interface{}) {
	var err error
	self.initializeFuncs()
	self.tpl, err = template.New("layout").Funcs(self.Funcs).ParseFiles(self.templates...)
	if err != nil{
		panic(err)
	}

	err = self.tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}


package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"

)

var tpl *template.Template

type BaseHandler struct {

}


func init(){
	templates,err := ioutil.ReadDir("./views/")
	var templateNames []string
	if err != nil{
		panic(err)
	}

	for _,x := range templates {
		templateNames = append(templateNames, "./views/" + x.Name())
	}

	tpl, err = template.New("layout").ParseFiles(templateNames...)
	if err != nil{
		panic(err)
	}

}



func (s BaseHandler) Render(w http.ResponseWriter, data interface{}) {
	err := tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}


package admin

import (
	"net/http"
	"github.com/juanfgs/cv/handlers"
	"github.com/juanfgs/cv/models"
	"fmt"
)


type SectionsHandler struct {
	handlers.BaseHandler
}


func (t SectionsHandler) Index(w http.ResponseWriter, r *http.Request){
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/sections/index.tpl")
	t.AddFunc("log", fmt.Sprintf)
	section := models.NewSection()
	sections := section.All()
	var sectionList []map[string]string
	t.Data = make(map[string]interface{})
 	for  sections.Next() {
		var id int64
		var name, icon, tagline, content string
		err := sections.Scan(&id, &name,&icon,&tagline,&content)
		

		if err != nil{
			panic(err)
		}
		
		var values map[string]string =  map[string]string{
			"Id" : fmt.Sprintf("%d",id),
			"Name": name,
			"Icon": icon,
			"Tagline": tagline,
			"Content": content,
		}
		sectionList = append(sectionList, values)
		
	}
	
	t.Data["SectionList"] = sectionList
	t.Data["Hola"] = "fooobar"
	t.Render(w,t.Data)
}

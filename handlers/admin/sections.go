package admin

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/juanfgs/cv/handlers"
	"github.com/juanfgs/cv/models"
	"strconv"
	"fmt"
	"log"
)


type SectionsHandler struct {
	handlers.BaseHandler
}


func (t SectionsHandler) Index(w http.ResponseWriter, r *http.Request){

	//	Adding templates	 
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/sections/index.tpl")
	t.AddFunc("log", fmt.Sprintf)
	// Instantiate section model
	section := models.NewSection()
	sections := section.All()
	var sectionList []map[string]string
	t.Data = make(map[string]interface{})
 	for  sections.Next() { // Loop through sections
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
	t.Render(w,t.Data)
}

func (t SectionsHandler) CreateForm(w http.ResponseWriter, r *http.Request){
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/sections/create.tpl")
	t.Render(w,"")
}

func (t SectionsHandler) Create( w http.ResponseWriter, r *http.Request){
	section := models.NewSection()

	
	section.Name = r.FormValue("Name")
	section.Icon = r.FormValue("Icon")
	section.Tagline = r.FormValue("Tagline")
	section.Content = r.FormValue("Content")
	
	_ = section.Create()
	http.Redirect(w,r, r.Referer(), http.StatusSeeOther )
}

func (t SectionsHandler ) Delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	section := models.NewSection()
	id , err := strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		log.Println("Invalid Id")
		return
	}
	
	section.FindById(id)
	section.Delete()
	http.Redirect(w,r, r.Referer(), http.StatusSeeOther )

}


func (t SectionsHandler) EditForm(w http.ResponseWriter, r *http.Request){
		t.Data = make(map[string]interface{})

	vars := mux.Vars(r)
	section := models.NewSection()
	id , err := strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		log.Println("Invalid Id")
		return
	}
	section.FindById(id)
	
	t.Data["Section"] = section
	
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/sections/edit.tpl")
	t.Render(w,t.Data)
}



func (t SectionsHandler) Edit( w http.ResponseWriter, r *http.Request){
	section := models.NewSection()
	id , err := strconv.ParseInt(r.FormValue("Id"),10,64)
	if err != nil {
		log.Println("Invalid Id")
		return
	}
	section.FindById(id)
	section.Name = r.FormValue("Name")
	section.Icon = r.FormValue("Icon")
	section.Tagline = r.FormValue("Tagline")
	section.Content = r.FormValue("Content")
	
	_ = section.Save()
	http.Redirect(w,r, r.Referer(), http.StatusSeeOther )
}

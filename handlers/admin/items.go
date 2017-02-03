package admin

import (
	"net/http"
	"github.com/juanfgs/cv/handlers"
	"github.com/juanfgs/cv/models"
	"fmt"
	"log"
	"strconv"
)


type ItemsHandler struct {
	handlers.BaseHandler
}


func (t ItemsHandler) Index(w http.ResponseWriter, r *http.Request){

	//	Adding templates	 
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/items/index.tpl")
	t.AddFunc("log", fmt.Sprintf)
	// Instantiate item model
	item := models.NewItem()
	items := item.All()
	var itemList []map[string]string
	t.Data = make(map[string]interface{})
	for  items.Next() { // Loop through items
		var id int64
		var sectionId int64
		var name, value string
		err := items.Scan(&id, &sectionId,&name, &value)
		

		if err != nil{
			panic(err)
		}
		
		var values map[string]string =  map[string]string{
			"Id" : fmt.Sprintf("%d",id),
			"SectionId": fmt.Sprintf("%d",sectionId),
			"Name": name,
			"Value": value,
		}
		itemList = append(itemList, values)
		
	}
	
	t.Data["ItemList"] = itemList
	t.Render(w,t.Data)
}

func (t ItemsHandler) CreateForm(w http.ResponseWriter, r *http.Request){
	t.AddTemplate("admin/layout.tpl")
	t.AddTemplate("admin/items/create.tpl")
	
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

func (t ItemsHandler) Create( w http.ResponseWriter, r *http.Request){
	item := models.NewItem()
	var err error

	
	item.Name = r.FormValue("Name")
	item.SectionId, err = strconv.Atoi(r.FormValue("SectionId"))
	if err != nil {
		log.Println(err)
		log.Println("Invalid section ID")
	}
	item.Value = r.FormValue("Value")
	
	_ = item.Create()
	http.Redirect(w,r, r.Referer(), http.StatusSeeOther )
}

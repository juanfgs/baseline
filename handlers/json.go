package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/juanfgs/cv/models"
	"fmt"
	"log"
	"strconv"
)

type JsonHandler struct{
	BaseHandler
}

func (t JsonHandler) Index(w http.ResponseWriter, r *http.Request) {

	section := models.NewSection()
	sections := section.All()
	var sectionList []map[string]interface{}
 	for  sections.Next() { // Loop through sections
		var id int64
		var name, icon, tagline, content string
		err := sections.Scan(&id, &name,&icon,&tagline,&content)
		
		if err != nil{
			panic(err)
		}

		item := models.NewItem()
		log.Println(id )
		items := item.AllWhere(" `section_id` = " + strconv.FormatInt(id,10) )
		log.Println(items)
		var itemList []map[string]string
		for  items.Next() { // Loop through items
			var id int64
			var sectionId int64
			var name, value string
			err := items.Scan(&id, &sectionId,&name, &value)
			
			
			if err != nil{
				panic(err)
			}
			
			var values map[string]string =  map[string]string{
				"id" : fmt.Sprintf("%d",id),
				"sectionId": fmt.Sprintf("%d",sectionId),
				"name": name,
				"value": value,
			}
			itemList = append(itemList, values)
			
		}
		

		var values map[string]interface{} =  map[string]interface{}{
			"id" : fmt.Sprintf("%d",id),
			"name": name,
			"icon": icon,
			"items": itemList,
			"tagline": tagline,
			"content": content,
		}
		sectionList = append(sectionList, values)
		
	}
	sectionsJson, err := json.Marshal(sectionList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(sectionsJson))
	
}


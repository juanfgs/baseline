package models

type Item struct {
	Model
	Id int64
	SectionId int
	Name string
	Value string
}


func NewItem() Item{
	return Item{ Model: Model{ tableName:"items", indexName: "id" } }
}

func (this *Item) Create() int64{
	stmt,err := DB.Prepare("INSERT INTO `" +this.tableName+ "` (`section_id`,`name`,`value`) VALUES (?,?,?)" )

	if err != nil{
		panic(err)
	}

	res, err := stmt.Exec(this.SectionId,this.Name,this.Value)
	this.Id, err = res.LastInsertId()
	if err != nil{
		panic(err)
	}

	return this.Id
}



func (this Item) Save() Item{
	stmt,err := DB.Prepare("UPDATE  `" +this.tableName+ "` SET `section_id` = ?, `name` = ?, `value` = ? WHERE `id` = ?" )
	if err != nil{
		panic(err)
	}

	stmt.Exec(this.SectionId,this.Name,this.Value,this.Id)
	return this
}

func (this *Item) FindById(id int64) {
	this.Model.FindById(id,&this.Id,&this.SectionId,&this.Name,&this.Value )
}

func (this Item) Delete(){
	this.Model.Delete(this.Id)
}

func (this *Item) FindByField(field string, value string){
	this.Model.FindByField(field, value,&this.Id,&this.SectionId,&this.Name,&this.Value )
}



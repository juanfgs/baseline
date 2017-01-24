package models


const (
	_ = iota
	TYPE_LIST
	TYPE_ACCORDION
	TYPE_PORTFOLIO
)

type Section struct {
	Model
	Id int64
	Name string
	Icon string
	Tagline string
	Content string
	itemType int
	Items []Item
}


func NewSection() Section{
	return Section{ Model: Model{ tableName:"sections", indexName: "id" } }
}

func (this *Section) Create() *Section{
	stmt,err := DB.Prepare("INSERT INTO `" +this.tableName+ "` (`name`,`icon`,`tagline`,`content`) VALUES (?,?,?,?)" )

	if err != nil{
		panic(err)
	}

	res, err := stmt.Exec(this.Name,this.Icon,this.Tagline,this.Content)
	this.Id, err = res.LastInsertId()
	if err != nil{
		panic(err)
	}

	return this
}



func (this Section) Save() Section{
	stmt,err := DB.Prepare("UPDATE  `" +this.tableName+ "` SET `name` = ?, `icon` = ?, `tagline` = ?, `content` = ? WHERE `id` = ?" )
	if err != nil{
		panic(err)
	}

	stmt.Exec(this.Name,this.Icon,this.Tagline,this.Content,this.Id)
	return this
}

func (this *Section) FindById(id int64) {
	this.Model.FindById(id,&this.Id,&this.Name,&this.Icon,&this.Tagline,&this.Content )
}


func (this *Section) FindByField(field string, value string){
	this.Model.FindByField(field, value,&this.Id,&this.Name,&this.Icon,&this.Tagline,&this.Content )
}

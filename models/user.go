package models

import(
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Model
	Id int64
	Name string
	password string
	Active bool
}


func NewUser() User{
	return User{ Model: Model{ tableName:"users", indexName: "id" } }
}

func (this *User) Create() int64{
	stmt,err := DB.Prepare("INSERT INTO `" +this.tableName+ "` (`name`,`password`,`active`) VALUES (?,?,?)" )

	if err != nil{
		panic(err)
	}

	if len(this.password) == 0 {
		panic("Password cannot be empty")
	}
	
	res, err := stmt.Exec(this.Name,this.password,this.Active)
	this.Id, err = res.LastInsertId()
	if err != nil{
		panic(err)
	}

	return this.Id
}



func (this User) Save() User{
	stmt,err := DB.Prepare("UPDATE  `" +this.tableName+ "` SET `name` = ?, `password` = ?, `active` = ? WHERE `id` = ?" )
	if err != nil{
		panic(err)
	}

	stmt.Exec(this.Name,this.password,this.Active,this.Id)
	return this
}

func (this *User) FindById(id int64) {
	this.Model.FindById(id,&this.Id,&this.Name,&this.password,&this.Active )
}

func (this User) Delete(){
	this.Model.Delete(this.Id)
}

func (this *User) FindByField(field string, value string){
	this.Model.FindByField(field, value,&this.Id,&this.Name,&this.password,&this.Active )
}

func (this *User) SetPassword( password string) {
	pass,err := bcrypt.GenerateFromPassword([]byte(password),0)
	
	
	if err != nil {
		log.Fatal(err)
	}

	this.password = string(pass)

}

func (this *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.password), []byte(password))
	
	
	if err != nil {
		return false
	}

	return true 
}

package models

import (
	"time"
)


type Post struct {
	Id int64
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
	Model
}


func NewPost() Post {
	post := Post{}
	post.SetIndex("id")
	post.SetTable("posts")

	return post
}


func (this *Post) FindById(id int64){
	this.Model.FindById(id, &this.Id, &this.Content, &this.CreatedAt)
}

func (this *Post) Save() {
	_,err := DB.Query("INSERT INTO `posts` (`content`, `created_at`) VALUES (?,?)", this.Content,this.CreatedAt)
	if err != nil {
		panic(err)
	}
}


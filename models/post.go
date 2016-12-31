package models

import (
	"time"
)


type Post struct {
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
	
}


func NewPost(content string) Post {
	return Post{ Content:content, CreatedAt: time.Now(), UpdatedAt: time.Now() }
}



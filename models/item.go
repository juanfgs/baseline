package models

type Item struct {
	Model
	Id int
	SectionId int
	Name string
	Value string
}

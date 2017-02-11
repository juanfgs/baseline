package models


import(
	"testing"
	"github.com/icrowley/fake"
)

func setupSectionTests(){
	MigrationDown()
	MigrationUp()
	stmt, _ := DB.Prepare("INSERT INTO `sections` (name,icon,tagline,content) VALUES (?,?,?,?)")
	stmt.Exec("test", "portfolio", fake.CharactersN(200), fake.CharactersN(600))
}



func TestFindById(t *testing.T){
	setupSectionTests()
	section := NewSection()

	section.FindById(1)
	if section.Name != "test" {
		t.Error("Should find post with ID 1")
	}

}

func TestFindByField(t *testing.T) {
	section := NewSection()

	section.FindByField("name","test")
	if section.Name != "test" {
		t.Error("Should find post with name \"test\"")
	}

}

func TestCreate(t *testing.T){
	section := NewSection()
	section.Name = "test_create"
	section.Icon = "iconfoo"
	section.Tagline = "taglinefoo"
	section.Content = "contentfoo"

	section.Create()
	if section.Id == 0 {
		t.Error("Should have created section")
	}
}


func TestSave(t *testing.T){
	section_proto := NewSection()
	section_proto.Name = "test_create"
	section_proto.Icon = "iconfoo"
	section_proto.Tagline = "taglinefoo"
	section_proto.Content = "contentfoo"

	section_proto.Create()

	section := NewSection()
	
	section.FindById(section_proto.Id)
	section.Name = "updated"
	section.Save()

	section2 := NewSection()
	section2.FindById(section_proto.Id)

	if section2.Name != "updated" {
		t.Error("Should have updated post name")
	}

}	


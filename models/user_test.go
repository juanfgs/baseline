package models

import(
	"testing"
)


func setupUserTests(){
	MigrationDown()
	MigrationUp()
}

func TestValidatePassword(t *testing.T){
	setupUserTests()
	u := NewUser()
	u.SetPassword("foobar")
	
	if !u.ValidatePassword("foobar") {
		t.Error("Password should be a valid bcrypt encrypted string")
	}

	
	if u.ValidatePassword("foobard") {
		t.Error("Password bcrypt hash shoudln't match")
	}
}

func TestSetPassword(t *testing.T){
	u := NewUser()
	u.SetPassword("foobar");

	if !u.ValidatePassword("foobar") {
		t.Error("Password should be a valid bcrypt encrypted string")
	}
	
}


func TestCreateUserFail(t *testing.T){
	u := NewUser()
	u.Name = "pepito"
	u.Active = true
	
	defer func (){
		if r:= recover(); r == nil{
			t.Error("Empty password should produce an error")
		}
	}()

	_ = u.Create()
}
func TestCreateUser(t *testing.T){
	u := NewUser()
	u.Name = "pepito2"
	u.Active = true
	u.SetPassword("foobar")

	id := u.Create()
	u2 := NewUser()
	u2.FindById(id)
	if u2.Name != "pepito2" {
		t.Error("User Should be created")
	}
}


package handlers

import (
	"net/http"
	"github.com/juanfgs/cv/models" 
	"github.com/juanfgs/cv/config"
	"github.com/gorilla/securecookie"
	"math/rand"
	"strconv"
	"log"
)


type AuthHandler struct{
	BaseHandler
}

func (t AuthHandler) SignInForm(w http.ResponseWriter, r *http.Request) {
	t.AddTemplate( "layout.tpl")
	t.AddTemplate( "login.tpl")
	t.Render(w, "")
}


func (t AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	
	username := r.FormValue("Email")
	password := r.FormValue("Password")
	

	user := models.NewUser()
	user.FindByField("name", username)
	
	if user.ValidatePassword(password)  {
		data := make(map[string]string)
		data["session_cookie"] = string( username + strconv.Itoa(rand.Intn(100000000)))
		
		var s = securecookie.New([]byte(config.Values.Secret), nil)
		
		if encoded, err := s.Encode("auth-logged-in", data); err == nil {
			cookie := &http.Cookie{
				Name:  "auth-logged-in",
				Value: encoded,
				Path:  "/",
			}
			
			http.SetCookie(w, cookie)
		}
		
		user.SessionCookie = []byte(data["session_cookie"])
		user.Save()
		http.Redirect(w,r, "/admin", http.StatusTemporaryRedirect)
	}
}



func (t AuthHandler) SignUpForm(w http.ResponseWriter, r *http.Request){
	if config.Values.UserRegistration {
		t.AddTemplate( "layout.tpl")
		t.AddTemplate( "signup.tpl")
		t.Render(w, "")
	} else {
		http.Redirect(w,r,"/auth/sign-in", http.StatusTemporaryRedirect)
	}
}

func (t AuthHandler) SignUp(w http.ResponseWriter, r *http.Request){
	
	username := r.FormValue("Email")
	password := r.FormValue("Password")
	
	user := models.NewUser()
	user.FindByField("name", username)
	log.Println(username)
	if user.Id != 0 {
		http.Redirect(w,r,"/auth/sign-in", http.StatusTemporaryRedirect)
		return 
	}

	user.Name = username
	
	user.Active = true
	user.SetPassword(password)
	user.Create()
	http.Redirect(w,r,"/auth/sign-in", http.StatusTemporaryRedirect)

}





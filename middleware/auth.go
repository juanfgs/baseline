package middleware

import (
	"net/http"
	"github.com/juanfgs/cv/models" 
	"github.com/juanfgs/cv/config"
	"github.com/gorilla/securecookie"
	"log"
)



// Authentication middleware
func Auth(  h http.Handler ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s = securecookie.New([]byte(config.Values.Secret), nil)
		

		
		value := make(map[string]string)
		if cookie, err := r.Cookie("auth-logged-in"); err == nil {
			
			if err = s.Decode("auth-logged-in", cookie.Value, &value); err == nil {
				
				user := models.NewUser()
				user.FindByField("session_cookie",value["session_cookie"])
				if user.Id != 0 {
					
					h.ServeHTTP(w, r) // call original
				}
				
			} else {
				log.Println(err)
			}
			
		}
		
		http.Redirect(w,r, "/auth/sign-in", http.StatusTemporaryRedirect)
	})
}

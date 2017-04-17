package handlers_test


import (
	"testing"
	"github.com/juanfgs/cv/handlers"
	"net/http/httptest"
	"net/http"
	"log"
	"io/ioutil"
)


func TestMain(t *testing.T){
	IsTest = true
}


func TestSignIn(t *testing.T){
	res := visit( handlers.AuthHandler{}.SignIn )
	
	page, _ := ioutil.ReadAll(res.Body)
	
	res.Body.Close()

	// Assert
	if res.StatusCode != 500 {
		t.Errorf("should fail if there is no session data")
	}
}

func visit( handler func(w http.ResponseWriter, r *http.Request) ) *http.Response {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err:= http.Get(ts.URL)
	if err != nil {
		log.Fatal("Error Reaching server")
	}
	return res
}

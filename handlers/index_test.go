package handlers


import (
	"testing"
	"net/http/httptest"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

var IsTest bool

func init(){
	IsTest = true
}


func TestIndex(t *testing.T){
	res := visit( FormHandler{}.Index )
	
	page, _ := ioutil.ReadAll(res.Body)
	
	res.Body.Close()

	i := strings.Index(string(page), "Start Bootstrap")

	if i == -1 {
		t.Errorf("Not found")
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

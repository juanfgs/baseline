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
	ts := httptest.NewServer(http.HandlerFunc(FormHandler{}.Index))
	defer ts.Close()

	res, err:= http.Get(ts.URL)
	if err != nil {
		log.Fatal("Error Reaching server")
	}
	page, err := ioutil.ReadAll(res.Body)
	
	res.Body.Close()

	i := strings.Index(string(page), "Start Bootstrap")

	if i == -1 {
		t.Errorf("Not found")
	} 
}

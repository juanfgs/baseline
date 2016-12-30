package config

import(
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	AppPath string
	Env string
}

var Values Config = Config{}

func init(){
	var configFile *os.File
	var openErr error
	// TODO: Find a way to detect environment other than system variables
	switch Values.Env = os.Getenv("APP_ENV"); Values.Env {
	case "testing":
		configFile, openErr = os.Open("../config/app.testing.json")
	case "development":
		configFile, openErr = os.Open("../config/app.development.json")
	default:
		configFile, openErr = os.Open("config/app.development.json")
	}
	
	if openErr != nil {
		log.Fatal("Opening app.json failed")
	}
	
	decoder := json.NewDecoder(configFile)
	
	parseErr := decoder.Decode(&Values)
	
	if parseErr != nil {
		log.Fatal("Parsing app.json failed")
	}
	
}

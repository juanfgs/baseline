package models

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/juanfgs/react-js-test/config"
)

var DB *sql.DB

type Model struct {
	
}

func init() {
	var dbError error
	DB, dbError = sql.Open(config.Values.Database.Driver, config.Values.Database.Username + ":" + config.Values.Database.Password + "@"+ config.Values.Database.Host +"/"+ config.Values.Database.Name )
	
	if dbError != nil {
		log.Fatal("Error : %s", dbError)
	}
} 


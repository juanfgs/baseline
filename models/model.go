package models

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/juanfgs/cv/config"
)

var DB *sql.DB

type Model struct {
	Written bool
	lastQuery string
	tableName string
	indexName string
}

func init() {
	var dbError error
	DB, dbError = sql.Open(config.Values.Database.Driver, config.Values.Database.Username + ":" + config.Values.Database.Password + "@tcp("+ config.Values.Database.Host +":"+ config.Values.Database.Port +")/"+ config.Values.Database.Name + "?parseTime=true")
	
	if dbError != nil {
		log.Fatal("Error : ", dbError)
	}
} 

func (this Model) FindById( id int64, fields ...interface{})  {
	err:= DB.QueryRow("SELECT * FROM `"+ this.tableName +"` WHERE `"+this.indexName+"` = ?", id ).Scan(fields...)
	if err != nil {
		log.Println(err)
	}

}

func (this Model) FindByField( field string, value string, fields ...interface{})  {
	
	err:= DB.QueryRow("SELECT * FROM `"+ this.tableName +"` WHERE `"+ field +"` = ?", value ).Scan(fields...)
	if err != nil {
		log.Println(err)
	}

}


func (this *Model) SetIndex(idx string){
	this.indexName = idx
}

func (this *Model) SetTable(table string){
	this.tableName = table
}


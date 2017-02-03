package models

import(
	"github.com/mattes/migrate/migrate"
	_ "github.com/mattes/migrate/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/juanfgs/cv/config"
)

var DB *sql.DB

var driverString string = config.Values.Database.Username + ":" + config.Values.Database.Password + "@tcp("+ config.Values.Database.Host +":"+ config.Values.Database.Port +")/"+ config.Values.Database.Name + "?parseTime=true"

type Model struct {
	Written bool
	lastQuery string
	tableName string
	indexName string
}

func init() {
	var dbError error
	DB, dbError = sql.Open(config.Values.Database.Driver, driverString)
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

func (this *Model) Delete(id int64) {
	_, err:= DB.Query("DELETE FROM `"+ this.tableName +"` WHERE `"+this.indexName+"` = ?", id )
	if err != nil {
		log.Println(err)
	}
}

func (this Model) All() *sql.Rows{
	rows, err:= DB.Query("SELECT * FROM `"+ this.tableName +"`" )
	if err != nil {
		log.Println(err)
	}

	return rows
}

func (this Model) AllWhere(where string) *sql.Rows{
	rows, err:= DB.Query("SELECT * FROM `"+ this.tableName +"` WHERE " + where )
	if err != nil {
		log.Println(err)
	}

	return rows
}



func (this *Model) SetIndex(idx string){
	this.indexName = idx
}

func (this *Model) SetTable(table string){
	this.tableName = table
}

func MigrationUp(){
	errors, ok := migrate.UpSync( config.Values.Database.Driver +"://"+ driverString, "../migrations/")
	if !ok {
		panic(errors)
	}
}

func MigrationDown(){
	errors, ok := migrate.DownSync(config.Values.Database.Driver + "://" + driverString, "../migrations/")
	if !ok {
		log.Println(errors)
		panic(errors)
	}
}

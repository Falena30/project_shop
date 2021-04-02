package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbmap = initDB()

//Conn merupakan fungsi untk menghubunkan ke dalam database
func initDB() *gorp.DbMap {
	dataSourceNameTamp := ExtrackDBYaml("db", ".")
	dataSourceName := fmt.Sprintf("%s:@%s(%s:%s)/%s", dataSourceNameTamp.Username, dataSourceNameTamp.Addr, dataSourceNameTamp.HOST, dataSourceNameTamp.PORT, dataSourceNameTamp.NameDB)
	//buka sqlnya
	db, err := sql.Open("mysql", dataSourceName)
	checkErr(err, "sql.open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

//checkErr untuk mengetahui apakah sql open errror atau tidak
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

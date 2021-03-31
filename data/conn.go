package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var dbmap = initDB()

//ExtrackDBYaml untuk extract data dari env
func ExtrackDBYaml(nameFileYML string, LocationFile string) *DBYaml {
	var db DBYaml
	//beritahu viper nama yamlnya
	viper.SetConfigName(nameFileYML)
	//beritahu lokasi filenya
	viper.AddConfigPath(LocationFile)
	//beri izin viper
	viper.AutomaticEnv()
	//beritahu viper tipe filenya
	viper.SetConfigType("yml")
	//cek apakah ada error atau tidak
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	//set undefined variabelnya
	viper.SetDefault("DB.HOST", "127.0.0.1")
	//ambil nilai vipernya dan tampung
	DBHost, ok := viper.Get("DB.HOST").(string)
	DBport, ok := viper.Get("DB.PORT").(string)
	DBUsername, ok := viper.Get("DB.USERNAME").(string)
	DBPass, ok := viper.Get("DB.PASWORD").(string)
	DBNameDB, ok := viper.Get("DB.DBNAME").(string)
	DBAddress, ok := viper.Get("DB.ADDRES").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	//masukkan kedalam struct
	db = DBYaml{
		HOST:     DBHost,
		PORT:     DBport,
		Username: DBUsername,
		Password: DBPass,
		NameDB:   DBNameDB,
		Addr:     DBAddress,
	}
	//kembalikan nilainya
	return &db
}

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

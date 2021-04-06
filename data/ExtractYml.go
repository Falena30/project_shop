package data

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

//ExtrackDBYaml untuk extract data dari env
func ExtrackDBYaml(nameFileYML string, LocationFile string) *DBYaml {
	var db DBYaml
	TemplateExtractYml(nameFileYML, LocationFile)
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

func ExtractQueryYml(name string, location string) *Query {
	var Que Query
	TemplateExtractYml(name, location)
	//set dauflt
	viper.SetDefault("QUERY.UNIQValueDefalt", "DefaultValue")
	//tampung nilainya
	Create, ok := viper.Get("QUERY.CREATE").(string)
	Read, ok := viper.Get("QUERY.READ").(string)
	Update, ok := viper.Get("QUERY.UPDATE").(string)
	Delete, ok := viper.Get("QUERY.DELETE").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	//pmasukkan data kedalam struct QUERY
	Que = Query{
		CREATE: Create,
		READ:   Read,
		UPDATE: Update,
		DELETE: Delete,
	}
	//kembalikan nilainya
	return &Que
}

func ExtractQueryUserYml(name string, location string) *QueryUser {
	var Que QueryUser
	TemplateExtractYml(name, location)
	//set dauflt
	viper.SetDefault("USER.UNIQValueDefalt", "ThisIsValueDefault")
	//tampung nilainya
	Create, ok := viper.Get("USER.CREATE").(string)
	Read, ok := viper.Get("USER.READ").(string)
	PartialRead, ok := viper.Get("USER.PARTIALREAD").(string)
	Update, ok := viper.Get("USER.UPDATE").(string)
	Delete, ok := viper.Get("USER.DELETE").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	//pmasukkan data kedalam struct QUERY
	Que = QueryUser{
		CREATE:      Create,
		READ:        Read,
		UPDATE:      Update,
		DELETE:      Delete,
		PARTIALREAD: PartialRead,
	}
	//kembalikan nilainya
	return &Que
}

func TemplateExtractYml(name string, location string) {
	//kasih tahu nama filenya
	viper.SetConfigName(name)
	//beritahu lokasinya
	viper.AddConfigPath(location)
	//beri izin viper
	viper.AutomaticEnv()
	//beritahu formtnya
	viper.SetConfigType("yml")
	//check apakah terjadi error pada saat eksekusi
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func ExtractQueryUserDetailYml(name string, location string) *QueryUser {
	var Que QueryUser
	TemplateExtractYml(name, location)
	//set dauflt
	viper.SetDefault("USERDETAIL.UNIQValueDefalt", "DefaultValue")
	//tampung nilainya
	Create, ok := viper.Get("USERDETAIL.CREATE").(string)
	Read, ok := viper.Get("USERDETAIL.READ").(string)
	Update, ok := viper.Get("USERDETAIL.UPDATE").(string)
	Delete, ok := viper.Get("USERDETAIL.DELETE").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	//pmasukkan data kedalam struct QUERY
	Que = QueryUser{
		CREATE: Create,
		READ:   Read,
		UPDATE: Update,
		DELETE: Delete,
	}
	//kembalikan nilainya
	return &Que
}

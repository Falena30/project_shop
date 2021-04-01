package data

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ExtractQueryYml(name string, location string) *Query {
	var Que Query
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

func GetDataBarang() *[]DataBarang {
	Query := ExtractQueryYml("query", ".")
	var data []DataBarang
	_, err := dbmap.Select(&data, Query.READ)

	if err != nil {
		fmt.Println(err.Error())
	}
	return &data

}

//GetBarangById merupakan fungsi untuk mencari barang berdasarkan idnya
func GetBarangById(Id int) (*DataBarang, error) {
	barang := GetDataBarang()
	for _, a := range *barang {
		return &a, nil
	}
	return nil, errors.New("Barang Tidak ada")
}

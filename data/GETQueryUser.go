package data

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ExtractQueryUserYml(name string, location string) *QueryUser {
	var Que QueryUser
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
	viper.SetDefault("QUERYUSER.UNIQValueDefalt", "DefaultValue")
	//tampung nilainya
	Create, ok := viper.Get("QUERYUSER.CREATE").(string)
	Read, ok := viper.Get("QUERYUSER.READ").(string)
	PartialRead, ok := viper.Get("QUERYUSER.PARTIALREAD").(string)
	Update, ok := viper.Get("QUERYUSER.UPDATE").(string)
	Delete, ok := viper.Get("QUERYUSER.DELETE").(string)
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

//GetAllUser adalah fungsi untuk GET semua user dari database
func GetAllUser() *[]UserData {
	//ambil dulu semua data dari queryUser
	QueryRead := ExtractQueryYml("queryUser", ".")
	//buatlah penampung
	var User []UserData
	//buat SQLnya
	_, err := dbmap.Select(&User, QueryRead.READ)
	if err != nil {
		fmt.Println(err.Error())
	}
	//kembalikan nilainya
	return &User
}

//GetUserByID adalah fungsi untuk GET User yang dicari
func GetUserByID(id int) (*UserData, error) {
	//tampung semua nilai user
	User := GetAllUser()
	//caari satu2
	for _, a := range *User {
		return &a, nil
	}
	return nil, errors.New("User tidak ada")
}

func GetUserLogin(user string, pass string) (*UserData, error) {
	//panggil extract value dari queryuser
	Query := ExtractQueryUserYml("queryUser", ".")
	//buatlah variabel untuk menampung nilainya nanti
	var userTamp UserData
	_, err := dbmap.Select(&userTamp, Query.PARTIALREAD, user, pass)
	if err != nil {
		return nil, err
	}
	//kembalikan nilainya
	return &userTamp, nil
}

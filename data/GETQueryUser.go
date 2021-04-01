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

//GetAllUser adalah fungsi untuk GET semua user dari database
func GetAllUser() *[]UserData {
	//ambil dulu semua data dari queryUser
	QueryRead := ExtractQueryUserYml("User", ".")
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

func GetUserLogin(user string) *UserData {
	//panggil extract value dari queryuser
	//Query := ExtractQueryUserYml("queryUser", ".")
	//buatlah variabel untuk menampung nilainya nanti
	userAll := GetAllUser()
	/*
		var userTamp UserData
		_, err := dbmap.Select(&userTamp, "SELECT * FROM User WHERE `U_Username` = ? LIMIT 1", user)
		if err != nil {
			return nil, err
		}
		//kembalikan nilainya
		return &userTamp, nil
	*/
	fmt.Println(userAll)
	for _, tamp := range *userAll {
		if tamp.Username == user {
			return &tamp
		}
	}
	return nil
}

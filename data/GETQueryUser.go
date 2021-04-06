package data

import (
	"errors"
	"fmt"
)

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
	userAll := GetAllUser()
	for _, tamp := range *userAll {
		if tamp.Username == user {
			return &tamp
		}
	}
	return nil
}

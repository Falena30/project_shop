package data

import (
	"errors"
	"fmt"
)

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
		if a.ID == Id {
			return &a, nil
		}
	}
	return nil, errors.New("Barang Tidak ada")
}

func GetDataBarangByIDUser(ID int) (*[]DataBarang, error) {
	var dataBarang []DataBarang
	_, err := dbmap.Select(&dataBarang, "SELECT Daftar_Barang.ID_Barang, Daftar_Barang.Nama_Barang,Daftar_Barang.Harga_Barang, User.U_Username FROM Daftar_Barang INNER JOIN User ON Daftar_Barang.ID_User = User.ID_User WHERE User.ID_User = ?", ID)
	if err != nil {
		return nil, err
	}
	return &dataBarang, nil
}

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
		return &a, nil
	}
	return nil, errors.New("Barang Tidak ada")
}

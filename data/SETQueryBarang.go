package data

func SetDataBarang(nama_barang string, harga_barang int, idUser int) error {
	QueryInsertBarang := ExtractQueryYml("query", ".")
	_, err := dbmap.Exec(QueryInsertBarang.CREATE, nil, nama_barang, harga_barang, idUser)
	if err != nil {
		return err
	}
	return nil
}

func PutDataBarang(ID_Barang int, Nama_Barang string, Harga_Barang int) error {
	QueryPut := ExtractQueryYml("query", ".")
	_, err := dbmap.Exec(QueryPut.UPDATE, Nama_Barang, Harga_Barang, ID_Barang)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDataBarang(ID_Barang int) error {
	QueryDelete := ExtractQueryYml("query", ".")
	Result, err := dbmap.Exec(QueryDelete.DELETE, ID_Barang)
	if Result == nil {
		return err
	}
	return nil
}

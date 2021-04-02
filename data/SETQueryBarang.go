package data

func SetDataBarang(nama_barang string, harga_barang int, idUser int) error {
	QueryInsertBarang := ExtractQueryYml("query", ".")
	_, err := dbmap.Exec(QueryInsertBarang.CREATE, nil, nama_barang, harga_barang, idUser)
	if err != nil {
		return err
	}
	return nil
}

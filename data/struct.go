package data

type DataBarang struct {
	ID    int    `db:"ID_Barang" json:"id"`
	Nama  string `db:"Nama_Barang" json:"nama"`
	Harga int    `db:"Harga_Barang" json:"harga"`
	User  string `db:"U_Username" json:"user"`
}

type UserData struct {
	ID       int    `db:"ID_User" json:"id_u"`
	Username string `db:"U_Username" json:"username"`
	Password string `db:"U_Password" json:"password"`
}

type UserDataDetail struct {
	ID       int    `db:"ID_User" json:"id_u"`
	Username string `db:"U_Username" json:"username"`
	Fname    string `db:"First_Name" json:"first_name"`
	LName    string `db:"Last_Name" json:"last_name"`
	Bio      string `db:"Biodata" json:"biodata"`
}

type DBYaml struct {
	HOST     string
	PORT     string
	Username string
	Password string
	NameDB   string
	Addr     string
}

type Query struct {
	CREATE string
	READ   string
	UPDATE string
	DELETE string
}

type UserDetail struct {
	QueryBasic Query
}
type QueryUser struct {
	CREATE      string
	READ        string
	UPDATE      string
	DELETE      string
	PARTIALREAD string
}

type RedisData struct {
	NETWORK  string
	ADDRESS  string
	PASSWORD string
	KEYPAIRS string
}

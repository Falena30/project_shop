package data

func CreateUser(username string, password string) error {
	QueryInsert := ExtractQueryUserYml("User", ".")
	_, err := dbmap.Exec(QueryInsert.CREATE, nil, username, password)
	if err != nil {
		return err
	}
	return nil
}

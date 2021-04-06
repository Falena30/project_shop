package data

import "fmt"

func GETAllUserDetail() (*[]UserDataDetail, error) {
	var AllUserDetail []UserDataDetail
	QueryGetAll := ExtractQueryUserDetailYml("userDetail", ".")
	_, err := dbmap.Select(&AllUserDetail, QueryGetAll.READ)
	if err != nil {
		return nil, err
	}
	return &AllUserDetail, nil
}

func GetUserDetail(username string) (*UserDataDetail, error) {
	var GetUser UserDataDetail
	AllUserDetail, err := GETAllUserDetail()
	fmt.Println(AllUserDetail)
	if err != nil {
		return nil, err
	}
	for _, a := range *AllUserDetail {
		if a.Username == username {
			GetUser = a
		} else {
			return nil, nil
		}
	}
	return &GetUser, nil
}

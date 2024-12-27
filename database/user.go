package database

import "errors"

func UserExist(username string) bool {
	var i int64
	DB.Model(&UserTable{}).Where("username = ?", username).Count(&i)
	return i != 0
}

func UserNew(username string, password string) error {
	if UserExist(username) {
		return errors.New("用户已存在")
	}
	var i UserTable
	i.Username = username
	i.Password = md5Generate(password)
	DB.Create(&i)
	return nil
}

func UserLogin(username string, password string) bool {
	if !UserExist(username) {
		return false
	}
	var i UserTable
	DB.Model(&UserTable{}).Where("username = ?", username).First(&i)
	return i.Password == md5Generate(password)
}

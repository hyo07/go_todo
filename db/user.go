package db

import (
	"go_todo/helper"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

//DB初期化
func UserInit() {
	db := dbOpen()
	db.AutoMigrate(&User{})
	defer db.Close()
}

//DB追加
func UserInsert(username string, hash string) {
	db := dbOpen()
	db.Create(&User{Username: username, Password: hash})
	defer db.Close()
}

//既にusername	が登録されていないか確認
func UserAlredy(username string) string {
	db := dbOpen()
	var user User
	db.Where("username = ?", username).Find(&user)
	db.Close()
	name := user.Username
	return name
}

//usernameとpasswordが正しいか検証
func UserCheck(username string, password string) error {
	db := dbOpen()
	var user User
	db.Where("username = ?", username).Find(&user)
	db.Close()

	err := helper.PasswordValid(user.Password, password)
	return err
}

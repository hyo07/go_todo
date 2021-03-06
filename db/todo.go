package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

//DBに接続
func dbOpen() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbOpen）")
	}
	return db
}

//DB初期化
func DbInit() {
	db := dbOpen()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//DB追加
func DbInsert(text string, status string) {
	db := dbOpen()
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func DbUpdate(id int, text string, status string) {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func DbDelete(id int) {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func DbGetAll() []Todo {
	db := dbOpen()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func DbGetOne(id int) Todo {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

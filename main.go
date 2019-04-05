package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
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
func dbInit() {
	db := dbOpen()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//DB追加
func dbInsert(text string, status string) {
	db := dbOpen()
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func dbUpdate(id int, text string, status string) {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func dbDelete(id int) {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func dbGetAll() []Todo {
	db := dbOpen()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func dbGetOne(id int) Todo {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	//ユーザー周り
	user := router.Group("/user")
	{
		user.POST("/signup", func(ctx *gin.Context) {
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")
			passwordConf := ctx.PostForm("passwordconf")
			if password != passwordConf {
				println("パスワードが一致していません")
				ctx.Redirect(302, "/")
				return
			}
			println("username: " + username)
			println("password: " + password)
			println("passwordConf: " + passwordConf)

			ctx.Redirect(302, "/")
		})
		user.POST("/login", func(ctx *gin.Context) {
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")
			println("username: " + username)
			println("password: " + password)

			ctx.Redirect(302, "/")
		})
	}

	//Index
	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbDelete(id)
		ctx.Redirect(302, "/")

	})

	//login
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})

	//signup
	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	})

	router.Run()
}

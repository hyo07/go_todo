package main

import (
	"go_todo/db"
	"go_todo/route"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	//dbInit()
	db.DbInit()
	db.UserInit()

	//Index
	router.GET("/", route.Index)
	//Create
	router.POST("/new", route.Create)
	//Detail
	router.GET("/detail/:id", route.Detail)
	//Update
	router.POST("/update/:id", route.Update)
	//削除確認
	router.GET("/delete_check/:id", route.DelConf)
	//Delete
	router.POST("/delete/:id", route.Delete)
	//login
	router.GET("/login", route.Login)
	//signup
	router.GET("/signup", route.Signup)

	//ユーザー周り
	user := router.Group("/user")
	{
		//singup
		user.POST("/signup", route.UserSignup)
		//login
		user.POST("/login", route.UserLogin)
	}

	router.Run()

}

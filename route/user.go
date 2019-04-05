package route

import (
	"go_todo/helper"

	"github.com/gin-gonic/gin"
)

//ログイン
func UserSignup(ctx *gin.Context) {
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
}

//アカウント作成
func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	println("username: " + username)
	println("password: " + password)

	hash, err := helper.PasswordHash(password)
	if err != nil {
		panic("hash err")
	}
	println("hashed PW >>>>", hash)
	err2 := helper.PasswordValid(hash, password)
	if err2 != nil {
		println("ERRRRER")
		println(err2)
	}

	ctx.Redirect(302, "/")
}

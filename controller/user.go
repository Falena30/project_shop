package controller

import (
	"fmt"
	"net/http"
	"project/shop/data"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	Username := c.PostForm("loginUsername")
	Password := c.PostForm("loginPass")
	User := data.GetUserLogin(Username)
	if User == nil {
		Render(c, gin.H{
			"tittle": "login",
			"status": "user tidak ada",
		}, "login.html")
	} else {
		if User.Password == Password {
			session.Set("id", User.ID)
			session.Set("username", User.Username)
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/dasbord/")
		} else {
			Render(c, gin.H{
				"tittle": "login",
				"status": "Password Salah",
			}, "login.html")
		}

	}
}

func RenderLogin(c *gin.Context) {
	Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func RenderDasbord(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	strUsername := fmt.Sprintf("%v", username)
	fmt.Println(strUsername)
	Render(c, gin.H{
		"title":   "Dasbord",
		"payload": strUsername,
	}, "dasbord.html")
}

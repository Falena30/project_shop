package controller

import (
	"net/http"
	"project/shop/data"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	Username := c.PostForm("loginUsername")
	Password := c.PostForm("loginPass")

	User, err := data.GetUserLogin(Username, Password)
	_ = User // harusnya nanti masukkan ke auth
	if err != nil {
		Render(c, gin.H{
			"tittle": "login",
			"status": err.Error(),
		}, "login.html")
	} else {
		session.Set("id", User.ID)
		session.Set("username", User.Username)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/dasbord/")
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
	Render(c, gin.H{
		"title":    "Dasbord",
		"username": username,
	}, "dasbord.html")
}

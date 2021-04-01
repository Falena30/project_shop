package controller

import (
	"net/http"
	"project/shop/data"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	Username := c.PostForm("loginUsername")
	Password := c.PostForm("loginPass")

	User, err := data.GetUserLogin(Username, Password)
	_ = User // harusnya nanti masukkan ke auth
	if err != nil {
		render(c, gin.H{
			"tittle": "login",
			"status": err.Error(),
		}, "login.html")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/dasbord/")
	}
}

func RenderLogin(c *gin.Context) {
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func RenderDasbord(c *gin.Context) {
	render(c, gin.H{
		"title": "Dasbord",
	}, "dasbord.html")
}

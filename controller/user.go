package controller

import (
	"fmt"
	"net/http"
	"project/shop/data"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		errDekrib := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(Password))
		if errDekrib == nil {
			//middleware.SetSession(User.ID,User.Username)
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

func Register(c *gin.Context) {
	session := sessions.Default(c)
	Username := c.PostForm("regisUsername")
	Password := c.PostForm("regisPass")
	HPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	err = data.CreateUser(Username, string(HPassword))
	if err != nil {
		Render(c, gin.H{
			"tittle": "register",
			"status": err,
		}, "register.html")
	} else {
		session.Set("id", Username)
		session.Set("username", Password)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/dasbord/")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	//coba clear dulu
	sessionID := session.Get("id")
	sessionUsername := session.Get("username")
	if sessionID != nil || sessionUsername != nil {
		session.Clear()
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func RenderLogin(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("id")
	if sessionID != nil {
		c.Redirect(http.StatusMovedPermanently, "/dasbord/")
	} else {
		Render(c, gin.H{
			"title": "Login",
		}, "login.html")

	}
}
func RenderRegister(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("id")
	if sessionID != nil {
		c.Redirect(http.StatusMovedPermanently, "/dasbord/")
	} else {
		Render(c, gin.H{
			"title": "Register",
		}, "register.html")
	}
}

func RenderDasbord(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	strUsername := fmt.Sprintf("%v", username)
	Render(c, gin.H{
		"title":   "Dasbord",
		"payload": strUsername,
	}, "dasbord.html")
}

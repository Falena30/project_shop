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
	//ambil semua user
	userAll := data.GetAllUser()
	session := sessions.Default(c)
	Username := c.PostForm("regisUsername")
	Password := c.PostForm("regisPass")
	HPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	for _, a := range *userAll {
		if a.Username == Username {
			Render(c, gin.H{
				"tittle": "register",
				"status": "username sudah ada",
			}, "register.html")
		} else {
			err = data.CreateUser(Username, string(HPassword))
		}
	}
	if err != nil {
		Render(c, gin.H{
			"tittle": "register",
			"status": err,
		}, "register.html")
	} else {
		session.Set("username", Username)
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
	//ambil nilai sessionya
	username := session.Get("username")
	sessionID := session.Get("id")
	//jika nilai sessionIDnya kosong
	//tentukan beri nilainya dengan mengambil id dari database
	if sessionID == nil {
		AllUser := data.GetAllUser()
		for _, a := range *AllUser {
			if a.Username == a.Username {
				session.Set("id", a.ID)
				sessionID = session.Get("id")
			}
		}
	}
	//tampung username kedalam string
	strUsername := fmt.Sprintf("%v", username)
	//cari barang yang di posting oleh user
	BarangByUser, err := data.GetDataBarangByIDUser(sessionID.(int))
	if err != nil {
		Render(c, gin.H{
			"title":    "Dasbord",
			"username": strUsername,
			"payload":  nil,
		}, "dasbord.html")
	} else {
		Render(c, gin.H{
			"title":    "Dasbord",
			"username": strUsername,
			"payload":  BarangByUser,
		}, "dasbord.html")
	}

}

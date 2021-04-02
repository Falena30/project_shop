package middleware

import (
	"project/shop/controller"
	"project/shop/data"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MiddleInputBarang(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("id")
	nama_barang := c.PostForm("namaBarang")
	harga_barang, err := strconv.Atoi(c.PostForm("hargaBarang"))
	if err != nil {
		panic(err)
	}
	err = data.SetDataBarang(nama_barang, harga_barang, sessionID.(int))
	if err != nil {
		controller.Render(c, gin.H{
			"tittle": "Input Barang",
			"status": err,
		}, "input.html")
	} else {
		controller.Render(c, gin.H{
			"tittle": "Input Barang",
			"status": "barang berhasil di input",
		}, "input.html")
	}
}

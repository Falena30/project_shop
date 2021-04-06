package middleware

import (
	"net/http"
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

func PutDataBarang(c *gin.Context) {
	namaBarang := c.PostForm("namaBarangEdit")
	hargaBarang, _ := strconv.Atoi(c.PostForm("hargaBarangEdit"))
	IDBarang, _ := strconv.Atoi(c.Param("barang_id"))

	if err := data.PutDataBarang(IDBarang, namaBarang, hargaBarang); err != nil {
		controller.Render(c, gin.H{
			"tittle": "Edit Barang",
			"status": err.Error(),
		}, "edit.html")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/dasbord/")
		c.Abort()
	}
}

func DeleteBarangOK() gin.HandlerFunc {
	return func(c *gin.Context) {
		IDBarang, _ := strconv.Atoi(c.Param("barang_id"))
		if err := data.DeleteDataBarang(IDBarang); err != nil {
			controller.Render(c, gin.H{
				"tittle": "Status Page",
				"status": err.Error(),
			}, "status.html")
		} else {
			controller.Render(c, gin.H{
				"tittle": "Status Page",
				"status": "Delete berehasil",
			}, "status.html")
		}
	}
}

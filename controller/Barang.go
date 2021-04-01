package controller

import (
	"net/http"
	"project/shop/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {

	//buatlah perpindahan apabla yang diinginkan JSON,XML, atau HTML
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		//respond with json
		c.JSON(http.StatusOK, data["payload"])
	case "application/XML":
		//respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		//jika responsenya HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

func ShowIndexPage(c *gin.Context) {
	barang := data.GetDataBarang()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": barang,
	}, "index.html")

}

func GetBarang(c *gin.Context) {
	//format param barang id menjadi int
	if barangID, err := strconv.Atoi(c.Param("barang_id")); err == nil {
		//dapatkan barangnya berdasarkan idnya
		if barang, err := data.GetBarangById(barangID); err == nil {
			//render barang ke hmtl
			render(c, gin.H{
				"title":   barang.Nama,
				"payload": barang,
			}, "barang.html")
		}
	}
}

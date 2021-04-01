package main

import (
	"project/shop/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templete/*")
	router.GET("/", controller.ShowIndexPage)
	router.GET("/barang/view/:barang_id", controller.GetBarang)
	login := router.Group("/login")
	{
		login.GET("/", controller.RenderLogin)
		login.POST("/", controller.Login)
	}
	dasbord := router.Group("/dasbord")
	{
		dasbord.GET("/", controller.RenderDasbord)
	}
	router.Run()
}

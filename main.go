package main

import (
	"project/shop/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templete/*")
	router.GET("/", controller.ShowIndexPage)
	router.Run()
}

package main

import (
	"project/shop/controller"
	"project/shop/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	redisTamp := middleware.ExtractRedisYml("redis", ".")
	store, _ := redis.NewStore(10, redisTamp.NETWORK, redisTamp.ADDRESS, redisTamp.PASSWORD, []byte(redisTamp.KEYPAIRS))
	router.Use(sessions.Sessions("mysession", store))
	/*
		store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
		//storing session
		router.Use(sessions.Sessions("mysession", store))
	*/
	router.LoadHTMLGlob("templete/*")
	router.GET("/", controller.ShowIndexPage)
	router.GET("/barang/view/:barang_id", controller.GetBarang)
	login := router.Group("/login")
	{
		login.GET("/", controller.RenderLogin)
		login.POST("/", controller.Login)
	}
	dasbord := router.Group("/dasbord")
	dasbord.Use(middleware.AuthUser())
	{
		dasbord.GET("/", controller.RenderDasbord)
	}
	router.Run()
}

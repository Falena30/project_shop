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
	router.LoadHTMLGlob("templete/*")
	//index
	router.GET("/", controller.ShowIndexPage)
	router.GET("/barang/view/:barang_id", controller.GetBarang)
	//login
	login := router.Group("/login")
	{
		login.GET("/", controller.RenderLogin)
		login.POST("/", controller.Login)
	}
	//logout
	logout := router.Group("/logout")
	{
		logout.POST("/", controller.Logout)
	}
	//register
	register := router.Group("/register")
	{
		register.GET("/", controller.RenderRegister)
		register.POST("/", controller.Register)
	}
	//dasbord
	dasbord := router.Group("/dasbord")
	dasbord.Use(middleware.AuthUser())
	{
		dasbord.GET("/", controller.RenderDasbord)
		dasbord.GET("/input", controller.GetInputBarang)
		dasbord.POST("/input", middleware.MiddleInputBarang)
		dasbord.GET("/barang/view/:barang_id", controller.GetDashbordViewBarang)
		dasbord.GET("/barang/view/:barang_id/edit/", controller.RenderPutBarang)
		dasbord.POST("/barang/view/:barang_id/edit/", middleware.PutDataBarang)
		dasbord.DELETE("/barang/view/:barang_id/delete", middleware.DeleteBarangOK())
	}
	router.Run()
}

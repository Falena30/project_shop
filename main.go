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
		barang := dasbord.Group("/barang")
		{
			barang.GET("/view/:barang_id", controller.GetDashbordViewBarang)
			barang.GET("/view/:barang_id/edit/", controller.RenderPutBarang)
			barang.POST("/view/:barang_id/edit/", middleware.PutDataBarang)
			barang.DELETE("/view/:barang_id/delete", middleware.DeleteBarangOK())
		}
		userRouter := dasbord.Group("/user")
		{
			userRouter.GET("/", controller.RenderUserDetail)
			userRouter.GET("/add/", controller.RenderAddUserDetail)
			userRouter.POST("/add/")
		}
	}
	router.Run()
}

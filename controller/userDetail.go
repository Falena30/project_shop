package controller

import (
	"net/http"
	"project/shop/data"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RenderUserDetail(c *gin.Context) {
	session := sessions.Default(c)
	sessionUsername := session.Get("username")
	UserDetail, err := data.GetUserDetail(sessionUsername.(string))
	if err != nil {
		panic(err)
	} else {
		if UserDetail != nil {
			Render(c, gin.H{
				"tittle":  "User Page",
				"payload": UserDetail,
			}, "user.html")
		} else {
			//redirect ke edit user
			c.Redirect(http.StatusMovedPermanently, "/dasbord/user/add")
		}
	}
}

func RenderAddUserDetail(c *gin.Context) {
	Render(c, gin.H{
		"tittle": "Add Detail User",
	}, "userAddDetail.html")
}

package routers

import (
	"demo_cms_1/auth"
	"demo_cms_1/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("./templates/**/*")
	indexGroup := r.Group("/")
	{
		indexGroup.GET("/", auth.AuthSession(), controller.IndexGet)
	}
	welcomeGroup := r.Group("welcome")
	{
		welcomeGroup.GET("/", auth.AuthSession(), controller.Welcome)
	}
	loginGroup := r.Group("login")
	{
		loginGroup.GET("/", controller.LoginGetController)
		loginGroup.POST("post", controller.LoginPostController)
	}
	adminGroup := r.Group("admin")
	{
		adminGroup.GET("/", auth.AuthSession(), controller.AdminGetAdmin)
		adminGroup.GET("adminAdd", auth.AuthSession(), controller.AdminGetAdminAdd)
		adminGroup.GET("adminEdit/:id", auth.AuthSession(), controller.AdminGetAdminEdit)
		adminGroup.POST("post", auth.AuthSession(), controller.AdminPostAdmin)
		adminGroup.DELETE("delete/:id", auth.AuthSession(), controller.AdminDeleteAdmin)
		adminGroup.PUT("put", auth.AuthSession(), controller.AdminPutAdmin)
		adminGroup.PUT("toDown", auth.AuthSession(), controller.AdminPutAdminPutAdminToDown)
		adminGroup.POST("logout", auth.AuthSession(), controller.AdminPutAdminPutAdminLogout)

	}
	return r
}

package router

import (
	"net/http"
	"work/controller"
	"work/middleware"
	"work/util"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CorsMiddleware())
	api := r.Group("/api")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		api.GET("/info", middleware.AuthMiddleware(), controller.UserInfo)
		api.POST("/upload", util.Upload)
	}

	r.NoRoute(func(context *gin.Context) { // 这里只是演示，不要在生产环境中直接返回HTML代码
		context.String(http.StatusNotFound, "404 Page Not Found")
	})
	return r
}

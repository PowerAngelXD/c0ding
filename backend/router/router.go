package router

import "github.com/gin-gonic/gin"

func ApiRouter(r *gin.Engine) {
	user := r.Group("/api/user")
	{
		user.POST("/register")
		user.POST("/login")
	}
}

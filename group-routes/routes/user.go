package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "All Users")
	})

	users.POST("/add", func(context *gin.Context) {
		context.JSON(http.StatusCreated, "Users Created")
	})

	users.GET("/comments", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Users Comments")
	})

	users.GET("status", func(context *gin.Context) {
		context.JSON(http.StatusOK, "User Status is ok")
	})
}

func addUserApiV1(rg *gin.RouterGroup) {
	v1 := rg.Group("/", authMiddleWare())

	v1.GET("/users", func(context *gin.Context) {
		context.JSON(http.StatusOK, "List of V1 Users")
	})

	v1.POST("users/add", func(context *gin.Context) {
		context.JSON(http.StatusOK, "V1 Users Added")
	})
}

func addUserApiV2(rg *gin.RouterGroup) {
	v2 := rg.Group("/", authMiddleWare())

	v2.GET("/users", func(context *gin.Context) {
		context.JSON(http.StatusOK, "List of V2 Users")
	})

	v2.POST("users/add", func(context *gin.Context) {
		context.JSON(http.StatusOK, "V2 Users Added")
	})
}

func authMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

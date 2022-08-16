package routes

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {
	getRoutes()
	router.Run("localhost:8081")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)

	v2 := router.Group("/v2")
	addUserRoutes(v2)

	apiv1 := router.Group("/api/v1")
	addUserApiV1(apiv1)

	apiv2 := router.Group("/api/v2")
	addUserApiV2(apiv2)
}

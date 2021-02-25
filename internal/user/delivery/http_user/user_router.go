package http_user

import (
	"go-clean-arch/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

// MakeUserRouters url handlers
func MakeUserRouters(route *gin.RouterGroup, service usecase.UseCaseUser) {
	route.POST("/user", createUser(service))
	route.GET("/user/:id", findUser(service))
	route.PUT("/user/:id", updateUser(service))
	route.DELETE("/user/:id", destroyUser(service))
}

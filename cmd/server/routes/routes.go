package routes

import (
	"errors"
	"fmt"
	"go-clean-arch/cmd/server/middleware"
	"go-clean-arch/internal/user/delivery/http_user"
	"go-clean-arch/internal/user/repository"
	"go-clean-arch/internal/user/usecase"
	"go-clean-arch/pkg/http_errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// MakeRouters add routes
func MakeRouters(postgresDB *sqlx.DB) *gin.Engine {

	app := gin.Default()

	app.Use(middleware.CORS())
	app.Use(middleware.HandleErrors())

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	userRepository := repository.NewUserRepository(postgresDB)
	userService := usecase.NewService(userRepository)

	api := app.Group("api")
	{
		http_user.MakeUserRouters(api, userService)
	}

	app.NoRoute(func(c *gin.Context) {
		panic(http_errors.NotFoundError(errors.New(fmt.Sprintf("Route '%s' was not found", c.Request.URL.Path))))
	})

	return app

}

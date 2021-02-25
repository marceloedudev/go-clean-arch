package http_user

import (
	"errors"
	"go-clean-arch/internal/user/usecase"
	"go-clean-arch/pkg/http_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createUser(us usecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			panic(err)
		}

		u, err := us.CreateUser(input.Email, input.Name, input.Password)

		if err != nil {
			panic(http_errors.InternalServerError(err))
		}

		result := User{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
		}

		c.JSON(http.StatusOK, gin.H{"user": result})

	}

}

func findUser(us usecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		id := c.Param("id")

		var userID int64
		var err error
		if userID, err = strconv.ParseInt(id, 10, 64); err != nil {
			panic(err)
		}

		u, err := us.FindUserByID(userID)

		if err != nil {
			panic(http_errors.InternalServerError(err))
		}

		if u == nil {
			panic(http_errors.BadRequestError(errors.New("User not found")))
		}

		result := User{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
		}

		c.JSON(http.StatusOK, gin.H{"user": result})

	}

}

func updateUser(us usecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		id := c.Param("id")

		var userID int64
		var err error
		if userID, err = strconv.ParseInt(id, 10, 64); err != nil {
			panic(err)
		}

		var input struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			panic(err)
		}

		err = us.UpdateUser(userID, input.Email, input.Name)

		if err != nil {
			panic(http_errors.InternalServerError(err))
		}

		c.AbortWithStatus(http.StatusOK)

	}

}

func destroyUser(us usecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		id := c.Param("id")

		var userID int64
		var err error
		if userID, err = strconv.ParseInt(id, 10, 64); err != nil {
			panic(err)
		}

		err = us.DestroyUser(userID)

		if err != nil {
			panic(http_errors.InternalServerError(err))
		}

		c.AbortWithStatus(http.StatusOK)

	}

}

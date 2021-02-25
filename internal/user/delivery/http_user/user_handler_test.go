package http_user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-clean-arch/cmd/server/middleware"
	"go-clean-arch/internal/user/entity"
	"go-clean-arch/internal/user/usecase"
	"go-clean-arch/internal/user/usecase/mock"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func AddRouters(service usecase.UseCaseUser) *gin.Engine {
	gin.SetMode(gin.TestMode)

	app := gin.New()

	app.Use(middleware.HandleErrors())

	api := app.Group("api")
	{
		MakeUserRouters(api, service)
	}

	return app
}

func generateUserData() *entity.User {

	return &entity.User{
		ID:        10,
		Email:     "example@gmail.com",
		Name:      "example",
		Password:  "12345678",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}

func TestCreateUser(t *testing.T) {

	t.Run("should be able create user", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockUseCaseUser(ctrl)

		user := generateUserData()

		m.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(user, nil)

		testRouter := AddRouters(m)

		FormUser := &struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}{}

		FormUser.Email = user.Email
		FormUser.Name = user.Name
		FormUser.Password = user.Password

		data, _ := json.Marshal(FormUser)

		req, err := http.NewRequest("POST", "/api/user", bytes.NewBufferString(string(data)))

		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		res := struct {
			User struct {
				Name string
			}
		}{}

		json.Unmarshal(body, &res)

		assert.Equal(t, resp.Code, http.StatusOK)

		assert.Equal(t, res.User.Name, "example")

	})

}

func TestFindUser(t *testing.T) {

	t.Run("should be able find user", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockUseCaseUser(ctrl)

		user := generateUserData()

		userID := int64(10)

		m.EXPECT().FindUserByID(gomock.Any()).Return(user, nil)

		testRouter := AddRouters(m)

		req, err := http.NewRequest("GET", fmt.Sprintf("/api/user/%d", userID), nil)

		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		res := struct {
			User struct {
				ID int
			}
		}{}

		json.Unmarshal(body, &res)

		assert.Equal(t, resp.Code, http.StatusOK)

		assert.Equal(t, res.User.ID, 10)

	})

}

func TestUpdateUser(t *testing.T) {

	t.Run("should be able update user", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockUseCaseUser(ctrl)

		userID := 10

		user := generateUserData()

		m.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		testRouter := AddRouters(m)

		FormUser := &struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}{}

		FormUser.Email = user.Email
		FormUser.Name = user.Name

		data, _ := json.Marshal(FormUser)

		req, err := http.NewRequest("PUT", fmt.Sprintf("/api/user/%d", userID), bytes.NewBufferString(string(data)))

		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		assert.Equal(t, resp.Code, http.StatusOK)

	})

}

func TestDestroyUser(t *testing.T) {

	t.Run("should be able destroy user", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockUseCaseUser(ctrl)

		userID := 10

		m.EXPECT().DestroyUser(gomock.Any()).Return(nil)

		testRouter := AddRouters(m)

		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/user/%d", userID), nil)

		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		assert.Equal(t, resp.Code, http.StatusOK)

	})

}

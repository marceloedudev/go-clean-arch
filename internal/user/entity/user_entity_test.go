package entity_test

import (
	"go-clean-arch/internal/user/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateUserData() *entity.User {

	return &entity.User{
		ID:       10,
		Email:    "test@gmail.com",
		Name:     "Test",
		Password: "12345678",
	}

}

func TestAddUser(t *testing.T) {

	t.Run("should be able add new user", func(t *testing.T) {

		fakeUser := generateUserData()

		user, err := entity.AddUser(fakeUser.Email, fakeUser.Name, fakeUser.Password)

		assert.NotNil(t, user)
		assert.Nil(t, err)

		assert.Equal(t, user.Email, fakeUser.Email)
		assert.Equal(t, user.Name, fakeUser.Name)

		assert.NotEqual(t, user.Password, fakeUser.Password)

	})

}

func TestSetUser(t *testing.T) {

	t.Run("should be able update user", func(t *testing.T) {

		fakeUser := generateUserData()

		user, err := entity.SetUser(fakeUser.Email, fakeUser.Name)

		assert.NotNil(t, user)
		assert.Nil(t, err)

		assert.Equal(t, user.Email, fakeUser.Email)
		assert.Equal(t, user.Name, fakeUser.Name)

	})

}

func TestCheckPassword(t *testing.T) {

	t.Run("should be able check password", func(t *testing.T) {

		fakeUser := generateUserData()

		user, _ := entity.AddUser(fakeUser.Email, fakeUser.Name, fakeUser.Password)
		assert.NotNil(t, user)

		err := user.CheckPassword(fakeUser.Password)
		assert.Nil(t, err)

		err = user.CheckPassword("random@pass")
		assert.NotNil(t, err)

	})

}

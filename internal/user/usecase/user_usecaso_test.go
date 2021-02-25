package usecase

import (
	"go-clean-arch/internal/user/entity"
	"go-clean-arch/internal/user/repository/fakes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generateUserData() *entity.User {

	return &entity.User{
		ID:        1,
		Email:     "test@gmail.com",
		Name:      "Test",
		Password:  "12345678",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}

func TestCreate(t *testing.T) {

	t.Run("should be able create user", func(t *testing.T) {

		fakeUser := generateUserData()
		repo := fakes.NewUserRepositoryFake()
		model := NewService(repo)

		user, err := model.CreateUser(fakeUser.Email, fakeUser.Name, fakeUser.Password)
		assert.NotNil(t, user)
		assert.Nil(t, err)

	})

}

func TestFindUserByID(t *testing.T) {

	t.Run("should be able find user", func(t *testing.T) {

		fakeUser := generateUserData()
		repo := fakes.NewUserRepositoryFake()
		model := NewService(repo)

		user, err := model.CreateUser(fakeUser.Email, fakeUser.Name, fakeUser.Password)
		assert.NotNil(t, user)
		assert.Nil(t, err)

		userFound, err := model.FindUserByID(user.ID)
		assert.NotNil(t, userFound)
		assert.Nil(t, err)

	})

}

func TestUpdate(t *testing.T) {

	t.Run("should be able update user", func(t *testing.T) {

		fakeUser := generateUserData()
		repo := fakes.NewUserRepositoryFake()
		model := NewService(repo)

		user, err := model.CreateUser(fakeUser.Email, fakeUser.Name, fakeUser.Password)
		assert.NotNil(t, user)
		assert.Nil(t, err)

		user.Name = "Example"

		err = model.UpdateUser(user.ID, user.Email, user.Name)
		assert.Nil(t, err)

		userFound, err := model.FindUserByID(user.ID)
		assert.NotNil(t, userFound)
		assert.Nil(t, err)
		assert.Equal(t, "Example", userFound.Name)

	})

}

func TestDestroy(t *testing.T) {

	t.Run("should be able update user", func(t *testing.T) {

		fakeUser := generateUserData()
		repo := fakes.NewUserRepositoryFake()
		model := NewService(repo)

		user, err := model.CreateUser(fakeUser.Email, fakeUser.Name, fakeUser.Password)
		assert.NotNil(t, user)
		assert.Nil(t, err)

		err = model.DestroyUser(user.ID)
		assert.Nil(t, err)

		userFound, err := model.FindUserByID(user.ID)
		assert.Nil(t, userFound)
		assert.NotNil(t, err)

	})

}

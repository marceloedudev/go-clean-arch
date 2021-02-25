package fakes

import (
	"errors"
	"go-clean-arch/internal/user/entity"
)

// UserRepositoryFake struct
type UserRepositoryFake struct {
	fake map[int64]*entity.User
}

// NewUserRepositoryFake fn
func NewUserRepositoryFake() *UserRepositoryFake {
	return &UserRepositoryFake{
		fake: map[int64]*entity.User{},
	}
}

// Create user
func (f *UserRepositoryFake) Create(u *entity.User) (*entity.User, error) {

	u.ID = int64(len(f.fake) + 1)

	f.fake[u.ID] = u

	return u, nil

}

// FindByID user
func (f *UserRepositoryFake) FindByID(userID int64) (user *entity.User, err error) {

	if f.fake[userID] == nil {
		return nil, errors.New("User not found")
	}

	return f.fake[userID], nil

}

// Update user
func (f *UserRepositoryFake) Update(userID int64, u *entity.User) (err error) {

	if _, err := f.FindByID(userID); err != nil {
		return err
	}

	f.fake[userID] = u

	return nil

}

// Destroy user
func (f *UserRepositoryFake) Destroy(userID int64) error {

	if _, err := f.FindByID(userID); err != nil {
		return err
	}

	f.fake[userID] = nil

	return nil

}

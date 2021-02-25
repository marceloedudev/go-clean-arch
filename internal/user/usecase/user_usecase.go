package usecase

import (
	"errors"
	"go-clean-arch/internal/user/entity"
)

// Service struct
type Service struct {
	repository Repository
}

// NewService Handle
func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

// CreateUser Create user
func (s *Service) CreateUser(email, name, password string) (user *entity.User, err error) {

	if user, err = entity.AddUser(email, name, password); err != nil {
		return user, err
	}

	return s.repository.Create(user)

}

// FindUserByID Get user by id
func (s *Service) FindUserByID(userID int64) (*entity.User, error) {

	if userID <= 0 {
		return nil, errors.New("User id invalid")
	}

	return s.repository.FindByID(userID)

}

// UpdateUser Update user
func (s *Service) UpdateUser(userID int64, email string, name string) error {

	if _, err := s.FindUserByID(userID); err != nil {
		return err
	}

	user, err := entity.SetUser(email, name)
	if err != nil {
		return err
	}

	return s.repository.Update(userID, user)

}

// DestroyUser Delete user
func (s *Service) DestroyUser(userID int64) error {

	if _, err := s.FindUserByID(userID); err != nil {
		return err
	}

	return s.repository.Destroy(userID)

}

package usecase

import "go-clean-arch/internal/user/entity"

// Repository interface
type Repository interface {
	Create(e *entity.User) (*entity.User, error)
	FindByID(userID int64) (*entity.User, error)
	Update(userID int64, u *entity.User) error
	Destroy(userID int64) error
}

// UseCaseUser interface
type UseCaseUser interface {
	CreateUser(email, name, password string) (*entity.User, error)
	FindUserByID(userID int64) (*entity.User, error)
	UpdateUser(userID int64, email string, name string) error
	DestroyUser(userID int64) error
}

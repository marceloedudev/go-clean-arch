package repository

import (
	"go-clean-arch/internal/user/entity"

	"github.com/jmoiron/sqlx"
)

// UserRepository struct
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository repo
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create user
func (u *UserRepository) Create(e *entity.User) (user *entity.User, err error) {

	var id int64
	if err := u.db.QueryRow(createUserQuery, e.Name, e.Password, e.Email, e.CreatedAt, e.UpdatedAt).Scan(&id); err != nil {
		return nil, err
	}

	e.ID = id

	return e, nil

}

// FindByID get user by id
func (u *UserRepository) FindByID(userID int64) (user *entity.User, err error) {

	user = &entity.User{}

	if err := u.db.QueryRow(findUserByIDQuery, userID).Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return user, nil

}

// Update user by id
func (u *UserRepository) Update(userID int64, doc *entity.User) (err error) {

	if _, err := u.db.Exec(updateUserQuery, doc.Name, doc.Email, doc.UpdatedAt, userID); err != nil {
		return err
	}

	return nil

}

// Destroy user by id
func (u *UserRepository) Destroy(userID int64) error {

	if _, err := u.db.Exec(destroyUserQuery, userID); err != nil {
		return err
	}

	return nil

}

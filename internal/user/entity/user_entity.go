package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User data
type User struct {
	ID        int64
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AddUser create user
func AddUser(email, name, password string) (*User, error) {

	hashPassword, err := generateHashPassword(password)

	if err != nil {
		return nil, err
	}

	rUser := &User{
		Email:     email,
		Name:      name,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = rUser.ValidateAddUser()

	if err != nil {
		return nil, err
	}

	return rUser, nil

}

// ValidateAddUser valitador
func (u *User) ValidateAddUser() error {

	if u.Email == "" {
		return errors.New("Field email is required")
	}

	if len(u.Email) > 100 {
		return errors.New("Email cannot be longer than 100")
	}

	if u.Name == "" {
		return errors.New("Field name is required")
	}

	if len(u.Name) > 60 {
		return errors.New("Name cannot be longer than 60")
	}

	if u.Password == "" {
		return errors.New("Field password is required")
	}

	if len(u.Password) < 8 || len(u.Password) > 64 {
		return errors.New("Password cannot be less than 8 or greater than 64 ")
	}

	return nil

}

// SetUser update user
func SetUser(email, name string) (*User, error) {

	rUser := &User{
		Email:     email,
		Name:      name,
		UpdatedAt: time.Now(),
	}

	if err := rUser.ValidateUpdateUser(); err != nil {
		return nil, err
	}

	return rUser, nil

}

// ValidateUpdateUser valitador
func (u *User) ValidateUpdateUser() error {

	if u.Email == "" {
		return errors.New("Field email is required")
	}

	if len(u.Email) > 100 {
		return errors.New("Email cannot be longer than 100")
	}

	if u.Name == "" {
		return errors.New("Field name is required")
	}

	if len(u.Name) > 60 {
		return errors.New("Name cannot be longer than 60")
	}

	return nil

}

// generateHashPassword
func generateHashPassword(password string) (string, error) {

	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil

}

// CheckPassword compare user password
func (u *User) CheckPassword(password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil

}

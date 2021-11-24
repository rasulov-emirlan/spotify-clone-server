package models

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

func TestUser() *User {
	return &User{
		ID:       1,
		Name:     "John Doe",
		Password: "123",
		Email:    "johndoe@gmail.com",
	}
}

func (u *User) BeforeCreate() error {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}
	newpwd, err := hashPassword(u.Password)
	u.Password = newpwd
	return err
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

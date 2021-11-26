package models

import (
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id" example:"0"`
	Name     string `json:"name" example:"John Doe"`
	Password string `json:"-" example:"QWERTY123"`
	Email    string `json:"-" example:"john@gmail.com"`
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

func (u *User) ComparePasswords(passwrod string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwrod))
}

func (u *User) GenerateJWT(jwtSecretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["name"] = u.Name
	claims["userid"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(jwtSecretKey))
}

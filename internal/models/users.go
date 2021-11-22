package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

func TestUser() *User {
	return &User{
		Name:     "John Doe",
		Password: "123",
		Email:    "johndoe@gmail.com",
	}
}

func userInitializer() {

}

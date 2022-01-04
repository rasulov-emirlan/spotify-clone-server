package models

const (
	Admin  = "admin"
	Singer = "singer"
)

type Role struct {
	Name string `json:"name"`
}

type Roles []string

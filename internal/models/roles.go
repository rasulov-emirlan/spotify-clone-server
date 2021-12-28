package models

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Roles []*Role

package models

type Genre struct {
	ID   int    `json"id"`
	Name string `json:"name"`
}

func TestGenre() *Genre {
	return &Genre{
		ID:   1,
		Name: "Classics",
	}
}

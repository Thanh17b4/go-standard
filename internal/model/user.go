package model

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address *string `json:"address"`
}

package model

type User struct {
	ID        string `json: "_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdat"`
}

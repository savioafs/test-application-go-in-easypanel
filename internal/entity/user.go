package entity

import "github.com/google/uuid"

type User struct {
	ID    string `json:"id" gorm:"primaryKey;type:uuid"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex"`
}

func NewUser(name string, email string) User {
	return User{
		ID:    uuid.NewString(),
		Name:  name,
		Email: email,
	}
}

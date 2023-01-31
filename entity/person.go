package entity

import "github.com/google/uuid"

type Person struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

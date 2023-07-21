package model

import (
	"github.com/google/uuid"
)

type Books struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	Price     float32   `json:"price"`
}

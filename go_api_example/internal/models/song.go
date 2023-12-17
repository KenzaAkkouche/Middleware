package models

import (
	"github.com/gofrs/uuid"
)
type Song struct {
	Id    *uuid.UUID `json:"id"`
	Title  string     `json:"title"`
	Artist string     `json:"artist"`
}

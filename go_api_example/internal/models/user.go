package models

import (
	"github.com/gofrs/uuid"
)
type User struct {
	Id    *uuid.UUID `json:"id"`
	Name  string     `json:"name"`
	Username string     `json:"username"`
	Password string     `json:"password"`

	// Ajoutez d'autres champs spécifiques à l'utilisateur
}

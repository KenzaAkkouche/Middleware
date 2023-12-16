package models

import (
	"github.com/gofrs/uuid"
)
type User struct {
	Id    *uuid.UUID `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
	// Ajoutez d'autres champs spécifiques à l'utilisateur
}

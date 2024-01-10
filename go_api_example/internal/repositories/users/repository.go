package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"errors"
)

// ErrUserNotFound est une erreur indiquant qu'un utilisateur n'a pas été trouvé.
var ErrUserNotFound = errors.New("user not found")

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Username,  &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.Name, &user.Username,  &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// UpdateUser met à jour un utilisateur dans la base de données
func UpdateUser(id uuid.UUID, newName, newUsername, newPassword string) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Vérifiez d'abord si l'utilisateur existe
	existingUser, err := GetUserById(id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, ErrUserNotFound
	}

	// Mettez à jour l'utilisateur dans la base de données
	_, err = db.Exec("UPDATE users SET name=?, username=?, password=? WHERE id=?", newName, newUsername, newPassword, id.String())
	if err != nil {
		return nil, err
	}

	// Récupérez l'utilisateur mis à jour depuis la base de données
	updatedUser, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// CreateUser crée un nouvel utilisateur dans la base de données
func CreateUser(name, username, password string) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Générez un nouvel UUID pour l'utilisateur
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Insérez le nouvel utilisateur dans la base de données
	_, err = db.Exec("INSERT INTO users (id, name, username, password) VALUES (?, ?, ?, ?)", newID.String(), name, username, password)
	if err != nil {
		return nil, err
	}

	// Récupérez l'utilisateur nouvellement créé depuis la base de données
	createdUser, err := GetUserById(newID)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// DeleteUser supprime un utilisateur de la base de données en utilisant son ID.
func DeleteUser(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Vérifiez d'abord si l'utilisateur existe
	_, err = GetUserById(id)
	if err != nil {
		return err
	}

	// Supprime l'utilisateur de la base de données
	_, err = db.Exec("DELETE FROM users WHERE id=?", id.String())
	if err != nil {
		return err
	}

	return nil
}

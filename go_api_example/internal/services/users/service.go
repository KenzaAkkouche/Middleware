package users

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"
	"net/http"
)


// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
	var err error
	// Call the repository
	users, err := repository.GetAllUsers()
	// Handle errors
	if err != nil {
		logrus.Errorf("error retrieving users: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

// GetUserById retrieves a user from the database by ID.
func GetUserById(id uuid.UUID) (*models.User, error) {
	// Call the repository to get the user by ID
	user, err := repository.GetUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

// CreateUser creates a new user in the database.
func CreateUser(name, email string) (*models.User, error) {
	// Call the repository to create a new user
	createdUser, err := repository.CreateUser(name, email)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return createdUser, nil
}

// UpdateUser updates a user in the database.
func UpdateUser(id uuid.UUID, newName, newEmail string) (*models.User, error) {
	// Call the repository to update the user
	updatedUser, err := repository.UpdateUser(id, newName, newEmail)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating user: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return updatedUser, nil
}

// DeleteUserByID deletes a user from the database using their ID.
func DeleteUserByID(id uuid.UUID) error {
	// Call the repository to delete the user
	err := repository.DeleteUser(id)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting user: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}

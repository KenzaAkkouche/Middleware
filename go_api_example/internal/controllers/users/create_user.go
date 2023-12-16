package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
)

// CreateUser
// @Tags         users
// @Summary      Create a new user.
// @Description  Create a new user.
// @Param        content       body   string  true  "Content of the new user"
// @Success      201            {object}  models.user
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var userRequest models.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		logrus.Errorf("error decoding request payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate request payload
	if userRequest.Name == "" {
		logrus.Error("Name cannot be empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userRequest.Email == "" {
    		logrus.Error("Email cannot be empty")
    		w.WriteHeader(http.StatusBadRequest)
    		return
    	}

	// Call service to create the user
	createdUser, err := users.CreateUser(userRequest.Name, userRequest.Email)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the created user
	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(createdUser)
	_, _ = w.Write(responseBody)
}

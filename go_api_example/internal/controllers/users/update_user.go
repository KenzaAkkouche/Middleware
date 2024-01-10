package users

import (
"encoding/json"
"net/http"
"errors"
"github.com/go-chi/chi/v5"
"github.com/sirupsen/logrus"
"middleware/example/internal/models"
"middleware/example/internal/services/users"
"github.com/gofrs/uuid"
)

// UpdateUser
// @Tags         users
// @Summary      Update an existing user.
// @Description  Update an existing user.
// @Param        id           path      string  true  "user UUID formatted ID"
// @Param        content       body      string  true  "New content for the user"
// @Success      200            {object}  models.User
// @Failure      400            "Invalid request payload"
// @Failure      404            "User not found"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [put]

// ErrUserNotFound is an error indicating that a user was not found.
var ErrUserNotFound = errors.New("user not found")

func UpdateUser(w http.ResponseWriter, r *http.Request) {
// Parse request body
var userRequest models.User
err := json.NewDecoder(r.Body).Decode(&userRequest)
if err != nil {
logrus.Errorf("error decoding request payload: %s", err.Error())
w.WriteHeader(http.StatusBadRequest)
return
}

// Validate request payload
if userRequest.Name == "" || userRequest.Username == "" || userRequest.Password == "" {
logrus.Error("name, username, and password cannot be empty")
w.WriteHeader(http.StatusBadRequest)
return
}

// Get user ID from URL parameters
userID, err := uuid.FromString(chi.URLParam(r, "id"))
if err != nil {
logrus.Errorf("parsing error : %s", err.Error())
w.WriteHeader(http.StatusUnprocessableEntity)
return
}

// Call service to update the user
updatedUser, err := users.UpdateUser(userID, userRequest.Name, userRequest.Username, userRequest.Password)
if err != nil {
if err == ErrUserNotFound {
logrus.Errorf("user not found: %s", err.Error())
w.WriteHeader(http.StatusNotFound)
} else {
logrus.Errorf("error updating user: %s", err.Error())
w.WriteHeader(http.StatusInternalServerError)
}
return
}

// Respond with the updated user
w.WriteHeader(http.StatusOK)
responseBody, _ := json.Marshal(updatedUser)
_, _ = w.Write(responseBody)
}

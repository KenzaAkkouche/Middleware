package users
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/services/users"
	"github.com/gofrs/uuid"
)
// DeleteUser
// @Tags         users
// @Summary      Delete an existing user.
// @Description  Delete an existing user.
// @Param        id           	path      string  true  "user UUID formatted ID"
// @Success      204            "No Content"
// @Failure      404            "user not found"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL parameters
	userID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("parsing error: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call service to delete the user
	err = users.DeleteUserByID(userID)
	if err != nil {
		if err == ErrUserNotFound {
			logrus.Errorf("user not found: %s", err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			logrus.Errorf("error deleting user: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Respond with No Content status
	w.WriteHeader(http.StatusNoContent)
}

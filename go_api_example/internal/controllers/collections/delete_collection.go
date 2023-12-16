package collections
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/services/collections"
	"github.com/gofrs/uuid"
)
// DeleteCollection
// @Tags         collections
// @Summary      Delete an existing collection.
// @Description  Delete an existing collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      204            "No Content"
// @Failure      404            "Collection not found"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [delete]
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	// Get collection ID from URL parameters
	collectionID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("parsing error: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call service to delete the collection
	err = collections.DeleteCollectionByID(collectionID)
	if err != nil {
		if err == ErrCollectionNotFound {
			logrus.Errorf("collection not found: %s", err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			logrus.Errorf("error deleting collection: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Respond with No Content status
	w.WriteHeader(http.StatusNoContent)
}

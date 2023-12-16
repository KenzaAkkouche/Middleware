package collections

import (
	"encoding/json"
	"net/http"
    "errors"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"github.com/gofrs/uuid"
)

// UpdateCollection
// @Tags         collections
// @Summary      Update an existing collection.
// @Description  Update an existing collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Param        content       body      string  true  "New content for the collection"
// @Success      200            {object}  models.Collection
// @Failure      400            "Invalid request payload"
// @Failure      404            "Collection not found"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [put]

// ErrCollectionNotFound est une erreur indiquant qu'une collection n'a pas été trouvée.
var ErrCollectionNotFound = errors.New("collection not found")

func UpdateCollection(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var collectionRequest models.Collection
	err := json.NewDecoder(r.Body).Decode(&collectionRequest)
	if err != nil {
		logrus.Errorf("error decoding request payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate request payload
	if collectionRequest.Content == "" {
		logrus.Error("content cannot be empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get collection ID from URL parameters
	collectionID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("parsing error : %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call service to update the collection
	updatedCollection, err := collections.UpdateCollection(collectionID, collectionRequest.Content)
	if err != nil {
		if err == ErrCollectionNotFound {
			logrus.Errorf("collection not found: %s", err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			logrus.Errorf("error updating collection: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Respond with the updated collection
	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(updatedCollection)
	_, _ = w.Write(responseBody)
}

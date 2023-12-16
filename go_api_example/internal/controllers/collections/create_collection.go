package collections

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"
)

// CreateCollection
// @Tags         collections
// @Summary      Create a new collection.
// @Description  Create a new collection.
// @Param        content       body   string  true  "Content of the new collection"
// @Success      201            {object}  models.Collection
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /collections [post]
func CreateCollection(w http.ResponseWriter, r *http.Request) {
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

	// Call service to create the collection
	createdCollection, err := collections.CreateCollection(collectionRequest.Content)
	if err != nil {
		logrus.Errorf("error creating collection: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the created collection
	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(createdCollection)
	_, _ = w.Write(responseBody)
}

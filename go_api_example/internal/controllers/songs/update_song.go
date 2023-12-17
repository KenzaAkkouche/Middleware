package songs

import (
	"encoding/json"
	"net/http"
    "errors"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"github.com/gofrs/uuid"
)

// UpdateSong
// @Tags         songs
// @Summary      Update an existing song.
// @Description  Update an existing song.
// @Param        id           	path      string  true  "song UUID formatted ID"
// @Param        content       body      string  true  "New content for the song"
// @Success      200            {object}  models.song
// @Failure      400            "Invalid request payload"
// @Failure      404            "song not found"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [put]

// ErrSongNotFound est une erreur indiquant qu'une song n'a pas été trouvée.
var ErrSongNotFound = errors.New("song not found")

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var songRequest models.Song
	err := json.NewDecoder(r.Body).Decode(&songRequest)
	if err != nil {
		logrus.Errorf("error decoding request payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate request payload
    if songRequest.Title == "" || songRequest.Artist == "" {
        logrus.Error("title and artist cannot be empty")
        w.WriteHeader(http.StatusBadRequest)
        return
    }

	// Get song ID from URL parameters
	songID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("parsing error : %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call service to update the song
	updatedSong, err := songs.UpdateSong(songID, songRequest.Title, songRequest.Artist)
	if err != nil {
		if err == ErrSongNotFound {
			logrus.Errorf("song not found: %s", err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			logrus.Errorf("error updating song: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Respond with the updated song
	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(updatedSong)
	_, _ = w.Write(responseBody)
}
package songs

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"
)

// CreateSong
// @Tags         songs
// @Summary      Create a new song.
// @Description  Create a new song.
// @Param        content       body   string  true  "Content of the new song"
// @Success      201            {object}  models.song
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /songs [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var songRequest models.Song
	err := json.NewDecoder(r.Body).Decode(&songRequest)
	if err != nil {
		logrus.Errorf("error decoding request payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate request payload
	if songRequest.Title == "" {
		logrus.Error("Title cannot be empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if songRequest.Artist == "" {
    		logrus.Error("Artist cannot be empty")
    		w.WriteHeader(http.StatusBadRequest)
    		return
    	}

	// Call service to create the song
	createdSong, err := songs.CreateSong(songRequest.Title, songRequest.Artist)
	if err != nil {
		logrus.Errorf("error creating song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the created song
	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(createdSong)
	_, _ = w.Write(responseBody)
}
package songs
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/services/songs"
	"github.com/gofrs/uuid"
)
// DeleteSong
// @Tags         songs
// @Summary      Delete an existing song.
// @Description  Delete an existing song.
// @Param        id           	path      string  true  "song UUID formatted ID"
// @Success      204            "No Content"
// @Failure      404            "song not found"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Get song ID from URL parameters
	songID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("parsing error: %s", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Call service to delete the song
	err = songs.DeleteSongByID(songID)
	if err != nil {
		if err == ErrSongNotFound {
			logrus.Errorf("Song not found: %s", err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			logrus.Errorf("error deleting song: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Respond with No Content status
	w.WriteHeader(http.StatusNoContent)
}
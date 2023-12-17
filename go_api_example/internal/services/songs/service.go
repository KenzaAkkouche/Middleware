package songs

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/songs"
	"net/http"
)


// GetAllsongs retrieves all songs from the database.
func GetAllSongs() ([]models.Song, error) {
	var err error
	// Call the repository
	songs, err := repository.GetAllSongs()
	// Handle errors
	if err != nil {
		logrus.Errorf("error retrieving songs: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

// GetSongById retrieves a song from the database by ID.
func GetSongById(id uuid.UUID) (*models.Song, error) {
	// Call the repository to get the song by ID
	song, err := repository.GetSongById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

// Createsong creates a new song in the database.
func CreateSong(title, artist string) (*models.Song, error) {
	// Call the repository to create a new song
	createdSong, err := repository.CreateSong(title, artist)
	if err != nil {
		logrus.Errorf("error creating song: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return createdSong, nil
}

// UpdateSong updates a song in the database.
func UpdateSong(id uuid.UUID, newTitle, newArtist string) (*models.Song, error) {
	// Call the repository to update the song
	updatedSong, err := repository.UpdateSong(id, newTitle, newArtist)
	if err != nil {
		if err == repository.ErrSongNotFound {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating song: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return updatedSong, nil
}

// DeleteSongByID deletes a song from the database using their ID.
func DeleteSongByID(id uuid.UUID) error {
	// Call the repository to delete the song
	err := repository.DeleteSong(id)
	if err != nil {
		if err == repository.ErrSongNotFound {
			return &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting song: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}
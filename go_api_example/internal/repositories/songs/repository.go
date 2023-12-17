package songs

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"errors"
)

// ErrSongNotFound est une erreur indiquant qu'un utilisateur n'a pas été trouvé.
var ErrSongNotFound = errors.New("song not found")

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Title, &data.Artist)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Title, &song.Artist)
	if err != nil {
		return nil, err
	}
	return &song, err
}

// UpdateSong met à jour un utilisateur dans la base de données
func UpdateSong(id uuid.UUID, newTitle, newArtist string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Vérifiez d'abord si l'utilisateur existe
	existingSong, err := GetSongById(id)
	if err != nil {
		return nil, err
	}
	if existingSong == nil {
		return nil, ErrSongNotFound
	}

	// Mettez à jour l'utilisateur dans la base de données
	_, err = db.Exec("UPDATE songs SET title=?, artist=? WHERE id=?", newTitle, newArtist, id.String())
	if err != nil {
		return nil, err
	}

	// Récupérez l'utilisateur mis à jour depuis la base de données
	updatedSong, err := GetSongById(id)
	if err != nil {
		return nil, err
	}

	return updatedSong, nil
}

// CreateSong crée un nouvel utilisateur dans la base de données
func CreateSong(title, artist string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Générez un nouvel UUID pour l'utilisateur
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Insérez le nouvel utilisateur dans la base de données
	_, err = db.Exec("INSERT INTO songs (id, title, artist) VALUES (?, ?, ?)", newID.String(), title, artist)
	if err != nil {
		return nil, err
	}

	// Récupérez l'utilisateur nouvellement créé depuis la base de données
	createdSong, err := GetSongById(newID)
	if err != nil {
		return nil, err
	}

	return createdSong, nil
}

// DeleteSong supprime un utilisateur de la base de données en utilisant son ID.
func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Vérifiez d'abord si l'utilisateur existe
	_, err = GetSongById(id)
	if err != nil {
		return err
	}

	// Supprime l'utilisateur de la base de données
	_, err = db.Exec("DELETE FROM songs WHERE id=?", id.String())
	if err != nil {
		return err
	}

	return nil
}
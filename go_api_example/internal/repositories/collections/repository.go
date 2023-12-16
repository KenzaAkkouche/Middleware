package collections

import (

	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"errors"
)

// ErrCollectionNotFound est une erreur indiquant qu'une collection n'a pas été trouvée.
var ErrCollectionNotFound = errors.New("collection not found")

func GetAllCollections() ([]models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM collections")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.Collection{}
	for rows.Next() {
		var data models.Collection
		err = rows.Scan(&data.Id, &data.Content)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetCollectionById(id uuid.UUID) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM collections WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.Collection
	err = row.Scan(&collection.Id, &collection.Content)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
// UpdateCollection met à jour une collection dans la base de données
func UpdateCollection(id uuid.UUID, newContent string) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Vérifiez d'abord si la collection existe
	existingCollection, err := GetCollectionById(id)
	if err != nil {
		return nil, err
	}
	if existingCollection == nil {
		return nil, ErrCollectionNotFound
	}

	// Mettez à jour la collection dans la base de données
	_, err = db.Exec("UPDATE collections SET content=? WHERE id=?", newContent, id.String())
	if err != nil {
		return nil, err
	}

	// Récupérez la collection mise à jour depuis la base de données
	updatedCollection, err := GetCollectionById(id)
	if err != nil {
		return nil, err
	}

	return updatedCollection, nil
}
// CreateCollection crée une nouvelle collection dans la base de données
func CreateCollection(content string) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Générez un nouvel UUID pour la collection
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Insérez la nouvelle collection dans la base de données
	_, err = db.Exec("INSERT INTO collections (id, content) VALUES (?, ?)", newID.String(), content)
	if err != nil {
		return nil, err
	}

	// Récupérez la collection nouvellement créée depuis la base de données
	createdCollection, err := GetCollectionById(newID)
	if err != nil {
		return nil, err
	}

	return createdCollection, nil
}
// DeleteCollection supprime une collection de la base de données en utilisant son ID.
func DeleteCollection(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Vérifiez d'abord si la collection existe
	_, err = GetCollectionById(id)
	if err != nil {
		return err
	}

	// Supprime la collection de la base de données
	_, err = db.Exec("DELETE FROM collections WHERE id=?", id.String())
	if err != nil {
		return err
	}

	return nil
}
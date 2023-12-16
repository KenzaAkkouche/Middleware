package collections

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllCollections() ([]models.Collection, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllCollections()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func GetCollectionById(id uuid.UUID) (*models.Collection, error) {
	collection, err := repository.GetCollectionById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}
func CreateCollection(content string) (*models.Collection, error) {
    	// appel du repository pour créer une nouvelle collection
    	createdCollection, err := repository.CreateCollection(content)
    	if err != nil {
    		logrus.Errorf("erreur lors de la création de la collection : %s", err.Error())
    		return nil, &models.CustomError{
    			Message: "Quelque chose s'est mal passé",
    			Code:    500,
    		}
    	}

    	return createdCollection, nil
  }


	func UpdateCollection(id uuid.UUID, newContent string) (*models.Collection, error) {
        // Appel du repository pour mettre à jour la collection
        updatedCollection, err := repository.UpdateCollection(id, newContent)
        if err != nil {
            if err == repository.ErrCollectionNotFound {
                return nil, &models.CustomError{
                    Message: "collection not found",
                    Code:    http.StatusNotFound,
                }
            }
            logrus.Errorf("error updating collection : %s", err.Error())
            return nil, &models.CustomError{
                Message: "Something went wrong",
                Code:    500,
            }
        }

        return updatedCollection, nil
    }
// DeleteCollectionByID supprime une collection de la base de données en utilisant son ID.
func DeleteCollectionByID(id uuid.UUID) error {
	// Appel du repository pour supprimer la collection
	err := repository.DeleteCollection(id)
	if err != nil {
		if err == repository.ErrCollectionNotFound {
			return &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting collection : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}


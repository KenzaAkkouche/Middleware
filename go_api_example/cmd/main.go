package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/collections"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/collections", func(r chi.Router) {
		r.Get("/", collections.GetCollections)
		r.Post("/", collections.CreateCollection)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetCollection)
			r.Put("/", collections.UpdateCollection)
			r.Delete("/", collections.DeleteCollection)

		})
	})


    	r.Route("/users", func(r chi.Router) {
    		r.Get("/", users.GetUsers)
    		r.Post("/", users.CreateUser)
    		r.Route("/{id}", func(r chi.Router) {
    			r.Use(users.Ctx)
    			r.Get("/", users.GetUser)
    			r.Put("/", users.UpdateUser)
    			r.Delete("/", users.DeleteUser)

    		})
    	})
	logrus.Info("[INFO] Web server started. Now listening on *:8081")
	logrus.Fatalln(http.ListenAndServe(":8081", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS collections (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			content VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS users (
        			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
        			name VARCHAR(255) NOT NULL,
        			email VARCHAR(255) NOT NULL
        		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}

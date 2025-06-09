/*
	Represents the core application, containing a
	logger, server handler, and database connection.

	Used by main.go
*/

package app

import (
	"database/sql"
	"fmt"
	"go-backend-server-template/internal/api"
	"go-backend-server-template/internal/store"
	"go-backend-server-template/migrations"
	"go-backend-server-template/utils"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger        *log.Logger
	ServerHandler *api.ServerHandler
	DB            *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	pgDB, err := store.Open()

	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// repository layer
	dataStore := store.NewPostgresStore(pgDB)

	// controller layer
	serverHandler := api.NewServerhandler(dataStore, logger)

	app := &Application{
		Logger:        logger,
		ServerHandler: serverHandler,
		DB:            pgDB,
	}

	return app, nil
}

func (a Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	a.Logger.Printf("Status is avaiable!\n")
	utils.WriteJson(w, http.StatusOK, utils.Envelope{"status": "OK"})
}

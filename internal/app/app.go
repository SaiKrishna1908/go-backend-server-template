package app

import (
	"database/sql"
	"fmt"
	"go-backend-server-template/internal/api"
	"go-backend-server-template/internal/store"
	"go-backend-server-template/migrations"
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

	dataStore := store.NewPostgresStore(pgDB)

	// handlers
	serverHandler := api.NewServerhandler(dataStore)

	app := &Application{
		Logger:        logger,
		ServerHandler: serverHandler,
		DB:            pgDB,
	}

	return app, nil
}

func (a Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available!\n")
}

/*
	Handlers are responsible for the logic that
	gets executed when a route is hit.

	Used by app.go
*/

package api

import (
	"go-backend-server-template/internal/store"
	"log"
	"net/http"
)

type ServerHandler struct {
	dataStore store.DataStore
	logger    *log.Logger
}

func NewServerhandler(dataStore store.DataStore, logger *log.Logger) *ServerHandler {
	return &ServerHandler{
		dataStore: dataStore,
		logger:    logger,
	}
}

func (sh *ServerHandler) HandleGetEntityById(w http.ResponseWriter, r http.Request) {

}

func (sh *ServerHandler) HandleCreateEntity(w http.ResponseWriter, r *http.Request) {

}

func (sh *ServerHandler) HandleUpdateEntity(w http.ResponseWriter, r *http.Request) {

}

func (sh *ServerHandler) HandleDeleteEntity(w http.ResponseWriter, r *http.Request) {

}

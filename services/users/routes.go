package users

import (
	"database/sql"

	"github.com/gorilla/mux"
)

const userPath = "/users"

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc(userPath+"/{id:[0-9]+}", h.handleGetUser).Methods("GET")
	router.HandleFunc(userPath, h.handleGetAllUsers).Methods("GET")
	router.HandleFunc(userPath, h.handleAddUser).Methods("POST")
	router.HandleFunc(userPath+"/{id:[0-9]+}", h.handleUpdateUser).Methods("PUT")
	router.HandleFunc(userPath+"/{id:[0-9]+}", h.handleDeleteUser).Methods("DELETE")
}

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
	router.HandleFunc(userPath+"/{userID}", h.handleGetUser).Methods("GET")
	router.HandleFunc(userPath, h.handleAddUser).Methods("POST")
	router.HandleFunc(userPath+"/{userID}", h.handleUpdateUser).Methods("PUT")
	router.HandleFunc(userPath, h.handleGetAllUsers).Methods("GET")
	router.HandleFunc(userPath+"/{userID}", h.handleDeleteUser).Methods("DELETE")
}

package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akeemdnl/go-crud/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

const prefix = "/users"

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc(prefix+"/{userID}", h.handleGetUser).Methods("GET")
	router.HandleFunc(prefix, h.handleAddUser).Methods("POST")
	router.HandleFunc(prefix+"/{userID}", h.handleUpdateUser).Methods("PUT")
	router.HandleFunc(prefix, h.handleGetAllUsers).Methods("GET")
	router.HandleFunc(prefix+"/{userID}", h.handleDeleteUser).Methods("DELETE")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUrlVariable("userID", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	user, err := getUserById(h.db, userID)
	if err == sql.ErrNoRows {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("User not found"))
		return
	}
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleAddUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := addUser(h.db, user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, user)
}

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["userID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user Id"))
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var updateUserPayload UpdateUserPayload
	if err := utils.ParseJSON(r, &updateUserPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := updateUser(h.db, &updateUserPayload, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := getAllUsers(h.db)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUrlVariable("userID", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := deleteUser(h.db, userID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.JsonMessage("Successfully deleted user"))
}

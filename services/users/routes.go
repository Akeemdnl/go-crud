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
	router.HandleFunc(prefix, h.handleGetUser).Methods("GET")
	router.HandleFunc(prefix, h.handleAddUser).Methods("POST")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user id"))
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	user, err := getUserBy(h.db, userID, "id")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("User not found"))
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

	_, err := h.db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, user)
}

package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akeemdnl/go-crud/utils"
)

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUrlVariable("id", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	user, err := getUserBy(h.db, "id", userID)
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

func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := getAllUsers(h.db)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleGetUserByName(w http.ResponseWriter, r *http.Request) {
	getUser(w, r, "name", h)
}

func (h *Handler) handleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	getUser(w, r, "email", h)
}

func getUser(w http.ResponseWriter, r *http.Request, param string, h *Handler) {
	value, err := utils.GetQueryParam(param, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := getUserBy(h.db, param, value)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

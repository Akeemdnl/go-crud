package users

import (
	"fmt"
	"net/http"

	"github.com/Akeemdnl/go-crud/utils"
	"github.com/go-playground/validator/v10"
)

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

	utils.WriteJSON(w, http.StatusCreated, user)
}

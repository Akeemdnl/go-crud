package users

import (
	"net/http"
	"strconv"

	"github.com/Akeemdnl/go-crud/utils"
)

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

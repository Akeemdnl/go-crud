package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akeemdnl/go-crud/utils"
	"github.com/gorilla/mux"
)

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
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

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(addr string) error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/users/{userID}", handleGetUser).Methods("GET")
	fmt.Println("Listening on", addr)
	return http.ListenAndServe(addr, router)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userID"]
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "OOPS"})
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "OK", "ID": str})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

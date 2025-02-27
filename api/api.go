package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Akeemdnl/go-crud/services/users"
	"github.com/Akeemdnl/go-crud/utils"
	"github.com/gorilla/mux"
)

func Run(addr string, db *sql.DB) error {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", handleHealthCheck).Methods("GET")
	v1 := router.PathPrefix("/api/v1").Subrouter()

	userHandler := users.NewHandler(db)
	userHandler.RegisterRoutes(v1)

	fmt.Println("Listening on", addr)
	return http.ListenAndServe(addr, router)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, utils.JsonMessage("OK!"))
}

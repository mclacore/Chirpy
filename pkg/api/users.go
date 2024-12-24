package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var resBody User

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&resBody)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid request payload")
		return
	}

	userEmail := sql.NullString{String: resBody.Email, Valid: true}
	createdUser, createErr := cfg.Database.CreateUser(r.Context(), userEmail)
	if createErr != nil {
		log.Println("Error creating user:", createErr)
		RespondWithError(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	dat, _ := json.Marshal(createdUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(dat)
}

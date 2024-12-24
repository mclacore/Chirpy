package api

import (
	"encoding/json"
	"log"
	"net/http"
)

var user User

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid request payload")
		return
	}

	createdUser, createErr := cfg.Database.CreateUser(r.Context(), user.Email)
	if createErr != nil {
		log.Println("Error creating user:", createErr)
		RespondWithError(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	RespondWithJSON(w, http.StatusCreated, createdUser)
}

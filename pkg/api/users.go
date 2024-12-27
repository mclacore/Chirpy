package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mclacore/Chirpy/internal/database"
)

var user User

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid request payload")
		return
	}

	userParams := database.CreateUserParams{
		Email:          user.Email,
		HashedPassword: user.Password,
	}

	createdUser, createErr := cfg.Database.CreateUser(r.Context(), userParams)
	if createErr != nil {
		log.Println("Error creating user:", createErr)
		RespondWithError(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	RespondWithJSON(w, http.StatusCreated, createdUser)
}

package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	var resBody user

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

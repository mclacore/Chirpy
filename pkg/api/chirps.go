package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mclacore/Chirpy/internal/database"
)

var chirp Chirp

func (cfg *ApiConfig) PostChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chirp)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(chirp.Body) > 140 {
		RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	chirp.CleanedBody = profaneToAsterisks(chirp.Body)

	chirpParams := database.PostChirpParams{
		Body:   chirp.Body,
		UserID: chirp.UserId,
	}
	postChirp, postErr := cfg.Database.PostChirp(r.Context(), chirpParams)
	if postErr != nil {
		log.Println("Error posting chirp:", postErr)
		RespondWithError(w, http.StatusInternalServerError, "Could not post chirp")
		return
	}

	RespondWithJSON(w, http.StatusCreated, postChirp)
}

func (cfg *ApiConfig) GetChirps(w http.ResponseWriter, r *http.Request) {
	getChirps, getErr := cfg.Database.GetChirps(r.Context())
	if getErr != nil {
		log.Println("Error fetching chirps:", getErr)
		RespondWithError(w, http.StatusInternalServerError, "Could not fetch chirps")
		return
	}

	RespondWithJSON(w, http.StatusOK, getChirps)
}

func (cfg *ApiConfig) GetChirp(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirpID")
	parsedUUID, parseErr := uuid.Parse(chirpId)
	if parseErr != nil {
		RespondWithError(w, http.StatusBadRequest, "Not a valid UUID")
		return
	}

	getChirp, getErr := cfg.Database.GetChirp(r.Context(), parsedUUID)
	if getErr != nil {
		log.Println("Error fetching chirp:", getErr)
		RespondWithError(w, http.StatusNotFound, "Chirp not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, getChirp)
}

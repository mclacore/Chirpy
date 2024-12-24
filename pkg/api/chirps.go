package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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

	chirp.CleanedBody = ProfaneToAsterisks(chirp.Body)

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

func ProfaneToAsterisks(s string) string {
	var cleanWords []string
	words := strings.Split(s, " ")
	for _, word := range words {
		switch strings.ToLower(word) {
		case "kerfuffle", "sharbert", "fornax":
			word = "****"
		}
		cleanWords = append(cleanWords, word)
	}
	return strings.Join(cleanWords, " ")
}

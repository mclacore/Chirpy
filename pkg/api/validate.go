package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

var reqBody Chirp

func ValidateChirpy(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(reqBody.Body) > 140 {
		RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	RespondWithJSON(w, http.StatusOK, reqBody)
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

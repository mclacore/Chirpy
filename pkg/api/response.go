package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	res := Chirp{Error: msg}
	dat, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func profaneToAsterisks(s string) string {
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

package api

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	res := Chirp{Error: msg}
	dat, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload string) {
	res := Chirp{CleanedBody: ProfaneToAsterisks(payload)}
	dat, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

package api

import (
	"log"
	"net/http"
	"os"
)

func (cfg *ApiConfig) Reset(w http.ResponseWriter, r *http.Request) {
	platform := os.Getenv("PLATFORM")
	if platform != "dev" {
		RespondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	if err := cfg.Database.DeleteAllChirps(r.Context()); err != nil {
		log.Println("Could not delete all chirps:", err)
		RespondWithError(w, http.StatusInternalServerError, "Could not delete chirps")
		return
	}
	cfg.FileserverHits.Store(0)
	RespondWithJSON(w, http.StatusOK, nil)
}

package api

import (
	"net/http"
	"os"
)

func (cfg *ApiConfig) Reset(w http.ResponseWriter, r *http.Request) {
	platform := os.Getenv("PLATFORM")
	if platform != "dev" {
		RespondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	cfg.FileserverHits.Store(0)
}

package api

import (
	"sync/atomic"

	"github.com/mclacore/Chirpy/internal/database"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	Database       *database.Queries
}

type Chirp struct {
	Body        string `json:"body"`
	Error       string `json:"error"`
	CleanedBody string `json:"cleaned_body"`
}

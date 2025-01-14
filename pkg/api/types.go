package api

import (
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/mclacore/Chirpy/internal/database"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	Database       *database.Queries
}

type Chirp struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Body        string    `json:"body"`
	UserId      uuid.UUID `json:"user_id"`
	CleanedBody string    `json:"cleaned_body"`
	Error       string    `json:"error"`
}

type User struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"hashed_password"`
}

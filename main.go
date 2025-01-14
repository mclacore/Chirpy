package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mclacore/Chirpy/internal/database"
	"github.com/mclacore/Chirpy/pkg/api"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Errorf("Error opening SQL database: %w\n", err)
	}
	dbQueries := database.New(db)

	port := "8081"
	apiCfg := &api.ApiConfig{
		Database: dbQueries,
	}

	mux := http.NewServeMux()
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))

	mux.Handle("/app/", apiCfg.MiddlewareMetricsInc(handler))
	mux.Handle("/assets", http.FileServer(http.Dir(".")))

	mux.HandleFunc("GET /api/chirps", apiCfg.GetChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.GetChirp)
	mux.HandleFunc("POST /api/chirps", apiCfg.PostChirp)

	mux.HandleFunc("POST /api/users", apiCfg.CreateUser)
	mux.HandleFunc("GET /api/healthz", api.HealthZHeader)

	mux.HandleFunc("POST /admin/reset", apiCfg.Reset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.Hits)
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Fatal(server.ListenAndServe())
}

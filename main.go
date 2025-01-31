package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type data struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubUrl       string `json:"github_url"`
}

var userData = data{
	Email:           "elizabethogundepo7@gmail.com",
	CurrentDateTime: getCurrentDateTime(),
	GithubUrl:       "https://github.com/girlincyberspace/HngStage0",
}

func getCurrentDateTime() string {
	now := time.Now()
	isoTimeStamp := now.Format(time.RFC3339)

	return isoTimeStamp
}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(userData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/get", enableCORS(http.HandlerFunc(getData)))

	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

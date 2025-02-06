package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Api Struct
type Api struct {
	Email            string `json:"email"`
	Current_datetime string `json:"current_datetime"`
	Github_url       string `json:"github_url"`
}

var now = time.Now().UTC().Format(time.RFC3339)
var api_content Api

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(api_content)

}

func main() {
	// Init Mux Browser
	r := mux.NewRouter()
	api_content = Api{Email: "chukwuebukanwokike2007@gmail.com",
		Current_datetime: now,
		Github_url:       "https://github.com/ebukamee/hng12-stage0-api"}

	r.HandleFunc("/", get).Methods("GET")

	// RUN SERVER

	log.Fatal(http.ListenAndServe(":8000", r))
}

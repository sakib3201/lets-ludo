package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type NumberInput struct {
	Number int `json:"number"` // JSON field "number"
}

func hello(writter http.ResponseWriter, route *http.Request) {
	writter.Write([]byte("Hello World!"))
}
func getRandom(writter http.ResponseWriter, route *http.Request) {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min
	writter.Write([]byte(strconv.Itoa(randomInt)))
}

func receiveNumber(w http.ResponseWriter, r *http.Request) {
	var input NumberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received number: %d\n", input.Number)
	response := map[string]string{
		"message": "Number received successfully!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	route := chi.NewRouter()
	route.Use(middleware.Logger)
	route.Get("/hello", hello)
	route.Get("/getRandom", getRandom)
	route.Post("/receiveNumber", receiveNumber)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", route)
}

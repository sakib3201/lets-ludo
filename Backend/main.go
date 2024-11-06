package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Remove this function after experimenting is over
func receiveNumber(writer http.ResponseWriter, r *http.Request) {
	var input_number int
	if err := json.NewDecoder(r.Body).Decode(&input_number); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received number: %d\n", input_number)

	response := map[string]string{
		"message": "Number received successfully!",
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

type ludoPiece struct {
	player       int // player number (the player the piece belongs to)
	piece_number int // piece number from 1-4 will work as an id for pieces
	position     int // index in game board
}

type gameStateResponse struct {
	dice_value            int
	game_board            [56]int
	current_player        int
	current_player_pieces [4]ludoPiece
	total_players         int
	player_base           int // the current players base position (index) in game board
}

// gameState responds with a JSON object containing the game state.
func gameState(writer http.ResponseWriter, _ *http.Request) {
	response := map[string]string{
		"game_state": "in_progress",
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

// getDiceValue generates a random number between 1 and 6 and responds with a JSON object
// containing the dice value. It will return an internal server error if encoding the response
// fails.
func getDiceValue(w http.ResponseWriter, _ *http.Request) {
	minNum := 1
	maxNum := 6

	rand.NewSource(time.Now().UnixNano())

	randomNum := rand.Intn(maxNum-minNum+1) + minNum

	response := map[string]int{
		"dice_value": randomNum,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func movePiece(w http.ResponseWriter, r *http.Request) {
	var piece_number int
	if err := json.NewDecoder(r.Body).Decode(&piece_number); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// change game state based on the piece number, where piece number is the selected piece of current player to move
	// TODO: implement the move piece logic and return new game state
}

// main initializes the HTTP server with routing configurations using the chi router.
// It sets up the middleware for logging requests and defines two routes:
// - GET /getRandom: handled by the getRandom function
// - POST /receiveNumber: handled by the receiveNumber function
// - GET /game: handled by the gameState function
// The server listens on port 8080.
func main() {
	route := chi.NewRouter()
	route.Use(middleware.Logger)

	route.Get("/rollDice", getDiceValue)
	route.Get("/movePiece", movePiece)
	route.Post("/receiveNumber", receiveNumber)
	route.Get("/game", gameState)

	http.ListenAndServe(":8080", route)
}

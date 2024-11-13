package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
<<<<<<< HEAD
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
=======
=======
>>>>>>> 2901888bfb5854409f69461839897aab3d4e54ba
type LudoBoard struct {
	UserNumber          int       `json:"userNumber"`
	RandomNumber        int       `json:"randomNumber"`
	UsersPiecesPosition [4][4]int `json:"usersPiecesPosition"`
}

type RandomNumberResponse struct {
	RandomNumber int `json:"randomNumber"`
}

var globalResponse = LudoBoard{
	UserNumber:   0,
	RandomNumber: 0,
	UsersPiecesPosition: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
}

func getRandom(Context *gin.Context) {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min

	response := RandomNumberResponse{
		RandomNumber: randomInt,
	}
	globalResponse.RandomNumber = randomInt
	Context.JSON(http.StatusOK, response)
}

func receivePieceNumber(Context *gin.Context) {
	variable := Context.Param("variable")
	pieceNumber, err := strconv.Atoi(variable)
	if err != nil {
		fmt.Println("Error converting variable to int:", err)
		pieceNumber = 0 // Or handle it however is appropriate for your app
<<<<<<< HEAD
	}

	log.Printf("Received number: %d\n", pieceNumber)
	globalResponse.UserNumber += 1
	user := globalResponse.UserNumber - 1
	diceNumber := globalResponse.RandomNumber
	globalResponse.UsersPiecesPosition[user][pieceNumber] += diceNumber
	response := LudoBoard{
		UserNumber:          user,
		RandomNumber:        diceNumber,
		UsersPiecesPosition: globalResponse.UsersPiecesPosition,
	}
	Context.JSON(http.StatusOK, response)
>>>>>>> 2901888bfb5854409f69461839897aab3d4e54ba
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
=======
	}

	log.Printf("Received number: %d\n", pieceNumber)
	globalResponse.UserNumber += 1
	user := globalResponse.UserNumber - 1
	diceNumber := globalResponse.RandomNumber
	globalResponse.UsersPiecesPosition[user][pieceNumber] += diceNumber
	response := LudoBoard{
		UserNumber:          user,
		RandomNumber:        diceNumber,
		UsersPiecesPosition: globalResponse.UsersPiecesPosition,
	}
	Context.JSON(http.StatusOK, response)
>>>>>>> 2901888bfb5854409f69461839897aab3d4e54ba
}

// main initializes the HTTP server with routing configurations using the chi router.
// It sets up the middleware for logging requests and defines two routes:
// - GET /getRandom: handled by the getRandom function
// - POST /receiveNumber: handled by the receiveNumber function
// - GET /game: handled by the gameState function
// The server listens on port 8080.
func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	route := chi.NewRouter()
	route.Use(middleware.Logger)

	route.Get("/rollDice", getDiceValue)
	route.Get("/movePiece", movePiece)
	route.Post("/receiveNumber", receiveNumber)
	route.Get("/game", gameState)

	http.ListenAndServe(":8080", route)
=======
=======
>>>>>>> 2901888bfb5854409f69461839897aab3d4e54ba
	route := gin.Default()
	route.Use(gin.Logger())

	route.GET("/getRandom", getRandom)

	route.POST("/receivePieceNumber/:variable", receivePieceNumber)

	log.Println("Starting server on :8080")
	route.Run(":8080")
<<<<<<< HEAD
>>>>>>> 2901888bfb5854409f69461839897aab3d4e54ba
=======
>>>>>>> 2901888bfb5854409f69461839897aab3d4e54ba
}

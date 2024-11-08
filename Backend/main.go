package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
}

func main() {
	route := gin.Default()
	route.Use(gin.Logger())

	route.GET("/getRandom", getRandom)

	route.POST("/receivePieceNumber/:variable", receivePieceNumber)

	log.Println("Starting server on :8080")
	route.Run(":8080")
}

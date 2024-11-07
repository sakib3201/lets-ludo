package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
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
	pieceNumber := variable
	if err := Context.ShouldBindJSON(&pieceNumber); err != nil {
		Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received number: %d\n", pieceNumber)
	user = globalResponse.UserNumber
	diceNumber = globalResponse.RandomNumber
	PiecePosition[4][4] := globalResponse.UsersPiecesPosition
	piecePosition[user][variable] += diceNumber
	response := LudoBoard {
		UserNumber: (user+1)%4,
		RandomNumber: diceNumber,
		UsersPiecesPosition: ,
	}
	response := map[string]string{
		"message": "Number received successfully!",
	}
	Context.JSON(http.StatusOK, response)

}

func main() {
	route := gin.Default()
	route.Use(gin.Logger())

	route.GET("/getRandom", getRandom)

	route.POST("/receivePieceNumber/:variable", receivePieceNumber) // কত নাম্বার গুটি চালবি

	log.Println("Starting server on :8080")
	route.Run(":8080")
}

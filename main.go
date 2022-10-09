package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type water struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

type water_request struct {
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Type   string `json:"type"`
}

var water_menu = []water{
	{Id: 0, Name: "Giga Chad Water", Origin: "Fiji", Type: "mineral", Price: 7, Amount: 10},
	{Id: 1, Name: "Big Chungus Drink", Origin: "Japan", Type: "sparkling", Price: 20, Amount: 5},
}

var water_menu_requests = []water_request{
	{Name: "Epic Water", Origin: "Canada", Type: "rain"},
}

func getWaterMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, water_menu)
}

func waterById(id int) (*water, error) {
	for i, w := range water_menu {
		if w.Id == id {
			return &water_menu[i], nil
		}
	}

	return nil, errors.New("water not found")
}

func getWaterById(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad query value"})
		return
	}

	foundWater, err := waterById(intId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Water not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, foundWater)
}

func waterByIdWithRequestCheck(c *gin.Context) (*water, error) {
	id, ok := c.GetQuery("id")
	operationFailed := errors.New("operation failed")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing query"})
		return nil, operationFailed
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad query value"})
		return nil, operationFailed
	}

	foundWater, err := waterById(intId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Water not found"})
		return nil, operationFailed
	}

	return foundWater, nil
}

func patchBuyWater(c *gin.Context) {
	foundWater, err := waterByIdWithRequestCheck(c)

	if err != nil {
		return
	}

	if foundWater.Amount <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This water is out of stock"})
		return
	}

	foundWater.Amount--
	c.IndentedJSON(http.StatusOK, foundWater)
}

func patchGiveWater(c *gin.Context) {
	foundWater, err := waterByIdWithRequestCheck(c)

	if err != nil {
		return
	}

	if foundWater.Amount >= 10 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Magazine is full"})
		return
	}

	foundWater.Amount++
	c.IndentedJSON(http.StatusOK, foundWater)
}

func main() {
	router := gin.Default()
	router.GET("/menu", getWaterMenu)

	router.GET("/water/:id", getWaterById)

	router.PATCH("/buy", patchBuyWater)
	router.PATCH("/give", patchGiveWater)

	router.Run("localhost:8080")
}

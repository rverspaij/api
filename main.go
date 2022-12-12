package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reservation struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// LastName	string		`json:"lastname"`
	// Time		int			`json:"time"`
}

var reservations = []reservation{
	{Username: "loloke1", Password: "Admin01"},
}

func addReservation(c *gin.Context) {
	var newReservation reservation

	version := c.Param("version")
	fmt.Println("Version", version)
	if version == "v2" {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	} else {
		c.IndentedJSON(http.StatusBadRequest, newReservation)
		log.Fatal("Not the right key")
	}

	if err := c.BindJSON(&newReservation); err != nil {

		return
	}

	reservations = append(reservations, newReservation)
	fmt.Println(reservations)
	c.IndentedJSON(http.StatusCreated, newReservation)
}

func main() {
	router := gin.Default()
	router.POST("/reservations/:version", addReservation)
	router.Run("localhost:8000")
}

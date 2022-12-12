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

	if err := c.BindJSON(&newReservation); err != nil {
		log.Fatal(err)
		return
	}
	reservations = append(reservations, newReservation)
	fmt.Println(reservations)
	c.IndentedJSON(http.StatusCreated, newReservation)
}

func main() {
	router := gin.Default()
	router.POST("/reservations", addReservation)
	router.Run("localhost:8080")
}

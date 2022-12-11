package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reservation struct {
	Username string `json:"inp_username"`
	Password string `json:"inp_password"`
	// LastName	string		`json:"lastname"`
	// Time		int			`json:"time"`
}

var reservations = []reservation{
	{Username: "1", Password: "Admin01"},
}

func addReservation(c *gin.Context) {
	var newReservation reservation

	version := c.Param("version")
	fmt.Println("Version", version)
	if version == "v2" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	}

	if err := c.BindJSON(&newReservation); err != nil {
		return
	}

	reservations = append(reservations, newReservation)
	c.IndentedJSON(http.StatusCreated, newReservation)
	fmt.Println(newReservation)
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/reservation", addReservation)
	router.Run("localhost:8080")
}

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type reservation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var reservations = []reservation{
	{Username: "loloke1", Password: "Admin01"},
}

func main() {
	router := gin.Default()
	router.POST("/reservations/:version", addReservation)
	router.Run("localhost:8000")
}

// Function to post data to API.
func addReservation(c *gin.Context) {
	var newReservation reservation

	version := c.Param("version")

	data, err := os.ReadFile("./api.key") // Checking for an API key in a .key file
	if err != nil {
		log.Fatal("Could not read the Key File: ", err)
	}
	apikey := string(data)

	fmt.Println("Version:", version)
	if version == apikey {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		if err := c.BindJSON(&newReservation); err != nil {
			errorHandler(err, nil)
		}
		reservations = append(reservations, newReservation)
		c.IndentedJSON(http.StatusCreated, newReservation)
		fmt.Println(newReservation)
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, newReservation)
		err := errors.New("unauthorized user tried to connect")
		errorHandler(nil, err)
		return
	}
}

// Handles errors to log file.
func errorHandler(err error, warning error) {
	file, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file.Close()

	warner := log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	if warning != nil {
		warner.Println("Unauthorized user tried to connect!")
	}

	logger := log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
	if err != nil {
		logger.Fatal(err)
	}
}

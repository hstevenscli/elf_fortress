package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Initialize a player object
	p := Player{ Name: "Elfo", Health: 100 }
	fmt.Println(p)

	w, h := 10, 10


	level1 := initZone(w, h)

	appendPlayerToZone(5, 8, &p, &level1)
	printZone(&level1)
	printZoneAll(&level1)

	////////// SERVER ROUTES HERE /////////////
	// Handlers will be in handlers.go

	// Can be like this with a function passed directly as the second parameter
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "pong"})
	})

	// Or just the name of a function
	router.GET("/test", testHandler)
	router.Run()
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Initialize a player object
	p := Player{ Name: "Elfo", Health: 100 }



	// INITIALIZE various zone objects
	// starterZone := Zone{
	// 	Name: "Starter Zone",
	// 	Zone: [][]Tile{ // Outer slice for rows
	// 		{ // First row
	// 			{GroundType: 0, Other: 1, CurrentPlayers: []*Player{ &p }},
	// 		},
	// 	},
	// }

	var sz SimpleZone = SimpleZone{ Name: "Hal", Land: [][]int{
		{
			0, 0, 1, 2,
		},
		{
			5, 5, 5, 25,
		},
	}}

	fmt.Println(p)
	fmt.Println(sz)
	// fmt.Println(starterZone)
	// fmt.Println(starterZone.Zone[0][0].CurrentPlayers[0])




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

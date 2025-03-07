package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello000w world")
	fmt.Println("hellrld")

	p := Player{ Name: "Elfo", Health: 100 }
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


}

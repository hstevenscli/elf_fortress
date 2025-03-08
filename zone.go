package main

import (
	"fmt"

)

type Tile struct {
	GroundType int
	Other int
	CurrentPlayers []*Player
}

type Zone struct {
	Name string
	Zone [][]Tile
}

type SimpleZone struct {
	Name string
	Land [][]int
}


// Put a player into a specific location in the zone
func appendPlayerToZone(row int, col int, p *Player, zp *Zone) {
	tile := &zp.Zone[row][col]
	tile.CurrentPlayers = append(tile.CurrentPlayers, p)
	tile.GroundType = 1
}

// Create and return a Zone filled with zeros of the given size
func initZone(width int, height int) Zone {
	z := Zone{ Name: "Level 1", Zone: make([][]Tile, height)}
	for i:= range z.Zone {
		t := Tile{GroundType: 0, Other: 0, CurrentPlayers: make([]*Player, 0)}
		z.Zone[i] = make([]Tile, width)
		for j := range z.Zone[i] {
			z.Zone[i][j] = t
		}
	}
	return z
}

func printZone(zp *Zone) {
	for i := range zp.Zone {
		for j := range zp.Zone[i] {
			fmt.Printf("%v", zp.Zone[i][j].GroundType)
		}
		fmt.Printf("\n")
	}
}

func printZoneAll(zp *Zone) {
	for i := range zp.Zone {
		for j := range zp.Zone[i] {
			fmt.Printf("%v", zp.Zone[i][j])
		}
		fmt.Printf("\n")
	}
}

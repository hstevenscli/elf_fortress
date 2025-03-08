package main

import (

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

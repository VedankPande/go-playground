package main

import "fmt"
//House struct
type House struct {
	wallType string
	doorType string
	rooms    int
}

// String representation for a house
func (h *House) String() (string){
	return fmt.Sprintf("A house with %v walls, a %v, and %v rooms", h.wallType, h.doorType, h.rooms)
}
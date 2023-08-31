package main

import (
	"fmt"
)

/* Builder pattern - */


// Builder interface
type iBuilder interface {
	setWallType()
	setDoorType()
	setRooms()
	getHouse() House
}

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

// Director struct that encapsulates a builder
type Director struct {
	builder iBuilder
}

func (d *Director) setBuilder(newBuilder iBuilder) {
	d.builder = newBuilder
}

// get house from builder
func (d *Director) buildHouse() House {
	return d.builder.getHouse()
}

func getDirector(b iBuilder) *Director {
	return &Director{builder: b}
}

// Builder 1
type SmallWoodenHouseBuilder struct {
	wallType string
	doorType string
	rooms    int
}

func (s *SmallWoodenHouseBuilder) setWallType() {
	s.wallType = "wooden"
}

func (s *SmallWoodenHouseBuilder) setDoorType() {
	s.doorType = "single door"
}

func (s *SmallWoodenHouseBuilder) setRooms() {
	s.rooms = 1
}

// Construct and return house
func (s *SmallWoodenHouseBuilder) getHouse() House {

	s.setWallType()
	s.setDoorType()
	s.setRooms()

	return House{
		wallType: s.wallType,
		doorType: s.doorType,
		rooms:    s.rooms,
	}
}

// Builder 2
type LargeConcreteHouseBuilder struct {
	wallType string
	doorType string
	rooms    int
}

func (l *LargeConcreteHouseBuilder) setWallType() {
	l.wallType = "concrete"
}

func (l *LargeConcreteHouseBuilder) setDoorType() {
	l.doorType = "double door"
}

func (l *LargeConcreteHouseBuilder) setRooms() {
	l.rooms = 3
}

func (l *LargeConcreteHouseBuilder) getHouse() House {

	l.setWallType()
	l.setRooms()
	l.setDoorType()

	return House{
		wallType: l.wallType,
		doorType: l.doorType,
		rooms:    l.rooms,
	}
}


func main() {
	
	//directors
	directorA := &Director{builder: &SmallWoodenHouseBuilder{}}
	directorB := &Director{builder: &LargeConcreteHouseBuilder{}}

	//house instances
	houseA := directorA.buildHouse()
	houseB := directorB.buildHouse()

	fmt.Println(&houseA)
	fmt.Println(&houseB)
}

package main

// Builder interface
type iBuilder interface {
	setWallType()
	setDoorType()
	setRooms()
	getHouse() House
}

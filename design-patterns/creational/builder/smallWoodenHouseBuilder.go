package main

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

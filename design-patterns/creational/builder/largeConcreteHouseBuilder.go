package main

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
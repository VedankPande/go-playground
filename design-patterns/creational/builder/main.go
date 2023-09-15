package main

import (
	"fmt"
)

/* Builder pattern - */

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

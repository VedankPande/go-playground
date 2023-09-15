package main

func main(){

	application := Application{}
	choice := "Adidas"

	switch choice {
	case "Nike":
		application.factory = NikeEquipmentFactory{}
	case "Adidas":
		application.factory = AdidasEquipmentFactory{}
	default:
		panic("Invalid input - choose Nike or Adidas")
	}

	application.createBasketball()
	application.createFootball()

	application.basketball.bounce()
	application.football.kick()
}
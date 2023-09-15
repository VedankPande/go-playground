package main

/* ------------- Application ----------- */
//The application relies only on abstractions (Dependency Inversion) - making it possible to pass any type of Factory in the future
type Application struct{
	factory Factory
	football Football
	basketball BasketBall
}

func (a *Application) getFactory() (Factory){
	return a.factory
}
func (a *Application) createFootball(){
	
	a.football = a.factory.createFootball()
}

func (a *Application) createBasketball(){

	a.basketball = a.factory.createBasketBall()
}


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
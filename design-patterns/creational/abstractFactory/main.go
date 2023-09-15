package main

import(
	"fmt"
)

/* ------------- Objects ----------- */
type Football interface{
	kick()
}

type NikeFootball struct{
	
}

func (n NikeFootball) kick(){
	fmt.Println("Kicked the Nike football")
}

type AdidasFootball struct{

}

func (a AdidasFootball) kick(){
	fmt.Println("Kicked the Adidas football")
}


type BasketBall interface{
	bounce()
}

type NikeBasketball struct{

}

func (n NikeBasketball) bounce(){
	fmt.Println("Bounced the Nike basketball")
}

type AdidasBasketball struct{
	
}

func (a AdidasBasketball) bounce(){
	fmt.Println("Bounced the Adidas basketball")
}


/* ------------- Factories ----------- */
type Factory interface{
	createFootball() Football
	createBasketBall() BasketBall
}

type NikeEquipmentFactory struct{

}

func (n NikeEquipmentFactory) createFootball() (Football){
	return NikeFootball{}
}

func (n NikeEquipmentFactory) createBasketBall() (BasketBall){
	return NikeBasketball{}
}

type AdidasEquipmentFactory struct{

}

func (a AdidasEquipmentFactory) createFootball() (Football){
	return AdidasFootball{}
}

func (a AdidasEquipmentFactory) createBasketBall() (BasketBall){
	return AdidasBasketball{}
}

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
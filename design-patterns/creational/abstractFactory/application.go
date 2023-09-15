package main

/* ------------- Application ----------- */
//The application relies only on abstractions (Dependency Inversion) - making it possible to pass any type of Factory in the future
type Application struct{
	factory Factory
	football Football
	basketball BasketBall
}

func (a *Application) createFootball(){
	
	a.football = a.factory.createFootball()
}

func (a *Application) createBasketball(){

	a.basketball = a.factory.createBasketBall()
}
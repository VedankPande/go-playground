package main

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
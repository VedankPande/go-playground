package main

type NikeEquipmentFactory struct{

}

func (n NikeEquipmentFactory) createFootball() (Football){
	return NikeFootball{}
}

func (n NikeEquipmentFactory) createBasketBall() (BasketBall){
	return NikeBasketball{}
}

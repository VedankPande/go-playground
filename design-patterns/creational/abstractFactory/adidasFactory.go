package main


type AdidasEquipmentFactory struct{

}

func (a AdidasEquipmentFactory) createFootball() (Football){
	return AdidasFootball{}
}

func (a AdidasEquipmentFactory) createBasketBall() (BasketBall){
	return AdidasBasketball{}
}
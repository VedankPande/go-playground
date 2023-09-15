package main

import "fmt"
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
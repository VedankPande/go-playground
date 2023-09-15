package main

type Factory interface{
	createFootball() Football
	createBasketBall() BasketBall
}
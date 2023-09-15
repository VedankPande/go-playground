package main

type Visitor interface{
	visitForDog()
	visitForCat()
	//add for other types 
}
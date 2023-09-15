package main

import "fmt"

type AnimalFeeder struct{

}

func (f AnimalFeeder) visitForCat(){
	fmt.Println("Feeding Cat")
}


func (f AnimalFeeder) visitForDog(){
	fmt.Println("Feeding Dog")
}

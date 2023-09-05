package main 

import(
	"fmt"
)

/* 
	The visitor pattern allows for the addition of new methods and behaviours in a class
	without having to modify the existing class (apart from the visitor pattern code).
*/

type Visitor interface{
	visitForDog()
	visitForCat()
	//add for other types 
}

type AnimalFeeder struct{

}

func (f AnimalFeeder) visitForCat(){
	fmt.Println("Feeding Cat")
}


func (f AnimalFeeder) visitForDog(){
	fmt.Println("Feeding Dog")
}

type Animal interface{
	Speak()
	Accept(Visitor)
}

type Dog struct{

}

type Cat struct{

}

func (c Cat) Speak(){
	fmt.Println("I'm a cat")
}

func (d Dog) Speak(){
	fmt.Println("I'm a dog")
}

func (c Cat) Accept(v Visitor){
	v.visitForCat()
}

func (d Dog) Accept(v Visitor){
	v.visitForDog()
}

func main(){
	dog := Dog{}
	cat := Cat{}

	dog.Speak()
	cat.Speak()

	feeder := AnimalFeeder{}

	dog.Accept(feeder)
	cat.Accept(feeder)
}
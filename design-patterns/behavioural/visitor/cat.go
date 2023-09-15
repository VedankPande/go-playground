package main

import "fmt"

type Cat struct{

}

func (c Cat) Speak(){
	fmt.Println("I'm a cat")
}

func (c Cat) Accept(v Visitor){
	v.visitForCat()
}
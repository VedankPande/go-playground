package main

import "fmt"

type Dog struct{

}

func (d Dog) Speak(){
	fmt.Println("I'm a dog")
}

func (d Dog) Accept(v Visitor){
	v.visitForDog()
}
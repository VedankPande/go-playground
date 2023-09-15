package main

type Animal interface{
	Speak()
	Accept(Visitor)
}
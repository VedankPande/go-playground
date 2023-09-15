package main 

/* 
	The visitor pattern allows for the addition of new methods and behaviours in a class
	without having to modify the existing class (apart from the visitor pattern code).
*/

func main(){
	dog := Dog{}
	cat := Cat{}

	dog.Speak()
	cat.Speak()

	feeder := AnimalFeeder{}

	dog.Accept(feeder)
	cat.Accept(feeder)
}
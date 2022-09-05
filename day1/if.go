package main

func main() {

	name := "Doe John"

	if name == "John" {
		println("Hello John")
	} else if name == "Doe" {
		println("Hello Doe")
	} else {
		println("Nice to meet you!")
	}

	if length := len(name); length == 3 {
		println("3 Character")
	} else {
		println("Greater 3")
	}

}

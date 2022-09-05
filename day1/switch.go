package main

func main() {
	fullname := "David"

	switch fullname {
	case "David":
		println("Hello David")
	case "BE":
		println("Hello BE")
	default:
		println("Nice to meet you!")
	}

	switch length := len(fullname); length == 5 {
	case true:
		println("The length is 5")
	case false:
		println("The length is NOT 5")
	}

	switch {
	case fullname == "David":
		println("This is David")
	case fullname == "BE":
		println("This is BE")
	default:
		println("Is not exist")
	}
}

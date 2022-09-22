package main

func getRandom() interface{} {
	return "Hello"
}

func getPanicMessage() {
	println("Don't panic!")
}

func main() {

	/*text := getRandom().(string)
	println(text)*/

	/*defer getPanicMessage() // always call if error
	number := getRandom().(int)
	println(number)*/

	text := getRandom()
	switch text.(type) {
	case string:
		println("Variable type string", text.(string))
	case int:
		println("Variable type integer", text.(int))
	case bool:
		println("Variable type boolean", text.(bool))
	default:
		println("Unknown")

	}
}

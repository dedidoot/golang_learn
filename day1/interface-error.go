package main

import "errors"

func getDiv(value1 int, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("value div cannot 0")
	}

	return value1 / value2, nil
}

func main() {

	/*exError := errors.New("ups error")
	println(exError.Error())*/

	div, err := getDiv(25, 0)

	println(div)
	println(err.Error())
}

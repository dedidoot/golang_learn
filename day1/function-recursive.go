package main

func factorialLoop(value int) int {
	result := 1
	for newVal := value; newVal > 0; newVal-- {
		result *= newVal
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	println(factorialLoop(5))
	println(factorialRecursive(5))

	println(5 * 4 * 3 * 2 * 1)
}

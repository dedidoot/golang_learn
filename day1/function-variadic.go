package main

func getSum(number ...int) int {
	sum := 0
	for _, num := range number {
		sum += num
	}
	/*for index := 0; index < len(number); index++ {
		sum += number[index]
	}*/
	return sum
}

func main() {
	test := getSum(10, 10, 10, 10)
	println(test)

	slice := []int{20, 20, 20, 20}
	println(getSum(slice...))
}

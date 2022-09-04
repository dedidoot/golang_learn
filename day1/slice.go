package main

import "fmt"

func main() {

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"Mei",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "June_edited"
	fmt.Println(slice1)

	slice1[0] = "Mei_edited"
	fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Added")
	fmt.Println(slice3)

	slice4 := months[10:11]
	fmt.Println(slice4)

	slice5 := append(slice4, "december_replaced")
	fmt.Println(slice5)
	fmt.Println()

	newSlice := make([]string, 5, 5)
	newSlice[0] = "John"
	newSlice[1] = "Doe"
	newSlice[2] = "Alex"
	newSlice[3] = "Sunday"
	newSlice[4] = "Monday"
	fmt.Println(newSlice)

	copyNewSlice := make([]string, len(newSlice), cap(newSlice))
	copy(copyNewSlice, newSlice)
	fmt.Println(copyNewSlice)

	thisIsArray := [5]int{1, 2, 3}
	fmt.Println(thisIsArray)

	thisIsSlice := []int{1, 2, 3}
	fmt.Println(thisIsSlice)
}

package main

func main() {

	counter := 0

	increment := func() {
		counter++
	}

	increment()
	println(counter)

}

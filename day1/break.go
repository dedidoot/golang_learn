package main

func main() {
	for index := 0; index < 10; index++ {
		if index == 4 {
			break
		}
		println("index", index)
	}
}

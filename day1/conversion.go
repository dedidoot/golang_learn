package main

func main() {
	var nilai32 int32 = 127
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	println(nilai32)
	println(nilai64)
	println(nilai8)

	realText := "Golang"
	indexTwo := realText[1]
	indexTwoString := string(indexTwo)

	println(realText)
	println(indexTwo)
	println(indexTwoString)

}

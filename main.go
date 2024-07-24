package main

import "fmt"

func main() {
	//classic for syntax
	var j int
	for j = 0; j < 10; j++ {
		fmt.Println(j)
	}

	//short classic for syntax
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for without init and increment/decrement
	sum := 1
	for sum < 20 {
		sum += sum
		fmt.Println(sum)
	}

	//for like a WHILE in other languages
	for sum <= 100 {
		sum += 1
		fmt.Println(sum)
	}

	// infinite cycle
	for {
		fmt.Println("Stop me please")
	}
}

package main

import "fmt"

func main() {
	var intPointer *int
	fmt.Printf("Int pointer is %v\n", intPointer)

	var a int = 9
	var aPointer *int = &a
	fmt.Printf("aPointer is %#v %#v \n", aPointer, *aPointer)

	var newPointer = new(int)
	fmt.Printf("newPointer is %#v %#v \n", newPointer, *newPointer)
	*newPointer = 10
	fmt.Printf("newPointer is %#v %#v \n", newPointer, *newPointer)

	num := 3
	square(num)
	fmt.Println(num)

	squarePointer(&num)
	fmt.Println(num)

	// empty value flag
	var wallet1 *int
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))
}

func square(num int) {
	num *= num
}

func squarePointer(num *int) {
	*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}

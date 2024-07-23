package main

import "fmt"

func main() {
	x, y := 2, 3

	greet()
	greetWithParameter("Vova")
	greetWithParameters("Vova", "Bilan")
	fmt.Println("Result of sum: ", sum(x, y))
	sum, mult := sumAndMultiply(x, y)
	fmt.Printf("Result of sum and multiply: %d and %d\n", sum, mult)
	_, mult32 := namedSumAndMultiply(x, y) // var _ - proxy
	fmt.Printf("Result of multiply: %d\n", mult32)
}

func greet() {
	fmt.Println("Wusup")
}

func greetWithParameter(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func greetWithParameters(name, surname string) {
	fmt.Printf("Hello, %s %s!\n", name, surname)
}

func sum(first int, second int) int {
	return first + second
}

func sumAndMultiply(first int, second int) (int, int) {
	return first + second, first * second
}

func namedSumAndMultiply(first int, second int) (sum int32, mult int32) {
	sum = int32(first + second)
	mult = int32(first * second)
	return // or return sum, mult
}

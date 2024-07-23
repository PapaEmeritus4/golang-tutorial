package main

import "fmt"

func main() {
	x, y := 5, 3
	xf, yf := 3.0, 2.0

	greet()
	greetWithParameter("Vova")
	greetWithParameters("Vova", "Bilan")
	fmt.Println("Result of sum: ", sum(x, y))
	sum, mult := sumAndMultiply(x, y)
	fmt.Printf("Result of sum and multiply: %d and %d\n", sum, mult)
	_, mult32 := namedSumAndMultiply(x, y) // var _ - proxy
	fmt.Printf("Result of multiply: %d\n", mult32)

	/*  func as var  */
	var multiplier func(x, y int) int
	multiplier = func(x, y int) int { return x * y }
	fmt.Printf("Result of func multiply as var: %d\n", multiplier(x, y))

	divider := func(x, y float32) float32 { return float32(x / y) }
	fmt.Printf("Result of func divide as var: %f\n", divider(float32(xf), float32(yf)))
	/*  func as var  */

	// func like parameter
	sumFunc := func(x, y int) int { return x + y }
	subFunc := func(x, y int) int { return x - y }
	fmt.Printf("Result of calculate: %d\n", calculate(x, y, sumFunc))
	fmt.Printf("Result of calculate: %d\n", calculate(x, y, subFunc))

	// Return new function
	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)
	fmt.Println(divideBy2(100))
	fmt.Println(divideBy10(100))

	//getter
	usd := 30
	getUsdValue := func() int { return usd }
	fmt.Println(getUsdValue())
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

// func like parameter
func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

// in func we return new func
func createDivider(divider int) func(x int) int {
	return func(x int) int {
		return x / divider
	}
}

/*
	function naming
	main() = private
	Main() = public
*/

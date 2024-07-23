package main

/*
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64
byte (uint8)
rune (int32)
float32, float64
complex64, complex128
*/

import (
	"fmt"
	"unsafe"
)

// standart syntax
var name string
var cond bool
var ourtCond = true

// const var
const age = 100

func main() {
	//local var
	var name string
	name = "localVar"
	fmt.Println(name)

	// short syntax (only in func)
	shortVar := 19

	fmt.Println(cond)
	cond = true
	fmt.Println(cond)

	//init
	name = "Vova"
	fmt.Printf("Type: %T Value: %v\n", name, name)

	fmt.Println(shortVar)
	fmt.Println(ourtCond)

	//cast
	var i1 int64 = 15
	var i2 int8 = 15
	fmt.Println(int8(i1) + i2)

	fmt.Println(unsafe.Sizeof(int8(1)))  //1 byte
	fmt.Println(unsafe.Sizeof(int16(1))) //2 byte
	fmt.Println(unsafe.Sizeof(int32(1))) //4 byte
	fmt.Println(unsafe.Sizeof(int64(1))) //8 byte
}

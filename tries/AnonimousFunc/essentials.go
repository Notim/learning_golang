package main

import (
	"reflect"
)
import 	"fmt"

func main() {
	// 1 -> Auto calling Anonimous function
	(func(name string){
		fmt.Println("With Auto calling ==> Hello" + name)
	})("Notim_")

	// 2 -> variable receive the anonimous func (but needs to be called after)
	received := func(name string){
		fmt.Println("With Explicit calling ==> Hello" + name)
	}
	received("Notim_") // explicit calling

	// the 1st and 2nd means the same thing !!!!
	fmt.Println(reflect.TypeOf(received))

	// look this shit KKKKKKKKKKKKKKKKK
	(func (A string) {
		(func (B string) {
			(func (C string){
				fmt.Println(C) // A + B + C => "> hello World, How Are You?"

			})(B + "How Are You?") // params
		})(A + "World, ") // params
	})("> hello ") // params

	// (func() to be executed)(params to send)
	(func(a string) {fmt.Println()})("> I m fine, cuz i have GoLang, thank you!!")

	// Oh My Fucking God look, this is so beautiful
	(func(in func(val string) string) {
		fmt.Println(in("Notim_"))
	})(func(val string) (value string) {
		value = "Hello " + val

		return
	})

	// Closure
	Add := (func() (Add func() int) { // this func returns another function
		index := 0 // isolated variable (Fake Private)

		Add = func() int { // Add receive a new instance of function
			index += 1

			return index
		}

		return // in go dont nedd especify the var to return if you already atributted him
	})()// void anonymous func

	fmt.Println(reflect.TypeOf(Add))

	/**
	 * the curious of this closures is the index value
	 * He changed always when the Add() is called
	 * like a property Obj from Poo
	 * remember, Go isn't OOP lang (100%)
	**/
	fmt.Println(Add()) // index(private) = 1
	fmt.Println(Add()) // index(private) = 2
	fmt.Println(Add()) // index(private) = 3
	fmt.Println(Add()) // index(private) = 4

	/**
	 * literals anon functions in go/js-similar lambda expr
	**/
	// Js similar: () => "some String Value"
	NoArgsRetStr :=  func() string {return "some String Value"} // no args return string
	fmt.Println(NoArgsRetStr())

	// Js similar: (arg) => "some String Value" + arg
	WithArgsRetStr := func(arg string) string {return "some String" + arg} // with string args and return string
	fmt.Println(WithArgsRetStr("Value"))

	// Js similar: () => {console.log("Some String Value")}
	NoArgNoRet := func() {fmt.Println("Some String Value")} // no args no return (void)
	NoArgNoRet()

	// Js similar: (arg) => {console.log("Some String Value" + arg)}
	WithArgNoRet := func(arg string) {fmt.Println("Some String " + arg)}
	WithArgNoRet ("Value")
}
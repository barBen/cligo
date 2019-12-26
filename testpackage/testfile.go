package testpackage

import (
	"fmt"
)

// A1 is an exported const
const A1 = 4

// a1 is an unexported const
const a1 = 4

// A is an exported function called A
// Test comment
// @CLIGO_COMMAND("a description for the remove sub-command")
// @CLIGO_OPTION(option1: "a description for option1")
// @CLIGO_OPTION(option2: "a description for option2")
// @CLIGO_ARGUMENT(arg1: "a description for arg1")
// @CLIGO_ARGUMENT(arg2: "a description for arg2")
func Remove(option1, option2 bool, arg1, arg2 int) {
	fmt.Println("executing function Remove()")
}

// @CLIGO_COMMAND()
// a is an unexported function called a
func Install() {
	fmt.Println("executing function Install()")
}

// @CLIGO_COMMAND()
// @CLIGO_ARGUMENT(a: min)
// @CLIGO_ARGUMENT(b: max)
// B is an exported function called B
func Build(a, b int, verbose bool) {
	fmt.Printf("executing function Build() with arguments - a = %d | b = %d | verbose = %v", a, b, verbose)
}

// c is an unexported function called c
func c() {
	fmt.Println("executing function c()")
}

// Func is an exported function called Func
func Func() {
	fmt.Println("executing function Func()")
}

// fn is an unexported function called fn
func fn() {
	fmt.Println("executing function fn()")
}

/*
Z is an exported function called Z
@CLIGO_COMMAND()
*/
func Chumis() {
	fmt.Println("executing function Chumis()")
}

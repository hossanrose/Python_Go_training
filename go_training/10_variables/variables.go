package main

import "fmt"

// package level variable declaration
var yo, java bool
var a, b int = 1, 2

func main() {
// function level variable declaration
	var i int
// if intializer is present the type can be omitted
	var c = "test"
// An implicit declaration can be done with ":=" construct , it is only available inside the dunction
	d := false
	fmt.Println(i, yo, java, a, b, c, d)
}

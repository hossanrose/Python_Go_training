package main

import "fmt"

// if both variables use same type you need to use the type only at the end
func add(x,y int) int {
	return x+y
}

func main() {
	fmt.Println(add(42,343))
}

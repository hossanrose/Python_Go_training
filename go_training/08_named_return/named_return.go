package main

import "fmt"

// return values can be named 
func mult(test int) (x int,y string) {
	x =  test * 5 
 	y = "testing"
	// return statment without arguments returns the named return values
	return 
}

func main(){
	fmt.Println(mult(4))
}

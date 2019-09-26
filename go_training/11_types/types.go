package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe bool = false
	Maxint  uint64 = 1<<64-1
	z 	complex128 =cmplx.Sqrt(-5+12i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", Maxint, Maxint)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

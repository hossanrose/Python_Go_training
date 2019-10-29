package main

import (
  "fmt"
  "math"
)

type vertex struct {
  X,Y float64
}

// Method is a function with reciver argument "(v vertex) is reciver"
func (v vertex) Abs() float64{
  return math.Sqrt((v.X*v.X)+(v.Y*v.Y))
}
// Same function declared in normal way
func Abs(v vertex) float64{
  return math.Sqrt((v.X*v.X)+(v.Y*v.Y))
}
func main(){
  v:=vertex{3,4}
  // Abs function is called in a object oriented syntax here
  fmt.Println(v.Abs())
  fmt.Println(Abs(v))
}

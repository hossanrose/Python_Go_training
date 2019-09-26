package main

import (
  "fmt"
  "math"
)

func main() {
  var x,y int = 2,3
  //  Type(variable) can be used to convert the Type
  var f float64 = math.Sqrt(float64(x*y + y*y))
  var z uint = uint(f)
  fmt.Print(x,y,f,z)
}

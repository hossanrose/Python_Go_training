// Reading : https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11

package main

import (
  "fmt"
  "math"
)

func compute(fn func(float64,float64) float64) float64{
  // parameters 3 and 4 will be passed to any function that is passed to compute
  // So in effect hypot will be passed to compute and compute will call hypot with fn(3,4)
  return fn(3,4)
}

func main (){
  hypot:= func(x,y float64) float64{
    return math.Sqrt(x+y)
  }
  fmt.Println(hypot(2,3))
  // Passing hypot function to the compute function
  fmt.Println(compute(hypot))
}

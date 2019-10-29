// Further reading

package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X,Y float64
}

func (v Vertex) Abs() float64{
  return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

// by passing a pointer as a receiver we will able to modify the value
// to which receiver points. Remove * from below line to check
//pointer receiver
func (v *Vertex) Scale(f float64){
  v.X = v.X*f
  v.Y = v.Y*f
}

func main(){
  v := Vertex{3,4}
  v.Scale(100)
  fmt.Println(v.Abs())
}

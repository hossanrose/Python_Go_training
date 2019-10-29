package main

import (
  "fmt"
  "math"
)

// defining an interface
type Shape interface{
  area() float64
}

// defining a type
type Circle struct{
  x,y,radius float64
}
// defining another type
type Rec struct{
  len,height float64
}

// define a method for circle
func (circle Circle) area() float64{
  return math.Pi* circle.radius*circle.radius
}

// define a method for rectangle
func (rec Rec) area() float64{
  return rec.len *rec.height
}

func getArea(shape Shape) float64{
  return shape.area()
}

func main(){
  circle:=Circle{x:0,y:0,radius:5}
  rectangle:=Rec{len:10,height:20}
  // Using methods
  fmt.Println(circle.area())
  fmt.Println(rectangle.area())
  // Using interface
  fmt.Println(getArea(circle))
}

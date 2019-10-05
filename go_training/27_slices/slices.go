// Slice type is a abstraction build on top of go Array

package main

import "fmt"

func main(){
  p:=[6]int{2,3,5,7,11,13}
  var s []int = p[1:4] // Creating slice from an array
  fmt.Println(s)
  names:= [4]string{
    "John",
    "Paul",
    "Holo",
    "hobo",
  }
  fmt.Println(names)
  a:=names[0:2] // Slice
  b:=names[1:3] //Slice
  fmt.Println(a,b)
  b[0]="XXX" //Changing the elements of a slice modifies the corresponding elements of its underlying array.
  fmt.Println(names)
  fmt.Println(b)
}

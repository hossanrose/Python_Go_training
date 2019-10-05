package main

import "fmt"

func main(){
  // make is used to create arrays that is dynamically sized
  // with make making a slice of length 5
  a:= make([]int, 5)
  printSlice(a)
  // with make making a slice of capacity 5 and length 0
  b:= make([]int,0, 5)
  printSlice(b)
}

func printSlice (s[]int){
  fmt.Printf("cap=%d,len=%d %v\n",cap(s),len(s),s)
}

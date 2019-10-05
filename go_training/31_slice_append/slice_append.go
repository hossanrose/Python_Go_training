package main

import "fmt"

func main(){
  var s[]int
  printSlice(s)
  s=append(s,5)
  printSlice(s)
  s=append(s,734,3,6,2)
  printSlice(s)
}

func printSlice(s[]int){
  fmt.Printf("cap=%d,len=%d,%v\n",cap(s),len(s),s)
}

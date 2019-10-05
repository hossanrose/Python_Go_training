package main

import "fmt"

func main(){
  s:=[]int{1,2,3,4,5,6,7}
  s=s[:3]
  // slice has both capacity and length
  // capacity is the length of the array
  // length is the number of element slice contains
  fmt.Println(s)
  fmt.Printf("len=%d,cap=%d",len(s),cap(s))
}

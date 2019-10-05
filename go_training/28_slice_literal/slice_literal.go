package main

import "fmt"

func main(){
  // creating a int slice literal
  q:=[]int{2,3,4,5,7,2}
  fmt.Println(q)
  // creating a struct slice literal
  s:=[] struct{
    i int
    j string
  }{
    {2,"test"},
    {3,"yo"},
    {4,"ute"},
  }
  fmt.Println(s)
// slice has the low bound as 0 and high bound as number of elements as default
 q=q[1:]
 fmt.Println(q)
 q=q[:3]
 fmt.Println(q)
}

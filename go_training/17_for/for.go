package main

import "fmt"

func main(){
  sum := 0
  for i :=0; i <10; i++{
    sum +=i
  }
  // int and post statments are optional
  for ; sum <1000;{
    sum +=sum
  }
  for sum <10000 {
    sum +=sum
  }
  fmt.Println(sum)
}

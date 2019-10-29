//https://www.calhoun.io/what-is-a-closure/
package main

import (
  "fmt"
)

func adder() func (int)int  {
  sum:=0
  // closure function has access to the value that is outside the func
  return func(x int) int {
    fmt.Println(x)
    sum +=x
    return sum
  }
}

func main(){
  pos:=adder();
  for i :=0; i <10; i++{
    fmt.Println(pos(i))
  }

}

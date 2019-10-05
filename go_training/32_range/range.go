package main

import "fmt"

var pow = []int{100,332,34,54,66,227}
func main(){
  // get both index and value
  for i,v := range pow {
    fmt.Printf("index=%d,value=%d\n",i,v)
  }
  // get only index
  for i := range pow {
    fmt.Printf("index=%d\n",i)
  }
  // get only vaiue
  for _,v :=range pow {
    fmt.Printf("value=%d\n",v)
  }
}

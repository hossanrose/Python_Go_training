package main

import "fmt"

func main() {
  i,j := 42,198
  p :=&i // pointing to i
  fmt.Println(*p) // printing the value of i
  *p = 25 // setting value of i through pointer
  fmt.Println(i)

  p = &j
  *p=*p+27 // setting the value of j with *package
  fmt.Println(j)
}

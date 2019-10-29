package main

import (
  "fmt"
)

func fibino (n int, c chan int){
  x,y := 0,1
  for i:=0; i<n;i++ {
    c <- x
    x,y = y, x+y
  }
  close(c)
}

func main(){
  c:=make(chan int,10)
  go fibino(cap(c),c)
  //range can be used to check if the channel is closed by sender
  for i:= range c{
    fmt.Println(i)
  }
}

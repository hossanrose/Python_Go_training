package main

import "fmt"

func add(x,y int) <-chan int {
  c := make(chan int)
  sum:=x
  go func(){
  for i:=x ; i<=y;i++{
    sum+=i
  }
  c<-sum
}()
  return c
}

func main(){
  c1 :=add(1,100)
  c2 :=add(101,200)
  c3 :=add(201,300)

  msg1:=<-c1
  msg2:=<-c2
  msg3:=<-c3
  fmt.Println(msg1+msg2+msg3)
}

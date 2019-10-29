package main

import "fmt"

func sum(s []int, c chan int) {
  sum := 0
  for _,v := range s {
    sum +=v
  }
  //send sum to channel c
  c <- sum
}

func main(){
  s := []int{7,2,4,6,23,6}

  c:= make(chan int)
  go sum(s[1:5],c)
  go sum(s[:len(s)/2],c)
  go sum(s[:len(s)],c)
  x,y,z := <-c,<-c,<-c

  fmt.Println(x,y,z)
}

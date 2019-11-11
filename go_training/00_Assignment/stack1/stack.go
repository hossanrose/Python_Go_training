package main

import (
  "fmt"
  "math/rand"
  "time"
  "sync"
)

type Array struct{
     x []int
  }
var wg sync.WaitGroup

type Arrayin interface{
  push(int) Array
  pop() Array
}

func (array Array) push(element int) Array{
  array.x = append(array.x, element )
  return array
}

func (array Array) pop() Array{
  fmt.Println(time.Now().Nanosecond(),"Popped Element :",array.x[len(array.x)-1])
  array.x = array.x[:len(array.x)-1]
  return array
}

func stack_push(element Arrayin, x int) Array{
  return element.push(x)
}
func stack_pop(element Arrayin) Array{
  return element.pop()
}
func random_num() int{
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(1000)
}

func producer(ch chan Array,id int) {
  defer wg.Done()
  rand:=random_num()
  select {
  case test:= <- ch:
    arr:=stack_push(test,rand)
    fmt.Println(time.Now().Nanosecond(),"PUSH--",id,"--After:",arr )
    ch <- arr
  default:
    var temp Array
    arr:=stack_push(temp,rand)
    fmt.Println(time.Now().Nanosecond(),"PUSH--",id,"--After:", arr )
    ch <- arr
  }
}

func consumer(ch chan Array,id int) {
  defer wg.Done()
  select {
  case temp:=<-ch:
  //fmt.Println(time.Now().Nanosecond(),"POP--",id,"--Before:", temp)
  if len(temp.x) >0{
    arr:=stack_pop(temp)
    fmt.Println(time.Now().Nanosecond(),"POP--",id,"--After:", arr)
    if len(arr.x) >0{
      ch <- arr
    }
  }else{
    fmt.Println("Nothing to pop")
  }
 }
}
func main(){
//NON parallel
  fmt.Println("========Begin non Parallel array stack========")
  var arr= Array{}
  //fmt.Println(arr)
  arr=stack_push(arr,29)
  fmt.Println(arr)
  arr=stack_push(arr,24)
  fmt.Println(arr)
  arr=stack_push(arr,234)
  fmt.Println(arr)
  arr=stack_pop(arr)
  fmt.Println(arr)
  fmt.Println("========End non Parallel array stack========")


// Parallel
  fmt.Println("========Begin Parallel array stack========")
  ch :=make(chan Array)
  num :=500
  for i:=0;i<num;i++{
    wg.Add(2)
    go producer(ch,i)
    //time.Sleep(1 * time.Second)
    go consumer(ch,i)
  }
  wg.Wait()
  fmt.Println("========End Parallel array stack========")
}

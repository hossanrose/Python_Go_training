package main

import (
  "fmt"
  "math/rand"
  "time"
)

type Array []int

type Elementin interface{
  push(int) Array
}

type Elementout interface{
  pop() Array
}

func (array Array) push(element int) Array{
  array = append(array, element )
  return array
}

func (array Array) pop() Array{
  if len(array) > 0{
    fmt.Println(time.Now().Nanosecond(),"Popped Element :",array[len(array)-1])
    array = array[:len(array)-1]
    //fmt.Println(array)

    return array
  }else{
    return array
  }

}

func stack_push(element Elementin, x int) Array{
  return element.push(x)
}
func stack_pop(element Elementout) Array{
  return element.pop()
}
func random_num() int{
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(1000)
}

func producer(ch chan Array,id int) {

  rand:=random_num()
  //fmt.Println( time.Now().Nanosecond(), "Push--",id,"--into stack:",rand )
  select {
  case test:= <- ch:
    arr:=stack_push(test,rand)
    fmt.Println(time.Now().Nanosecond(),"PUSH--",id,"--After:",arr )
    ch <- arr
  default:
    var temp Array
    arr:=stack_push(temp,rand)
    fmt.Println(time.Now().Nanosecond(),"PUSH--",id,"--Afewrwetter:", arr )
    ch <- arr
  }
}

func consumer(ch chan Array,id int) {
    temp:=<-ch
    //fmt.Println(time.Now().Nanosecond(),"POP--",id,"--Before:", temp)
    arr:=stack_pop(temp)
    fmt.Println(time.Now().Nanosecond(),"POP--",id,"--After:", arr)
    ch <- arr
}

func main(){
//NON parallel
  fmt.Println("========Begin non Parallel array stack========")
  var arr Array
  fmt.Println(arr)
  arr=stack_push(arr,29)
  fmt.Println(arr)
  arr=stack_push(arr,24)
  fmt.Println(arr)
  arr=stack_push(arr,234)
  fmt.Println(arr)
  arr=stack_pop(arr)
  fmt.Println(arr)
  fmt.Println("========Begin non Parallel array stack========")

// Parallel
  ch :=make(chan Array)
  num :=50

  for i:=0;i<num;i++{
    go producer(ch,i)

    //time.Sleep(1*time.Second)

     go consumer(ch,i)
}
}

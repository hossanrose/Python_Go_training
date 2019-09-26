package main
import "fmt"

func main(){
  // defer will run the deferred function only after the surrounding functions
  defer fmt.Print("world")
  fmt.Print("Hello ")

  fmt.Println("counting")
  //Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
  for i:=0;i<10;i++{
    defer fmt.Println(i)
  }
  fmt.Println("done")
}

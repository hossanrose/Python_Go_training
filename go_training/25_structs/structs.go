package main

import "fmt"

// Struct is used to group together multile vars and create a new type
type Vertex struct{
  X int
  Y int
  Z string
}

func main(){
  fmt.Println(Vertex{1,2,"test"})
  v := Vertex{4,5,"lol"}
  fmt.Println(v)
  v.X=8 // Struct values can be individualy accessed by dot operator
  fmt.Println(v)
  p:=&v // Struct values also can be accessed through a pointer
  p.Y=1e9 // Go alows to use pointer as p.Y instead of (*p).Y
  fmt.Println(v)
}

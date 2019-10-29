package main

import "fmt"

type Vertex struct {
  Lat,Long float64
}

var m = map[string]Vertex{
  "Bell Labs" : Vertex{
    343.23434,
    32423.232,
  },
  // You can omit the type also here
  "Google" : {
    223.122,
    542.232,
  },
}
func main(){
  fmt.Println(m)
}

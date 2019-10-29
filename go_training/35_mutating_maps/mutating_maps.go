package main

import "fmt"

func main(){
  m:= make(map[string]int)

  m["Input1"]=324
  fmt.Println("The value",m["Input1"])
  // delete
  delete(m,"Input1")
  fmt.Println("The value",m["Input1"])
  // he first value (i) is assigned the value stored under the key "route". If that key doesn't exist, i is the value type's zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.
  i,ok:=m["Input1"]
  fmt.Println(i,ok)
}

package main
import (
	"fmt"
  "sync"
)
func worker(ch_sum chan int,inFrom int, inTo int,  wg *sync.WaitGroup) {
    var sum = 0
        for i:=inFrom; i<=inTo; i++ { // assigning 1 to i
             sum += i // sum = sum + i
        }
        ch_sum <- sum
    wg.Done()
}
func main() {
//    var inVal int
//    fmt.Print("Enter number of numbers to generate: ")
//    _, err := fmt.Scanf("%d", &inVal)
//    if err != nil {
//            fmt.Println(err)
//     }
    inVal :=1500
    var bottom int = 0
    var top int
    var numOfroutine int = (inVal / 100) + 1
    fmt.Printf("I will generate %d \n ",numOfroutine)
    var jump = inVal % 100
    var wg sync.WaitGroup
    ch_sum := make(chan int, jump)
    for i := 1; i <= jump; i++ {
        wg.Add(1)
        top = (i*100) + jump
        go worker(ch_sum,bottom,top, &wg)
        bottom = top + 1
    }
     wg.Wait()
     total := 0
     for i :=0 ; i <jump;i++ {
       total += <-ch_sum
     }
   fmt.Println(total)
  }

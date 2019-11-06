package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"bufio"
    "os"
)

// lock -Read write
type Production struct { //define prodcution
	index           int    //production id index
	name_production string //name of production
	producer        string
}

func main() {
	ch := make(chan Production)
	n := 3
	for i := 1; i < n; i++ {
		go producer(ch, i)
		time.Sleep(1 * time.Second)
		go consumer(ch, i)
		time.Sleep(1 * time.Second)
	}

	//	go producer(ch, 1)
	//	time.Sleep(3 * time.Second)
	//	go consumer(ch, 1)
	//	time.Sleep(5 * time.Second)
	//	go producer(ch, 2)
	//	time.Sleep(2 * time.Second)
	//	go consumer(ch, 2)
	//	time.Sleep(6 * time.Second)

	// Just delay the end of main
 	reader := bufio.NewReader(os.Stdin)
    txt, _ := reader.ReadString('\n')
    fmt.Println(txt)
    fmt.Println("program complete")
}

func producer(c chan<- Production, id int) {
	var unit_production Production
	var name_product string
	for i := 0; i < 10; i++ {

		name_product = randUpString(1)
		producer_name := "producer"
		producer_name += strconv.Itoa(id)
		unit_production = Production{i, name_product, producer_name}
		fmt.Println("This is number", i, "--producer:", id, "--name", name_product)
		c <- unit_production
		time.Sleep(1 * time.Second)

		//for see the process of production and consume --use channel control
	}

}
func consumer(c <-chan Production, id int) {
	for i := 0; i < 10; i++ {

		buy := <-c
		fmt.Println("Consumer--", id, " --buy the number--", buy)
		time.Sleep(1 * time.Second)

	}
}

func randInt(min int, max int) byte {
	rand.Seed(time.Now().UnixNano())
	return byte(min + rand.Intn(max-min))
}
func randUpString(l int) string {
	var result bytes.Buffer
	var temp byte
	for i := 0; i < l; {
		if randInt(65, 91) != temp {
			temp = randInt(65, 91)
			result.WriteByte(temp)
			i++
		}
	}
	return result.String()
}

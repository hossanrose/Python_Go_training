package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

type Node struct {
	data       int
	next      *Node
}

type Array struct{
        x []int
}


type Stack interface {
	pop()
	push(int)
	printList()
}

func stack_push(n Stack, elem int) {
	n.push(elem)
}
func stack_pop(n Stack ) {
	n.pop()
}
func printme(n Stack ) {
	n.printList()
}


func main() {
	var wg sync.WaitGroup
        var mx sync.Mutex
	fmt.Println("-----Stack -----")
	fmt.Println("1. Linked list: Non parallel")
	fmt.Println("2. Linked list: Parallel")
	fmt.Println("3. Array: No Parallel")
	fmt.Println("4. Array: Parallel")
	// Non parallel linked list
	var choice int
        fmt.Scan(&choice)
        switch choice {
        case 1:
	// Non parallel linked list
		n := &Node{}
		fmt.Println("Enter Loop length:")
		var choice int
		fmt.Scan(&choice)
		start := time.Now()
		for i:=0; i < choice ; i++ {
			stack_push(n, i)
		}
		for i:=choice ; i> 0 ; i-- {
			stack_pop(n)
		}
		elapsed := time.Since(start)
		fmt.Println("Time of execution -------", elapsed)
	case 2:
	// Parallel linked list 
		n := &Node{}
		fmt.Println("Enter no of goroutines:")
		var choice int
		fmt.Scan(&choice)
		start := time.Now()
                for i:=0;i<choice;i++{
                        wg.Add(2)
                        go producer_list(n,i,&wg,&mx)
                        go consumer_list(n,i,&wg,&mx)
                }
		elapsed := time.Since(start)
                wg.Wait()
		fmt.Println("Time of execution -------:", elapsed)
	case 3:
	// Non parallel Array
		arr:= &Array{}
		fmt.Println("Enter Loop length:")
		var choice int
		fmt.Scan(&choice)
		start := time.Now()
                for i:=0; i < choice; i++ {
                        stack_push(arr, i)
                }
                for i:=choice ; i> 0 ; i-- {
                        stack_pop(arr)
                }
		elapsed := time.Since(start)
		fmt.Println("Time of execution -------", elapsed)
	case 4:
	// Parallel Array
		ch :=make(chan *Array)
		fmt.Println("Enter no of goroutines:")
		var choice int
		fmt.Scan(&choice)
		start := time.Now()
		for i:=0;i<choice;i++{
			wg.Add(2)
			go producer_array(ch,i,&wg)
			go consumer_array(ch,i,&wg)
		}
		elapsed := time.Since(start)
		wg.Wait()
		fmt.Println("Time of execution -------", elapsed)
        default:
                fmt.Println("Enter a valid  option")

	}
}

func producer_list(n *Node, id int, wg *sync.WaitGroup, mx *sync.Mutex) {
        defer wg.Done()
        mx.Lock()
	rand:=rand.Intn(100)
	fmt.Println("Producer", id,  "-- pushing:",rand)
        stack_push(n, rand)
        mx.Unlock()
}
func consumer_list(n *Node, id int, wg *sync.WaitGroup, mx *sync.Mutex) {
        defer wg.Done()
        mx.Lock()
	fmt.Println("Consumer", id, "-- poping")
        stack_pop(n)
        mx.Unlock()
}

func producer_array(ch chan *Array,id int,wg *sync.WaitGroup) {
	defer wg.Done()
	rand:=rand.Intn(100)
	select {
	case test:= <-ch:
		fmt.Println("Producer", id , "-- pushing:",rand)
		stack_push(test,rand)
		ch <- test
	default:
		temp:=&Array{}
		fmt.Println("Producer" ,id ,"-- pushing:",rand)
		stack_push(temp,rand)
		ch <- temp
  }
}

func consumer_array(ch chan *Array,id int, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case temp:=<-ch:
		if len(temp.x) >0{
			stack_pop(temp)
			fmt.Println("Consumer" ,id, "-- poping")
			if len(temp.x) >0{
				ch <- temp
			}
		} else {
			fmt.Println("Nothing to pop")
		}
	}
}


func (n *Node) printList() {
	var output string
	for i := n; i.next != nil; i = i.next {
		output =  fmt.Sprintf("%s %v", output,  i.data )
	}
	fmt.Println("Stack:", output)
}

func (n *Node) push(elem int)  {
	if n == nil {
		fmt.Println("Empty")
	}
	for i := n; i != nil; i = i.next {
		if i.next == nil {
			i.next =  &Node{data:elem,next:nil}
			printme(n)
			break
		}
	}
}

func (n *Node) pop()  {
	if n == nil {
		fmt.Println("Empty")
	}
	for i := n; i.next != nil; i = i.next {
		if i.next.next == nil {
			i.next = nil
			printme(n)
			break
	}
   }
}

func (a *Array) printList() {
	fmt.Println("Array:", a)
}

func (array *Array) push(element int) {
	array.x = append(array.x, element )
	fmt.Println(array)
}

func (array *Array) pop() {
	array.x = array.x[:len(array.x)-1]
	fmt.Println(array)
}


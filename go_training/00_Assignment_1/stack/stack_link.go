package main

import (
	"fmt"
	"sync"
	"math/rand"
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
}

func stack_push(n Stack, elem int) {
	n.push(elem)
}
func stack_pop(n Stack ) {
	n.pop()
}


func main() {
	var wg sync.WaitGroup
        var mx sync.Mutex
        fmt.Println("1. Non parallel Linked list")
        fmt.Println("2. Parallel Linked list")
        fmt.Println("3. Non parallel Circular Queue")
        fmt.Println("4. Parallel Circular Queue")
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
		for i:=0; i < choice ; i++ {
			stack_push(n, i)
		}
		for i:=choice ; i> 0 ; i-- {
			stack_pop(n)
		}
	case 2:
	// Parallel linked list 
		n := &Node{}
		fmt.Println("Enter no of goroutines:")
		var choice int
		fmt.Scan(&choice)
                for i:=0;i<choice;i++{
                        wg.Add(2)
                        go producer_list(n,i,&wg,&mx)
                        go consumer_list(n,i,&wg,&mx)
                }
                wg.Wait()
                printList(n)
	case 3:
		arr:= &Array{}
		fmt.Println("Enter Loop length:")
		var choice int
		fmt.Scan(&choice)
                for i:=0; i < 10; i++ {
                        stack_push(arr, i)
                }
                for i:=10 ; i> 0 ; i-- {
                        stack_pop(arr)
                }
	case 4:
		ch :=make(chan *Array)
		fmt.Println("Enter no of goroutines:")
		var choice int
		fmt.Scan(&choice)
		for i:=0;i<choice;i++{
			wg.Add(2)
			go producer_array(ch,i,&wg)
			go consumer_array(ch,i,&wg)
		}
		wg.Wait()
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


func printList(n *Node) {
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
			printList(n)
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
			printList(n)
			break
	}
   }
}

func (array *Array) push(element int) {
	array.x = append(array.x, element )
	fmt.Println(array)
}

func (array *Array) pop() {
	array.x = array.x[:len(array.x)-1]
	fmt.Println(array)
}


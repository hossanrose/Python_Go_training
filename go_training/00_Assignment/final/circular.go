package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Circle interface{
	insert(int)
	delete()
	printList()
}


type Queue struct{
	size,rear, front int
	array [10]int
}

type Node struct{
	data int
	next *Node
	last *Node
}

func push (c Circle, elem int){
	c.insert(elem)
}

func pop (c Circle){
	c.delete()
}

func printme (c Circle){
	c.printList()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	fmt.Println("----- Circular Buffer -----")
	fmt.Println("1. Array: Non Parallel")
	fmt.Println("2. Array: Parallel")
	fmt.Println("3. Linked list: Non Parallel")
	fmt.Println("4. Linked list: Parallel")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
	// Non prallel Circular queue
		a:= &Queue{rear:-1,front:-1,size:10}
		fmt.Println("Enter Loop length:")
                var choice int
                fmt.Scan(&choice)
                for i:=0; i < choice ; i++ {
                        push(a, i)
                }
                for i:=choice ; i> 0 ; i-- {
                        pop(a)
                }
	case 2:
       // Parallel Circular queue
                a:= &Queue{rear:-1,front:-1,size:10}
                fmt.Println("Enter no of goroutines:")
                var choice int
                fmt.Scan(&choice)
                for i:=0;i<choice;i++{
                        wg.Add(2)
                        go producer_array(a,i,&wg,&mx)
                        go consumer_array(a,i,&wg,&mx)
                }
                wg.Wait()
                fmt.Println(a)

	case 3:
        // Non Parallel Circular linked
                n := &Node{}
                fmt.Println("Enter Loop length:")
                var choice int
                fmt.Scan(&choice)
                for i:=0; i < choice ; i++ {
                        push(n, i)
                }
                for i:=choice ; i> 0 ; i-- {
                        pop(n)
                }

	case 4:
	// Parallel Linkedlist queue
		n := &Node{}
		fmt.Println("Enter no of goroutines:")
                var choice int
                fmt.Scan(&choice)
		for i:=0;i<choice;i++{
			wg.Add(2)
			go producer(n,i,&wg,&mx)
			go consumer(n,i,&wg,&mx)
		}
		wg.Wait()
		printme(n)
	default:
		fmt.Println("Enter a valid  option")
	}
}

func producer(n *Node,id int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	rand:=rand.Intn(100)
	fmt.Println("Producer", id,  "-- pushing:",rand)
	push(n, rand)
	mx.Unlock()
}

func consumer(n *Node,id int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	fmt.Println("Consumer", id, "-- poping")
	pop(n)
	mx.Unlock()
}

func producer_array(a *Queue, id int, wg *sync.WaitGroup, mx *sync.Mutex) {
        defer wg.Done()
        mx.Lock()
	rand:=rand.Intn(100)
	fmt.Println("Producer", id,  "-- pushing:",rand)
        push(a, rand)
        mx.Unlock()
}

func consumer_array(a *Queue,id int, wg *sync.WaitGroup, mx *sync.Mutex) {
        defer wg.Done()
        mx.Lock()
	fmt.Println("Consumer", id, "-- poping")
        pop(a)
        mx.Unlock()
}

func (a *Queue) insert(  elem int) {
	//Checks if  queue is full
	if  (a.front == 0 &&  a.rear  == a.size-1 ) || ( a.rear ==  (a.front-1 ) % a.size-1) {
		fmt.Println("Full")
	// Checks if this is the first insertion
	} else if  a.front == -1 {
		a.front , a.rear = 0,0
		a.array[a.rear]=elem
		printme(a)
	// Checks if rear is at the end	
	} else if  a.rear == a.size-1  &&  a.front != 0 {
		a.rear = 0
		a.array[a.rear]=elem
		printme(a)
	}else{
		a.rear++
		a.array[a.rear]=elem
		printme(a)
	}
}

func (a *Queue ) delete( ) {
	if ( a.front == -1 ){
		fmt.Println("Empty")
	} else {
		// Putting in zero as placeholder value
		a.array[a.front]=0
		printme(a)
		if ( a.front == a.rear ){
			a.front,a.rear =-1,-1
		} else if a.front == a.size -1 {
			a.front = 0
		} else {
			a.front++
		}
	}
}

func (n *Queue) printList( ) {
	fmt.Println("Array", n)
}

func  (n *Node )insert( elem int) {
	// Create head node and circular link
	if n.next == nil {
		n.data=elem
		n.next=n
		n.last=n
		printme(n)
	}else {
		new_node:=&Node{data:elem,next:n.last.next}
		n.last.next=new_node
		//fmt.Println(n.last.next)
		n.last=new_node
		printme(n)
	}
}


func (n *Node) delete () {
	// If empty or only head node 
	if n.next==nil || n == n.next {
		fmt.Println("Empty")
	}else {
		n.data=n.next.data
		n.next=n.next.next
		printme(n)
	}
}

func (n *Node) printList (){
        var output string
        for i := n; i != n.last ; i = i.next {
                output =  fmt.Sprintf("%s %v", output,  i.data )
		if i==i.next{
			break
		}
        }
	fmt.Println("Stack:", output)
}

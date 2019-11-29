package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Circle interface{
	insert(int)
	delete()
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


func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	fmt.Println("1. Non parallel Circular Queue")
	fmt.Println("2. Non parallel Circular Linked list")
	fmt.Println("3. Parallel Circular Queue")
	fmt.Println("4. Parallel Circular Linked list")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
	// Non prallel Circular queue
	a:= &Queue{rear:-1,front:-1,size:10}
	push(a,1)
	push(a,2)
	push(a,3)
	push(a,4)
	push(a,5)
	push(a,6)
	push(a,7)
	pop(a)
	pop(a)
	push(a,8)
	push(a,9)
	push(a,10)
	push(a,11)
	fmt.Println(a)
	case 2:
	// Non Parallel Circular linked
	n := &Node{}
	push(n,1)
	push(n,2)
	push(n,3)
	pop(n)
	push(n,4)
	push(n,5)
	push(n,6)
	pop(n)
	push(n,7)
	push(n,8)
	push(n,9)
	pop(n)
	push(n,10)
	push(n,11)
	case 3:
		// Parallel Circular queue
		a:= &Queue{rear:-1,front:-1,size:10}
		num := 10
		for i:=0;i<num;i++{
			wg.Add(2)
			go producer_array(a,&wg,&mx)
			go consumer_array(a,&wg,&mx)
		}
		wg.Wait()
		fmt.Println(a)
	case 4:
		// Parallel Linkedlist queue
		n := &Node{}
		num := 10
		for i:=0;i<num;i++{
			wg.Add(2)
			go producer(n,&wg,&mx)
			go consumer(n,&wg,&mx)
		}
		wg.Wait()
		printList(n)
	default:
		fmt.Println("Enter a valid  option")
	}
}
func producer(n *Node, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	fmt.Println("Test")
	fmt.Println(rand.Intn(100))
	push(n, rand.Intn(100))
	mx.Unlock()
}
func consumer(n *Node, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	pop(n)
	mx.Unlock()
}
func producer_array(a *Queue, wg *sync.WaitGroup, mx *sync.Mutex) {
        defer wg.Done()
        mx.Lock()
        fmt.Println("Test")
        fmt.Println(rand.Intn(100))
        push(a, rand.Intn(100))
        mx.Unlock()
}
func consumer_array(a *Queue, wg *sync.WaitGroup, mx *sync.Mutex) {
        defer wg.Done()
        mx.Lock()
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
	// Checks if rear is at the end	
	} else if  a.rear == a.size-1  &&  a.front != 0 {
		a.rear = 0
		fmt.Println("Test")
		a.array[a.rear]=elem
	}else{
		a.rear++
		a.array[a.rear]=elem
	}
}

func (a *Queue ) delete( ) {
	if ( a.front == -1 ){
		fmt.Println("Empty")
	} else {
		fmt.Println(a.front)
		fmt.Println(a.array[a.front])
		a.array[a.front]=0
		if ( a.front == a.rear ){
			a.front,a.rear =-1,-1
		} else if a.front == a.size -1 {
			a.front = 0
		} else {
			a.front++
		}
	}
}

func  (n *Node )insert( elem int) {
	// Create head node and circular link
	if n.next == nil {
		n.data=elem
		n.next=n
		//n=temp
		n.last=n
	}else {
	new_node:=&Node{data:elem,next:n.last.next}
	n.last.next=new_node
	//fmt.Println(n.last.next)
	n.last=new_node
}
}

func printList (n *Node){
	for i:=n; i!=n.last;i=i.next{
		fmt.Println(i)
	}
}

func (n *Node) delete () {
	// If empty or only head node 
	if n==nil ||  n.next == n {
		n.data=n.next.data
		n.next=n.next.next
	}
}

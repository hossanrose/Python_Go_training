package main

import "fmt"

type Queue struct{
	size,rear, front int
	array [10]int
}

func main() {
	a:= &Queue{rear:-1,front:-1,size:10}
	fmt.Println(a)
	insert(a,34)
	fmt.Println(a)
	insert(a,4)
	insert(a,5)
	insert(a,6)
	insert(a,7)
	insert(a,8)
	insert(a,9)
	insert(a,10)
	insert(a,11)
	insert(a,12)
	fmt.Println(a)
	insert(a,13)
	fmt.Println(a)
	insert(a,14)
	insert(a,15)
	insert(a,16)
	insert(a,17)
	insert(a,18)
	fmt.Println(a)
	fmt.Println(a)
	delete(a)
	fmt.Println(a)
	delete(a)
	fmt.Println(a)
	insert(a,13)
	fmt.Println(a)
}

func insert( a *Queue ,  elem int) *Queue{
	//Checks if  queue is full
	if  (a.front == 0 &&  a.rear  == a.size-1 ) || ( a.rear ==  (a.front-1 ) % a.size-1) {
		return a
	// Checks if this is the first insertion
	} else if  a.front == -1 {
		a.front , a.rear = 0,0
		a.array[a.rear]=elem
		return a
	// Checks if rear is at the end	
	} else if  a.rear == a.size-1  &&  a.front != 0 {
		a.rear = 0
		fmt.Println("Test")
		a.array[a.rear]=elem
		return a
	}else{
		a.rear++
		a.array[a.rear]=elem
		return a
	}
}

func delete( a *Queue ) *Queue{
	if ( a.front == -1 ){
		fmt.Println("Empty")
		return a
	}
	fmt.Println(a.array[a.front])
	a.array[a.front]=0
	if ( a.front == a.rear ){
		a.front,a.rear =-1,-1
	} else if a.front == a.size -1 {
		a.front = 0
	}else{
		a.front++
	}
	return a
}



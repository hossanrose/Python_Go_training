package main

import "fmt"

type Node struct{
	data int
	next *Node
}

func main () {
	n := &Node{}
	n=push(n,25)
	n=push(n,26)
	n=push(n,27)
	n=pop(n)
	n=pop(n)
	n=push(n,28)
	n=push(n,29)
	n=push(n,30)
	n=push(n,31)
	printList(n)
}

func push(n *Node, elem int) *Node{
	// Create head node and circular link
	if n.next == nil {
		n.next=n
		n.data=elem
		return n
	}
	new_node:=&Node{data:elem,next:n.next}
	n.next=new_node
	//fmt.Println(new_node)
	return new_node
}

func printList (n *Node){
	for i:=n.next;i!=n;i=i.next{
		fmt.Println(i)
		if i.next==n{
			fmt.Println(i.next)
		}
	}
}

func pop (n *Node) *Node {
	// If empty or only head node 
	if n.next==nil &&  n.next==n{
		return n
	}
	n.next=n.next.next
	return n
}

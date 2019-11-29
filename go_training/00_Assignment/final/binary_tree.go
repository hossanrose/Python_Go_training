package main

import (
	"fmt"
)

var array []int

type Tree struct{
	data int
	left,right *Tree
}

type Stack interface{
	push(elem int)
	pop(elem int, prev *Tree)
}

func main() {
	// Non parallel binary tree
	T:=&Tree{data:10}
	stack_push(T,1)
	stack_push(T,20)
	stack_push(T,2)
	stack_push(T,30)
	stack_push(T,40)
	stack_push(T,0)
	fmt.Println("Non parallel binary tree: After push")
	printMe(T)
	temp:=array[len(array)-1]
	array=array[:len(array)-1]
	stack_pop(T,temp,T)
	temp=array[len(array)-1]
	array=array[:len(array)-1]
	stack_pop(T,temp,T)
	temp=array[len(array)-1]
	array=array[:len(array)-1]
	stack_pop(T,temp,T)
	fmt.Println("Non parallel binary tree: After pop")
	printMe(T)
	//fmt.Println(array)
}


func printMe( T *Tree){
	fmt.Println("Elements", T.data)
	if  T.left != nil{
		 printMe(T.left)
	}
	if  T.right != nil{
                printMe(T.right)
        }
}


func stack_push( Tree Stack, elem int ) {
	Tree.push(elem)
}
func stack_pop( Tree Stack, elem int, prev *Tree ) {
	Tree.pop(elem,prev )
}


func (t *Tree) push(elem int)  {
        if t == nil {
		fmt.Println("Empty")
        }
	if  elem == t.data {
		fmt.Println(elem ,"is the same value in root node")
	} else if  elem < t.data {
                if t.left == nil {
			t.left =&Tree{data:elem,left:nil,right:nil}
			array=append(array, elem)
		} else {
			stack_push(t.left, elem)
		}
	} else  {
		if  t.right == nil {
			t.right = &Tree{data:elem,left:nil,right:nil}
			array=append(array, elem)
                } else {
			stack_push(t.right, elem)
		}
	}
}

func (t *Tree ) pop( elem int,prev *Tree ) {
        if t == nil {
		fmt.Println("Empty")
        }
	if  t!= nil && elem < t.data  {
		prev:=t
		t=t.left
		stack_pop(t,elem, prev)
	}else if t!= nil &&  elem > t.data {
		prev:=t
		t=t.right
		stack_pop(t,elem,prev)
	}else{
			if prev.right !=nil && prev.right.data == elem {
				prev.right = nil
			} else if prev.left != nil &&  prev.left.data == elem {
				prev.left = nil
			} else {
				t=nil
			}
		}
	}




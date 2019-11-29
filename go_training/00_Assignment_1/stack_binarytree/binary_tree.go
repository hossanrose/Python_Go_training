package main

import (
	"fmt"
)


type binaryTree struct{
	data int
	left,right *binaryTree
}

type myStack interface{
//	removeNodeEnd() *binaryTree
	addNodeEnd(*binaryTree) *binaryTree
}

func main() {
	A:=&binaryTree{left:nil,data:1,right:nil}
	Tree:=  A
	B:=&binaryTree{left:nil,data:2,right:nil}
	C:=&binaryTree{left:nil,data:3,right:nil}
	D:=&binaryTree{left:nil,data:4,right:nil}
	Tree=stack_push(B,Tree)
	Tree=stack_push(C,Tree)
	Tree=stack_push(D,Tree)
	printMe(Tree)
}


/*
func printMe( Tree *binaryTree) {
	fmt.Println(Tree)
	for t := Tree; t != nil; t = t.left {
		if t.left == nil {
			fmt.Println(t)
		} else  if  t.right == nil {
			fmt.Println(t)
		}
	}
}
*/

func printMe( Tree *binaryTree){
	fmt.Println(Tree)
	if  Tree.left != nil{
		printMe(Tree.left)
	}
	if  Tree.right != nil{
                printMe(Tree.right)
        }
}


func stack_push(newNode *binaryTree, Tree myStack ) *binaryTree{
	return Tree.addNodeEnd(newNode)
}


/*
func stack_pop(Tree myStack ) *binaryTree{
	return Tree.removeNodeEnd()
}
*/


func (Tree *binaryTree) addNodeEnd(newNode *binaryTree) *binaryTree {
        if Tree == nil {
                return Tree
        }
	if  newNode.data < Tree.data {
                if Tree.left == nil {
                        Tree.left = newNode
                        return Tree
		} else {
			stack_push(newNode, Tree.left)
		}
	} else  {
		if  Tree.right == nil {
			Tree.right = newNode
			return Tree
                } else {
			stack_push(newNode, Tree.right)
		}
	}

        return Tree
}

/*
func (Tree *binaryTree) removeNodeEnd( ) *binaryTree {
        if Tree == nil {
                return Tree
        }
}
*/

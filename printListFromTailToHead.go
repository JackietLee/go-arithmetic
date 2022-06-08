package main

import "fmt"

type NodeList struct {
	val  int
	next *NodeList
}

func printListFromTailToHead(head *NodeList) {
	if head != nil {
		printListFromTailToHead(head.next)
		fmt.Println(head.val)
	}
}

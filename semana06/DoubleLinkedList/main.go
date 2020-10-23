package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List
	l.PushBack(1)
	l.PushBack("hello")
	l.PushBack("world")
	l.PushFront(3.4)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(l)
}

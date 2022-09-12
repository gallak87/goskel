package main

import (
	"fmt"
)

type Node struct {
	index int
	next  *Node
}

type Graph struct {
	size    int
	adjList []*Node
}

func (n *Node) getLast() *Node {
	tmp := n
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	return tmp
}

// adds edge in both directions
func (g Graph) addEdge(from, to int) {
	if from >= g.size || to >= g.size {
		panic(fmt.Sprintf("addEdge: from %d or to %d exceeds size of graph", from, to))
	}
	if from == to {
		panic("addEdge: from and to cannot be the same")
	}
	g.adjList[from].getLast().next = &Node{index: to}
	g.adjList[to].getLast().next = &Node{index: from}
}

func (g Graph) addNode(i int) {
	if i >= g.size {
		panic("addNode: i is out of bounds")
	}
	g.adjList[i] = &Node{index: i}
}

func (g Graph) print() {
	for i := 0; i < g.size; i++ {
		fmt.Printf("node %d -> [ ", g.adjList[i].index)
		for n := g.adjList[i].next; n != nil; n = n.next {
			fmt.Printf("%d ", n.index)
		}
		fmt.Println("]")
	}
}

func main() {
	g := Graph{
		size:    5,
		adjList: make([]*Node, 5, 5),
	}
	for i := 0; i < g.size; i++ {
		g.addNode(i)
	}
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 4)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)

	g.print()
}

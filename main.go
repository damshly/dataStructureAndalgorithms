package main

import (
	"fmt"

	"github.com/data-structure/Graph"
)

func main() {
	graph := graph.NewGraph()

	// إضافة العقد
	nodeA := graph.AddNode("A")
	nodeB := graph.AddNode("B")
	nodeC := graph.AddNode("C")
	nodeD := graph.AddNode("D")

	// إضافة الحواف
	graph.AddEdge(nodeA, nodeB)
	graph.AddEdge(nodeA, nodeC)
	graph.AddEdge(nodeB, nodeD)
	graph.AddEdge(nodeC, nodeD)

	// طباعة الرسم البياني
	fmt.Println(graph.HasEdge(nodeA,nodeC))
	graph.PrintGraph()
	graph.RemoveEdge(nodeA,nodeC)
	fmt.Println("__________")
	graph.PrintGraph()
	fmt.Println(graph.HasEdge(nodeA,nodeC))

}



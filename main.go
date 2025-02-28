package main

import (


	"github.com/dataStructureAndalgorithms/mapGraph"
)


func main() {
	mapgraph := mapGraph.NewGraph()
	// fmt.Println(queue.Peek())
	// queue.Enqueue("damshly")
	// fmt.Println(queue.Peek())
	// queue.Dequeue()
	// fmt.Println(queue.Peek())
	// إضافة العقد
	// Manually adding 20 alphabetic nodes
	nodeA := mapgraph.AddNode("A")
	nodeB := mapgraph.AddNode("B")
	nodeC := mapgraph.AddNode("C")
	nodeD := mapgraph.AddNode("D")
	nodeE := mapgraph.AddNode("E")
	nodeF := mapgraph.AddNode("F")
	nodeG := mapgraph.AddNode("G")
	nodeH := mapgraph.AddNode("H")
	nodeI := mapgraph.AddNode("I")
	nodeJ := mapgraph.AddNode("J")
	nodeK := mapgraph.AddNode("K")
	nodeL := mapgraph.AddNode("L")
	nodeM := mapgraph.AddNode("M")
	nodeN := mapgraph.AddNode("N")
	nodeO := mapgraph.AddNode("O")
	nodeP := mapgraph.AddNode("P")
	nodeQ := mapgraph.AddNode("Q")
	nodeR := mapgraph.AddNode("R")
	nodeS := mapgraph.AddNode("S")
	nodeT := mapgraph.AddNode("T")

	// Manually adding 30 edges
	mapgraph.AddEdge(nodeB, nodeA)
	mapgraph.AddEdge(nodeB, nodeE)
	mapgraph.AddEdge(nodeA, nodeB)
	mapgraph.AddEdge(nodeA, nodeC)
	mapgraph.AddEdge(nodeK, nodeL)
	mapgraph.AddEdge(nodeG, nodeH)
	mapgraph.AddEdge(nodeC, nodeD)
	mapgraph.AddEdge(nodeI, nodeJ)
	mapgraph.AddEdge(nodeA, nodeE)
	mapgraph.AddEdge(nodeM, nodeN)
	mapgraph.AddEdge(nodeO, nodeP)
	mapgraph.AddEdge(nodeE, nodeF)
	mapgraph.AddEdge(nodeH, nodeL)
	mapgraph.AddEdge(nodeB, nodeF)
	mapgraph.AddEdge(nodeC, nodeG)
	mapgraph.AddEdge(nodeS, nodeT)
	mapgraph.AddEdge(nodeQ, nodeR)
	mapgraph.AddEdge(nodeD, nodeH)
	mapgraph.AddEdge(nodeE, nodeI)
	mapgraph.AddEdge(nodeG, nodeK)
	mapgraph.AddEdge(nodeL, nodeP)
	mapgraph.AddEdge(nodeI, nodeM)
	mapgraph.AddEdge(nodeF, nodeJ)
	mapgraph.AddEdge(nodeJ, nodeN)
	mapgraph.AddEdge(nodeK, nodeO)
	mapgraph.AddEdge(nodeM, nodeQ)
	mapgraph.AddEdge(nodeN, nodeR)
	mapgraph.AddEdge(nodeO, nodeS)
	mapgraph.AddEdge(nodeP, nodeT)

	// Print the graph
	

	// pathatoj := mapgraph.FindPaths(nodeA, nodeJ)
	mapgraph.PrintFindPaths(nodeA, nodeM)
	// fmt.Println(queue.Peek())
	// mapgraph.PrintGraph(false)
	// // // طباعة الرسم البياني
	// // fmt.Println(graph.HasEdge(nodeA,nodeC))
	// // fmt.Println(mapgraph.HasEdge(nodeA,nodeB))
	// // mapgraph.RemoveEdge(nodeA,nodeB)
	// // fmt.Println(mapgraph.HasEdge(nodeA,nodeB))
	// // mapgraph.RemoveNode(nodeC)
	// mapgraph.Clear()
	// fmt.Println("++++++++")
	// mapgraph.PrintGraph()
	// fmt.Println(mapgraph.NodeCount())
	// graph.RemoveEdge(nodeA,nodeC)
	// fmt.Println("__________")
	// graph.PrintGraph()
	// fmt.Println(graph.HasEdge(nodeA,nodeC))

}


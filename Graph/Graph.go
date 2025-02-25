package graph


import (
	"fmt"
)

// Node يمثل عقدة في الرسم البياني
type Node struct {
	value    string
	neighbors []*Node
}

// Graph يمثل الرسم البياني باستخدام العقد
type Graph struct {
	nodes []*Node
}

// NewGraph ينشئ رسماً بيانياً جديداً
func NewGraph() *Graph {
	return &Graph{nodes: []*Node{}}
}

// AddNode يضيف عقدة جديدة للرسم البياني
func (g *Graph) AddNode(value string) *Node {
	node := &Node{value: value}
	g.nodes = append(g.nodes, node)
	return node
}

// AddEdge يضيف حافة بين عقدتين
func (g *Graph) AddEdge(node1, node2 *Node) {
	node1.neighbors = append(node1.neighbors, node2)
	node2.neighbors = append(node2.neighbors, node1) // إذا كان غير موجه
}

// PrintGraph يطبع الرسم البياني
func (g *Graph) PrintGraph() {
	for _, node := range g.nodes {
		fmt.Printf("%s -> ", node.value)
		for _, neighbor := range node.neighbors {
			fmt.Printf("%s ", neighbor.value)
		}
		fmt.Println()
	}
}

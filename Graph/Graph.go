package graph


import (
	"fmt"
)
func remove(slice []*Node, index int) []*Node {
	return append(slice[:index], slice[index+1:]...)
}
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

func (g *Graph)FindNode(value string) *Node{
	for _,node := range g.nodes{
		if node.value == value{return node}
	}
	return nil
}
func (g *Graph) HasEdge(node1, node2 *Node) bool {
	for _,neighbors := range node1.neighbors{
		if node2 == neighbors{
			return true}
	}
	return false
}

func (g *Graph)RemoveEdge(node1 ,node2 *Node){
	for k,v := range node1.neighbors{
		if v.value == node2.value{
			node1.neighbors = remove(node1.neighbors,k)
			break
		}}
	for k,v := range node2.neighbors{
		if v.value == node1.value{
			node2.neighbors = remove(node2.neighbors,k)
			break
	}
}
}
package mapGraph

import (
	"fmt"

	"github.com/dataStructureAndalgorithms/Queue"
)

type Node struct {
	Value            string
	Neighbors        map[string]*Node
	Revarceneighbors map[string]*Node
}

type MapGraph struct {
	nodes map[string]*Node
}
func contains(path []*Node, node *Node) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

func NewGraph() *MapGraph {
	return &MapGraph{nodes: make(map[string]*Node)}
}

func (g *MapGraph) PrintGraph(b bool) {
	if g.IsEmpty() {
		fmt.Println("The graph is empty")
		return
	}
	for key, node := range g.nodes {
		if !b && len(node.Neighbors) == 0 {
			continue
		}

		fmt.Printf("%s ->", key)
		for _, neighbor := range node.Neighbors {
			fmt.Printf(" %s", neighbor.Value)
		}
		fmt.Println()
	}
}

func (g *MapGraph) AddNode(value string) *Node {
	n := &Node{
		Value:            value,
		Neighbors:        make(map[string]*Node),
		Revarceneighbors: make(map[string]*Node),
	}
	g.nodes[value] = n // هون أضفنا العقدة للخريطة باستخدام قيمتها كمفتاح
	return n
}

func (g *MapGraph) AddEdge(node1, node2 *Node) error {
	if err := g.EnsureNodeExists(node1); err != nil {
		return err
	}
	if err := g.EnsureNodeExists(node2); err != nil {
		return err
	}
	n1Key := node1.Value
	n2Key := node2.Value

	// أضف node2 كجار لـ node1
	node1.Neighbors[n2Key] = node2

	// أضف node1 كـ reverse neighbor لـ node2
	node2.Revarceneighbors[n1Key] = node1
	return nil
}

func (g *MapGraph) HasEdge(node1, node2 *Node) (bool, error) {
	if err := g.EnsureNodeExists(node1); err != nil {
		return false, err
	}
	if err := g.EnsureNodeExists(node2); err != nil {
		return false, err
	}
	_, v := node1.Neighbors[node2.Value]
	if v {
		return true, nil
	}

	return false, nil
}

func (g *MapGraph) RemoveEdge(node1, node2 *Node) error {
	exists, err := g.HasEdge(node1, node2)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("there is no edge %s -> %s", node1.Value, node2.Value)
	}
	delete(node1.Neighbors, node2.Value)
	delete(node2.Revarceneighbors, node1.Value)
	return nil
}

func (g *MapGraph) RemoveNode(node *Node) error {
	if err := g.EnsureNodeExists(node); err != nil {
		return err
	}
	for _, v := range node.Neighbors {
		g.RemoveEdge(node, v) // هاد السطر بكفي، لأنه بحذف العلاقة العكسية كمان
	}
	// حذف كل الحواف العكسية (لو في عقد بتشير لهادي العقدة)
	for _, v := range node.Revarceneighbors {
		g.RemoveEdge(v, node)
	}
	delete(g.nodes, node.Value)
	return nil
}

func (g *MapGraph) GetNeighbors(node *Node) ([]*Node, error) {
	if err := g.EnsureNodeExists(node); err != nil {
		return nil, err
	}

	list := make([]*Node, 0, len(node.Neighbors))
	for key := range node.Neighbors {
		list = append(list, g.nodes[key])
	}
	return list, nil
}

func (g *MapGraph) EdgeCount(node *Node) (int, error) {
	if err := g.EnsureNodeExists(node); err != nil {
		return 0, err
	}
	return len(node.Neighbors), nil
}

func (g *MapGraph) EnsureNodeExists(node *Node) error {
	if _, exist := g.nodes[node.Value]; !exist {
		return fmt.Errorf("the node %s does not exist", node.Value)
	}
	return nil
}

func (g *MapGraph) GetReverseNeighbors(node *Node) ([]string, error) {
	if err := g.EnsureNodeExists(node); err != nil {
		return nil, err
	}

	list := make([]string, 0, len(node.Neighbors))
	for key := range node.Revarceneighbors {
		list = append(list, key)
	}
	return list, nil
}

func (g *MapGraph) NodeCount() int {
	return len(g.nodes)
}

func (g *MapGraph) IsEmpty() bool {
	return len(g.nodes) == 0
}

func (g *MapGraph) Clear() error {
	g.nodes = make(map[string]*Node)
	return nil
}

func (g *MapGraph) BFS(node *Node, mapgraph *MapGraph) {
	queue := Queue.NewQueue[*Node]()
	visited := make(map[*Node]bool)
	queue.Enqueue(node)
	visited[node] = true
	for !queue.IsEmpty() {
		neighbors, _ := mapgraph.GetNeighbors(queue.Peek())
		elemnt, _ := queue.Dequeue()
		fmt.Println(elemnt.Value)
		for _, v := range neighbors {
			if !visited[v] {
				queue.Enqueue(v)
				visited[v] = true
			}
		}
	}

}

func (g *MapGraph) FindPaths(start *Node, end *Node) [][]*Node {
	queue := Queue.NewQueue[[]*Node]()
	queue.Enqueue([]*Node{start})
	paths := [][]*Node{}
	for !queue.IsEmpty() {
		currentPath,_:= queue.Dequeue()
		currentNode := currentPath[len(currentPath)-1]
		if currentNode == end {
			paths = append(paths,currentPath)
			continue
		}
		neighbors,_ := g.GetNeighbors(currentNode)
		for _,neighbor := range neighbors {
			if !contains(currentPath, neighbor) { // تجنب الحلقات
                newPath := append([]*Node{}, currentPath...) // نسخة جديدة
                newPath = append(newPath, neighbor)
                queue.Enqueue(newPath) // إضافة المسار الجديد إلى الطابور
            }
		}

		}
		return paths
}

func (g *MapGraph) PrintFindPaths(start *Node, end *Node) {
	paths := g.FindPaths(start, end)
	if len(paths) == 0 {
		fmt.Println("No paths found")
		return
	}

	shortestPathLen := len(paths[0]) // أقصر مسار هو أول عنصر
	for _, path := range paths {
		if len(path) == shortestPathLen { // فقط أطبع المسارات اللي بنفس طول أقصر مسار
			for i, node := range path {
				if i > 0 {
					fmt.Print(" -> ")
				}
				fmt.Print(node.Value)
			}
			fmt.Println()
		}
	}
}



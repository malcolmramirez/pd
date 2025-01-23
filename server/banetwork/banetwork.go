package banetwork

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type BANetwork struct {
	M     int
	Edges map[Edge]bool
	Nodes map[Node]bool
}

type Edge struct {
	Left  Node
	Right Node
}

type Node struct {
	Value int
}

func NewNetwork(m int) *BANetwork {
	network := new(BANetwork)
	network.Edges = make(map[Edge]bool)
	network.Nodes = make(map[Node]bool)
	network.M = m

	return network
}

func InitializeNetwork(m int, size int) *BANetwork {
	network := NewNetwork(m)
	for i := 0; i < size; i++ {
		AddNode(network)
	}
	return network
}

func AddNode(network *BANetwork) {
	newNode := Node{
		Value: rand.IntN(100000),
	}

	edges := []Edge{}
	for edge := range network.Edges {
		edges = append(edges, edge)
	}

	nodes := []Node{}
	for node := range network.Nodes {
		nodes = append(nodes, node)
	}

	numEdgesToAdd := min(len(nodes), network.M)

	for i := 0; i < numEdgesToAdd; i++ {
		var selectedNode Node
		if len(edges) > 0 {
			selectedEdge := Sample(edges)
			selectedNode = Sample([]Node{selectedEdge.Left, selectedEdge.Right})
		} else {
			selectedNode = Sample(nodes)
		}

		newEdge := Edge{
			Left:  newNode,
			Right: selectedNode,
		}
		network.Edges[newEdge] = true
	}

	// Add node to graph
	network.Nodes[newNode] = true
}

func WriteDot(network *BANetwork) {
	if len(network.Edges) <= 0 {
		return
	}

	var sb strings.Builder
	for edge := range network.Edges {
		sb.WriteString(fmt.Sprintf("\t%d -- %d\n", edge.Left, edge.Right))
	}

	file, err := os.Create("graph.gv")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	file.WriteString("graph BANetwork {\n")
	file.WriteString(sb.String())
	file.WriteString("}")
}

func Sample[T any](slice []T) T {
	return slice[rand.IntN(len(slice))]
}

func SampleNode(network *BANetwork) Node {
	nodes := []Node{}
	for node := range network.Nodes {
		nodes = append(nodes, node)
	}
	return Sample(nodes)
}

func SampleEdge(network *BANetwork) Edge {
	edges := []Edge{}
	for edge := range network.Edges {
		edges = append(edges, edge)
	}
	return Sample(edges)
}

func GetNode(network *BANetwork, nodeValue int) *Node {
	for node := range network.Nodes {
		if node.Value == nodeValue {
			return &node
		}
	}
	return nil
}

func GetEdges(network *BANetwork, node Node) []Edge {
	edges := []Edge{}
	for edge := range network.Edges {
		if edge.Left == node || edge.Right == node {
			edges = append(edges, edge)
		}
	}
	return edges
}

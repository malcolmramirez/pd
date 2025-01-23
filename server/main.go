package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/malcolmramirez/algorithm-server/v2/banetwork"
)

var network = banetwork.InitializeNetwork(3, 50)

func main() {
	listener, err := net.Listen("tcp", ":8080") // Listen on port 8080
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept() // Accept a connection

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn) // Handle connection in a goroutine
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer) // Read data from socket
	if err != nil {
		fmt.Println("Error reading from socket:", err)
		return
	}
	message := string(buffer[:n]) // Interpret bytes as a string
	message = message[:len(message)-2]

	fmt.Println("Received message:", message)
	response := "undefined"
	if strings.HasPrefix(message, "banetwork") {
		response = handleBaNetworkMessage(message)
	}

	n, err = conn.Write([]byte(response))
	fmt.Println("Wrote response:", response)
}

func handleBaNetworkMessage(message string) string {
	split := strings.Split(message, " ")
	if len(split) < 2 || len(split) > 4 {
		return "invalid format"
	}

	operation := split[1]
	if operation == "sampedge" {
		var node banetwork.Node
		if len(split) == 3 {
			i, err := strconv.Atoi(split[2])
			if err != nil {
				return "number format error"
			}
			nodePtr := banetwork.GetNode(network, i)
			if nodePtr == nil {
				return "node not found"
			}
			node = *nodePtr
		} else {
			node = banetwork.SampleNode(network)
		}

		edge := banetwork.Sample(banetwork.GetEdges(network, node))
		return fmt.Sprintf("%d %d;", edge.Left.Value, edge.Right.Value)
	}

	return "invalid format"
}

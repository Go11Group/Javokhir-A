package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	connections = make(map[net.Conn]bool)
	connMutex   sync.Mutex
)

func handleConnection(con net.Conn) {
	defer func() {
		connMutex.Lock()
		delete(connections, con)
		connMutex.Unlock()
		con.Close()
	}()

	receivedData := make([]byte, 1024)
	for {
		n, err := con.Read(receivedData)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		message := string(receivedData[:n])
		fmt.Println("Received:", message)

		broadcastMessage(con, message)
	}
}

func broadcastMessage(sender net.Conn, message string) {
	connMutex.Lock()
	defer connMutex.Unlock()

	for con := range connections {
		if con != sender {
			_, err := con.Write([]byte(message))
			if err != nil {
				fmt.Println("Error writing to connection:", err)
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port:", listener.Addr().String())

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		connMutex.Lock()
		connections[con] = true
		connMutex.Unlock()

		fmt.Println("Connection from:", con.RemoteAddr().String())

		go handleConnection(con)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message := make([]byte, 1024)

		n, err := conn.Read(message)
		if err != nil {
			fmt.Println("Reading from server failed:", err)
			return
		}
		fmt.Printf("\n%s:%s", conn.LocalAddr().String(), string(message[:n]))
	}
}

func main() {
	serverAddr := "localhost:8082"
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	go handleConnection(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Reading input failed:", err)
			break
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Writing to connection failed:", err)
			break
		}
	}
}

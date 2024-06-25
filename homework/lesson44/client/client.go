package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type UserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("client")

	client, err := rpc.DialHTTP("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	userReq := new(UserRequest)
	userReq.Id = 1
	UserData := new(UserData)

	if err := client.Call("UserServer.GetUser", userReq, &UserData); err != nil {
		log.Fatal(err)
	}

	fmt.Println(UserData)
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type UserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserServer struct {
}

var Users []CreateUser

func main() {
	server := new(UserServer)
	if err := rpc.Register(server); err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	http.Serve(l, nil)
}

func (s *UserServer) CreateUser(req *UserRequest, res *CreateUser) (*CreateUser, error) {
	id := len(Users) + 1

	res.Id = id

	newUser := &CreateUser{
		Id:   id,
		Name: req.Name,
		Age:  req.Age,
	}

	Users = append(Users, *newUser)

	fmt.Println("Users:", Users)

	res = newUser

	return newUser, nil
}

func (s *UserServer) GetUser(req *UserRequest, res *UserData) (*UserData, error) {

	for _, user := range Users {
		if user.Id == req.Id {
			res.Id = user.Id
			res.Name = user.Name
			res.Age = user.Age
			return (*UserData)(&user), nil
		}
	}

	return nil, fmt.Errorf("User %v not found", req.Id)
}

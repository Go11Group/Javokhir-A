package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	pb "github.com/Go11Group/Javokhir-A/atLesson/lesson45/protos/genproto/generator"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneratorServer
}

func (s *server) RandomPicker(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	numbers := []int{}
	for i := 1; i <= 24; i++ {
		if _, ok := in.Exception[int32(i)]; !ok {
			numbers = append(numbers, i)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	assignment := make(map[string]int32)
	for i, name := range in.Names {
		assignment[name] = int32(numbers[i])
	}

	return &pb.Response{Result: assignment}, nil
}

func (s *server) GetNameBySurname(ctx context.Context, UserName *pb.UserName) (*pb.UserSurname, error) {

	for _, name := range names {
		fullName := strings.Split(name, " ")
		if len(fullName) < 2 {
			return &pb.UserSurname{Surname: UserName.Name}, fmt.Errorf("User Name is not valid: ", fullName)
		}
		if UserName.Name == fullName[0] {
			return &pb.UserSurname{Surname: fullName[1]}, nil
		}
	}

	return nil, fmt.Errorf("Not found")
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := server{}
	grpc := grpc.NewServer()
	pb.RegisterGeneratorServer(grpc, &s)

	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}

var names = []string{
	"Abbos Qambarov",
	"Azizbek Qobulov",
	"Bekzod Qo'chqarov",
	"Diyorbek Nematov Dadajon o'g'li",
	"Faxriddin Raximberdiyev Farxodjon o'g'li",
	"Fazliddin Xayrullayev",
	"Hamidjon Nuriddinov",
	"Hamidulloh Hamidullayev",
	"Ibrohim Umarov Fazliddin o'g'li",
	"Jamshidbek Hatamov Erkin o'g'li",
	"Javohir Abdusamatov",
	"Muhammadaziz Yoqubov",
	"Muhammadjon Ko'palov",
	"Nurmuhammad",
	"Ozodjon A'zamjonov",
	"Sanjarbek Abduraxmonov",
	"Yusupov Bobur",
	"Firdavs",
	"Ozodbek",
	"Abdulaziz Xoliqulov",
	"Intizor opa",
}

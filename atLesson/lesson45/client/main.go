package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Go11Group/Javokhir-A/atLesson/lesson45/protos/genproto/generator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	gen := pb.NewGeneratorClient(conn)
	// req := &pb.Request{
	// 	Limit:     24,
	// 	Exception: map[int32]bool{3: true, 4: false, 5: true},
	// 	Names:     names,
	// }
	// resp, err := gen.RandomPicker(context.Background(), req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Seat - Name")
	// for k, v := range resp.Result {
	// 	fmt.Printf("%2d   - %s\n", v, k)
	// }
	var name string
	fmt.Println("Enter name:")

	fmt.Scan(&name)
	user := pb.UserName{Name: name}

	res, err := gen.GetNameBySurname(context.Background(), &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s, Surname: %s\n", user.Name, res.Surname)
}

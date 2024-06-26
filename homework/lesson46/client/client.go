package main

import (
	"context"
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson45/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ClientID string

func displayMenu(cl pb.LibraryClient) {
	for {
		fmt.Println("1. Add Book")
		fmt.Println("2. Search Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. See all books")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addBook(cl)
		case 2:
			searchBook(cl)
		case 3:
			borrowBook(cl)
		case 4:
			getAllBooks(cl)
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func main() {
	ClientID = uuid.NewString()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	cl := pb.NewLibraryClient(conn)

	fmt.Println("Welcome to Library Management System")
	displayMenu(cl)
}

func addBook(cl pb.LibraryClient) {
	req := pb.AddBookRequest{}
	fmt.Print("Enter book title: ")
	fmt.Scanln(&req.Title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&req.Author)
	fmt.Print("Enter book year published: ")
	fmt.Scanln(&req.YearPublished)

	resp, err := cl.AddBook(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Book added successfully: ID: " + resp.BookId)

}

func searchBook(cl pb.LibraryClient) {

	req := pb.SearchBookRequest{}
	fmt.Print("Enter book title: ")
	fmt.Scanln(&req.Title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&req.Author)
	fmt.Print("Enter book year published: ")
	fmt.Scanln(&req.YearPublished)

	resp, err := cl.SearchBook(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found Books:")

	for i, book := range resp.Books {
		fmt.Printf("%d-book: %v\n", i, book)
	}
}

func borrowBook(cl pb.LibraryClient) {
	req := pb.BorrowBookRequest{}

	req.UserId = ClientID

	fmt.Print("Enter book id: ")
	fmt.Scanln(&req.BookId)

	resp, err := cl.BorrowBook(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Book borrowed successfully", resp.Message)
}

func getAllBooks(cl pb.LibraryClient) {
	resp, err := cl.GetAllBooks(context.Background(), &pb.GetAllBooksRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("All books:")

	for i, book := range resp.Books {
		fmt.Printf("%d-book: %v\n", i, book)
	}
}

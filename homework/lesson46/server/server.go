package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"

	"github.com/Go11Group/Javokhir-A/homework/lesson46/pb"
	"google.golang.org/grpc"
)

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

var Users []User

var Books []*pb.Book //serves as a storage for books

type BorrowedBook struct {
	BookId string `json:"book_id"`
	UserId string `json:"user_id"`
}

var BorrowedBooks []*BorrowedBook // for borrowed

type LibraryService struct {
	pb.UnimplementedLibraryServer
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterLibraryServer(grpcServer, &LibraryService{})

	go func() {
		log.Println("Starting gRPC server on localhost:50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	log.Println("Stopping gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped gracefully")
}

func (ls *LibraryService) AddBook(c context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	newBookId := uuid.NewString()

	book := pb.Book{
		BookId:        newBookId,
		Title:         req.Title,
		Author:        req.Author,
		YearPublished: req.YearPublished,
	}

	Books = append(Books, &book)

	return &pb.AddBookResponse{BookId: newBookId}, nil
}

func (ls *LibraryService) SearchBook(c context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	foundBooks := []*pb.Book{}

	for _, book := range Books {
		if (req.Author != "" && book.Author == req.Author) ||
			(req.Title != "" && book.Title == req.Title) ||
			(req.YearPublished != "" && book.YearPublished == req.YearPublished) {
			foundBooks = append(foundBooks, book)
		}
	}

	return &pb.SearchBookResponse{Books: foundBooks}, nil
}

func (ls *LibraryService) BorrowBook(c context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	var foundUser *User

	for _, user := range Users {
		if user.UserId == req.UserId {
			foundUser = &user
		}
	}

	if foundUser != nil {
		return &pb.BorrowBookResponse{Message: false}, fmt.Errorf("User %v not found", req.UserId)
	}

	borrowingBook := BorrowedBook{
		UserId: req.UserId,
	}

	for i, book := range Books {
		if book.BookId == req.BookId {
			borrowingBook.BookId = book.BookId

			//remove book from books list
			Books = append(Books[:i], Books[i+1:]...)

			//add book to borrowed books list
			BorrowedBooks = append(BorrowedBooks, &borrowingBook)

			return &pb.BorrowBookResponse{Message: true}, nil
		}
	}

	return &pb.BorrowBookResponse{Message: false}, fmt.Errorf("Book %v not found", req.BookId)
}

func (ls *LibraryService) GetAllBooks(c context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	return &pb.GetAllBooksResponse{Books: Books}, nil
}

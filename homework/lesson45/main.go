package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/Go11Group/Javokhir-A/homework/lesson45/grpc/translator"
	"google.golang.org/grpc"
)

type Word struct {
	EnWord string `json:"en_word,omitempty"`
	UzWord string `json:"uz_word,omitempty"`
}

var words []Word

type server struct {
	translator.UnimplementedTranslatorServiceServer
}

func (s *server) Translator(ctx context.Context, req *translator.TranslatingRequest) (*translator.TranslatingResponse, error) {
	log.Printf("Received: %v", req)
	en := translate(req.UzWord)
	return &translator.TranslatingResponse{EnWord: en}, nil
}

func translate(uzWord string) string {

	file, err := os.Open("resources/words.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	if err := json.NewDecoder(file).Decode(&words); err != nil {
		log.Fatal(err)
	}

	for _, word := range words {
		if word.UzWord == uzWord {
			return word.EnWord
		}
	}
	return "Translated: " + uzWord
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	translator.RegisterTranslatorServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

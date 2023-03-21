package main

import (
	"context"
	"fmt"
	"grpc-test/pb"
	"grpc-test/repository"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BookServiceClient struct {
	pb.UnimplementedBookServiceServer
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string
	LastName  string
}

// var books []Book
var Book2 Book

func (c *BookServiceClient) GetBooks(ctx context.Context, req *pb.Book) (*pb.Books, error) {
	fmt.Println("GetBooks was invoked")

	repo := repository.NewRepository()
	res, err := repo.FindAllBooks(ctx, req)
	fmt.Println("res:", res)
	if err != nil {
		panic("error")
	}

	return res, nil
}

func (c *BookServiceClient) GetBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	fmt.Println("GetBook was invoked")

	repo := repository.NewRepository()
	res, err := repo.FindBookById(ctx, req)
	fmt.Println("res:", res)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("res2:", res)
	return res, nil
}

func (c *BookServiceClient) CreateBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	fmt.Println("CreateBook was invoked")
	fmt.Println("req:", req)
	repo := repository.NewRepository()
	res, err := repo.AddBook(ctx, req)
	fmt.Println("res:", res)
	if err != nil {
		panic("error")
	}

	return res, nil
}

func (c *BookServiceClient) UpdateBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	fmt.Println("UpdateBook was invoked")
	fmt.Println("req:", req)
	repo := repository.NewRepository()
	res, err := repo.EditBook(ctx, req)
	fmt.Println("res:", res)
	if err != nil {
		panic("error")
	}

	return res, nil
}

func (c *BookServiceClient) DeleteBook(ctx context.Context, req *pb.Book) (*pb.Books, error) {
	fmt.Println("DeleteBook was invoked")
	fmt.Println("req:", req)
	repo := repository.NewRepository()
	res, err := repo.EliminateBook(ctx, req)
	fmt.Println("res:", res)
	if err != nil {
		panic("error")
	}

	return res, nil
}

func main() {

	// con, er := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8mb4")
	// if er != nil {
	// 	panic(er)
	// }
	// // defer con.Close() //DBの開放処理
	// boil.SetDB(con)

	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("failed to list:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &BookServiceClient{})
	fmt.Println("server is running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve:%v", err)
	}
}

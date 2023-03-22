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

// BookServerという構造体を定義
type BookServer struct {
	pb.UnimplementedBookServiceServer //gRPCサーバー実装時にUnimplementedBookServiceServerを埋め込んだ構造体を作る必要ある。
	//こちらの構造体下記メソッドが定義されている。
}

// サーバー側のメソッドを定義
func (c *BookServer) GetBooks(ctx context.Context, req *pb.Book) (*pb.Books, error) {
	fmt.Println("GetBooks was invoked")

	repo := repository.NewRepository()
	res, err := repo.FindAllBooks(ctx, req)
	fmt.Println("res:", res)
	if err != nil {
		fmt.Println("err:", err)
	}

	return res, nil
}

func (c *BookServer) GetBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
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

func (c *BookServer) CreateBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
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

func (c *BookServer) UpdateBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
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

func (c *BookServer) DeleteBook(ctx context.Context, req *pb.Book) (*pb.Books, error) {
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

	lis, err := net.Listen("tcp", "localhost:8082") //tcp通信を設定。
	if err != nil {
		log.Fatalf("failed to list:%v", err)
	}
	s := grpc.NewServer()                          //gRPCサーバーを作成。sはServer構造体。
	pb.RegisterBookServiceServer(s, &BookServer{}) //gRPCサーバーにサービスを登録→sにBookServerを登録。
	fmt.Println("server is running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve:%v", err)
	}
}

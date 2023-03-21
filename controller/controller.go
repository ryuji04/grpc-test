package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-test/pb"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (c Controller) Execute() {
	fmt.Println("Execute")
	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail connect:%v", err)
	}
	defer conn.Close()
	client := pb.NewBookServiceClient(conn)

	router.GET("/api/books", func(c *gin.Context) {
		getBooks(ctx, c, client)
	})
	router.GET("/api/book/:id", func(c *gin.Context) {
		getBook(ctx, c, client)
	})
	router.POST("api/book/create", func(c *gin.Context) {
		createBook(ctx, c, client)
	})
	router.POST("api/book/update", func(c *gin.Context) {
		updateBook(ctx, c, client)
	})
	router.POST("api/book/delete", func(c *gin.Context) {
		deleteBook(ctx, c, client)
	})

	router.Run(":8080")
}

func getBooks(ctx context.Context, c *gin.Context, client pb.BookServiceClient) {
	fmt.Println("getBooks")
	BookPb := pb.Book{}
	res, err := client.GetBooks(context.Background(), &BookPb)

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res.BookList)
	if err != nil {
		fmt.Println(err)
	}
}
func getBook(ctx context.Context, c *gin.Context, client pb.BookServiceClient) {
	fmt.Println("getBook")
	id := c.Param("id")
	BookPb := pb.Book{
		ID: id,
	}
	res, err := client.GetBook(context.Background(), &BookPb)

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res)
	if err != nil {
		fmt.Println(err)
	}
}

func createBook(ctx context.Context, c *gin.Context, client pb.BookServiceClient) {
	fmt.Println("crateBook")
	BookPb := pb.Book{}
	if err := c.ShouldBindJSON(&BookPb); err != nil {
		panic(err)
	}
	fmt.Println("BookPb:", &BookPb)
	res, err := client.CreateBook(context.Background(), &BookPb)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res)
	if err != nil {
		fmt.Println(err)
	}
}

func updateBook(ctx context.Context, c *gin.Context, client pb.BookServiceClient) {
	fmt.Println("crateBook")
	BookPb := pb.Book{}
	if err := c.ShouldBindJSON(&BookPb); err != nil {
		panic(err)
	}
	fmt.Println("BookPb:", &BookPb)
	res, err := client.UpdateBook(context.Background(), &BookPb)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteBook(ctx context.Context, c *gin.Context, client pb.BookServiceClient) {
	fmt.Println("crateBook")
	BookPb := pb.Book{}
	if err := c.ShouldBindJSON(&BookPb); err != nil {
		panic(err)
	}
	fmt.Println("BookPb:", &BookPb)
	res, err := client.DeleteBook(context.Background(), &BookPb)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res)
	if err != nil {
		fmt.Println(err)
	}
}

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
	//Controllerの構造体を返す
	return Controller{}
}

func (c Controller) Execute() {
	fmt.Println("Execute")

	ctx := context.Background() //空のContextを生成

	gin.SetMode(gin.ReleaseMode) //Ginのデバックモードを無効にする
	router := gin.Default()

	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure()) //指定したターゲット(localhost:8082)へ接続を作成している。 rpc.WithInsecure()は「認証をしないで接続する」という指定をしている。
	if err != nil {
		log.Fatalf("fail connect:%v", err)
	}
	defer conn.Close()                      //一通りの処理がclient側で完了したら接続を閉じる
	client := pb.NewBookServiceClient(conn) //gRPCクライアントを生成している(gRPCクライアントを「localhost:8082」に接続)

	//パスとメソッドを紐づけている。
	//gRPC通信をする為に引数には必ずgRPCクライアントを持たせている。

	//HTTP通信でGETメソッドを指定している
	router.GET("/api/books", func(c *gin.Context) {
		getBooks(ctx, c, client) //gin.Contextを引数に持たせてメソッドを呼び出している
	})
	router.GET("/api/book/:id", func(c *gin.Context) {
		//パラメーターからidを取得する
		getBook(ctx, c, client)
	})
	//HTTP通信でPOSTメソッドを指定している
	router.POST("api/book/create", func(c *gin.Context) {
		createBook(ctx, c, client)
	})
	router.POST("api/book/update", func(c *gin.Context) {
		updateBook(ctx, c, client)
	})
	router.POST("api/book/delete", func(c *gin.Context) {
		deleteBook(ctx, c, client)
	})

	router.Run(":8080") //ルーターをhttp.Serverに接続
}

func getBooks(ctx context.Context, c *gin.Context, client pb.BookServiceClient) {
	fmt.Println("getBooks")
	BookPb := pb.Book{}
	res, err := client.GetBooks(context.Background(), &BookPb) //ここでサーバ側へ接続している

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &res)                            //サーバから正常にレスポンスが返ってきた場合JSON形式で結果が表示される。→POSTMANを使用した際に戻り値がJSONで表示される。
	err = json.NewEncoder(os.Stdout).Encode(&res.BookList) //pb.Book構造体をJSONに変換。「os.Stdout」→コンソール上で結果(JSON形式)を表示。
	if err != nil {
		fmt.Println(err)
	}
}

//以下同様

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
	fmt.Println("updateBook")
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

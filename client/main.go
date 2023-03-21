package main

import (
	"grpc-test/controller"
)

func main() {
	cont := controller.NewController()
	cont.Execute()
}

// func main() {
// 	gin.SetMode(gin.ReleaseMode)

// 	router := gin.Default()
// 	router.GET("/api/books", getBooks)
// 	fmt.Println("serve")
// 	router.Run(":8081")
// }

// func getBooks(c *gin.Context) {

// 	fmt.Println("test")

// 	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("fail connect:%v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewBookServiceClient(conn)

// 	BookPb := pb.Book{
// 		ID:     "1",
// 		Title:  "Book one",
// 		Author: &pb.Author{FirstName: "Philip", LastName: "Williams"},
// 	}
// 	res, err := client.GetBooks(context.Background(), &BookPb)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// err = json.NewEncoder(os.Stdout).Encode(&res.BookList)
// json.Unmarshal([]byte(j), &decoded)
// if err := json.NewDecoder(strings.NewReader(j)).Decode(&Book2); err != nil {
// 	fmt.Println(err)
// 	return
// }
//json.NewEncoder(os.Stdout).Encode(&Book2)
// json.NewEncoder(w).Encode(books)
// }

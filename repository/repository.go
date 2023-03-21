package repository

import (
	"context"
	"database/sql"
	"fmt"
	"grpc-test/pb"
	"log"

	"grpc-test/config/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	//"github.com/vattle/sqlboiler/boil"
)

func init() {
	con, er := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8mb4")
	if er != nil {
		panic(er)
	}
	// defer con.Close() //DBの開放処理
	boil.SetDB(con)
}

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) FindAllBooks(ctx context.Context, req *pb.Book) (*pb.Books, error) {

	// con, er := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8mb4")
	// if er != nil {
	// 	panic(er)
	// }
	//  defer con.Close() //DBの開放処理
	// boil.SetDB(con)

	var PbBooks = &pb.Books{}
	var PbBook = &pb.Book{}

	Books, err := models.Books().AllG(ctx)
	log.Print("Books:", Books)
	for _, Book := range Books {
		PbBook = &pb.Book{
			ID:    Book.ID,
			Title: Book.Title.String,
		}
		PbBooks = &pb.Books{
			BookList: append(PbBooks.BookList, PbBook),
		}
	}

	return PbBooks, err
}

func (c Repository) FindBookById(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	var PbBook = &pb.Book{}
	Book, err := models.FindBookG(ctx, req.ID)
	if err != nil {
		return PbBook, err
	}
	PbBook = &pb.Book{
		ID:    Book.ID,
		Title: Book.Title.String,
	}
	return PbBook, err
}

func (c Repository) AddBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	model := &models.Book{
		ID:    req.ID,
		Title: null.StringFrom(req.Title),
	}
	fmt.Println("req.ID:", req.ID, "req.Title:", req.Title)
	err := model.InsertG(ctx, boil.Infer())
	if err != nil {
		panic(err)
	}
	Book, err := models.FindBookG(ctx, req.ID)
	var PbBook = &pb.Book{
		ID:    Book.ID,
		Title: Book.Title.String,
	}
	return PbBook, err
}

func (c Repository) EditBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	Book, err := models.FindBookG(ctx, req.ID)
	if err != nil {
		panic(err)
	}

	Book.ID = req.ID
	Book.Title = null.StringFrom(req.Title)

	rowsAff, err := Book.UpdateG(ctx, boil.Infer())
	fmt.Println("rowAff:", rowsAff)

	var PbBook = &pb.Book{
		ID:    Book.ID,
		Title: Book.Title.String,
	}
	return PbBook, err
}

func (c Repository) EliminateBook(ctx context.Context, req *pb.Book) (*pb.Books, error) {
	Book, err := models.FindBookG(ctx, req.ID)
	if err != nil {
		panic(err)
	}

	rowsAff, err := Book.DeleteG(ctx)

	fmt.Println("rowAff:", rowsAff)

	var PbBooks = &pb.Books{}
	var PbBook = &pb.Book{}

	Books, err := models.Books().AllG(ctx)
	log.Print("Books:", Books)
	for _, Book := range Books {
		PbBook = &pb.Book{
			ID:    Book.ID,
			Title: Book.Title.String,
		}
		PbBooks = &pb.Books{
			BookList: append(PbBooks.BookList, PbBook),
		}
	}

	return PbBooks, err

}

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

// init()は必ず最初に実行されるメソッド
func init() {
	con, er := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8mb4") //DBコネクションを設定
	if er != nil {
		panic(er)
	}
	boil.SetDB(con) //コネクション情報をグローバル領域に保存
}

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) FindAllBooks(ctx context.Context, req *pb.Book) (*pb.Books, error) {

	var PbBooks = &pb.Books{} //空のpb.Books構造体を作成。※構造体を生成する際は{}が必要。
	var PbBook = &pb.Book{}

	Books, err := models.Books().AllG(ctx) //sqlboilerのメソッドはAllGにする事で引数にDB接続情報を入れなくてよくなる。
	if err != nil {
		return PbBooks, err
	}
	//DBから全件取得したBooks情報(型:models.BookSlice)を戻り値であるPbBooks(型:pb.Books)に格納する
	for _, Book := range Books {
		PbBook = &pb.Book{
			ID:    Book.ID,
			Title: Book.Title.String,
		}
		PbBooks = &pb.Books{
			BookList: append(PbBooks.BookList, PbBook), //PbBooks.BookListにPbBookを格納
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
	var PbBook = &pb.Book{}
	model := &models.Book{
		ID:    req.ID,
		Title: null.StringFrom(req.Title),
	}
	fmt.Println("req.ID:", req.ID, "req.Title:", req.Title)
	err := model.InsertG(ctx, boil.Infer())
	if err != nil {
		return PbBook, err
	}
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

func (c Repository) EditBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	var PbBook = &pb.Book{}
	Book, err := models.FindBookG(ctx, req.ID)
	if err != nil {
		return PbBook, err
	}

	Book.ID = req.ID
	Book.Title = null.StringFrom(req.Title)

	rowsAff, err := Book.UpdateG(ctx, boil.Infer()) //rowsAff:アップデート件数
	fmt.Println("rowAff:", rowsAff)

	PbBook = &pb.Book{
		ID:    Book.ID,
		Title: Book.Title.String,
	}
	return PbBook, err
}

func (c Repository) EliminateBook(ctx context.Context, req *pb.Book) (*pb.Books, error) {
	var PbBooks = &pb.Books{}
	var PbBook = &pb.Book{}
	Book, err := models.FindBookG(ctx, req.ID)
	if err != nil {
		return PbBooks, err
	}

	rowsAff, err := Book.DeleteG(ctx)

	fmt.Println("rowAff:", rowsAff)

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

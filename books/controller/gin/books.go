package gin

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	mysql "github.com/Mictrlan/library/books/model/mysql"
	"github.com/gin-gonic/gin"
)

var (
	errInvalidServer = errors.New("[RegisterRouter]: server is nil")
)

// BooksController -
type BooksController struct {
	db *sql.DB
}

// New create new bookController
func New(db *sql.DB) *BooksController {
	return &BooksController{
		db: db,
	}
}

// Register register books router
func (bc *BooksController) Register(r gin.IRouter) error {

	if r == nil {
		log.Fatal(errInvalidServer)
	}

	err := mysql.BooksCreateTable(bc.db)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("api/v2/books/booksadd", bc.booksadd)
	r.POST("api/v2/books/bookQueryByID", bc.booksInquireByID)
	r.POST("api/v2/books/bookQueryByISBN", bc.booksInquireByISBN)
	r.POST("api/v2/books/bookQueryByTitle", bc.booksInquireByTitle)
	r.POST("api/v2/books/bookQueryByAuthor", bc.booksInquireByAuthor)
	r.POST("api/v2/books/booksUpdateTotalByISBN", bc.booksUpdateTotal)
	r.POST("api/v2/books/booksUpdateCountByISBN", bc.booksUpdateCountByISBN)
	r.POST("api/v2/books/booksDeleteByISBN", bc.booksDeleteByISBN)

	return nil
}

func (bc *BooksController) booksadd(ctx *gin.Context) {
	var (
		req struct {
			Title    string `json:"title"`
			ISBN     string `json:"ISBN"`
			Author   string `json:"author"`
			Duration string `json:"duration"`
			Total    int    `json:"total"`
			Ebook    bool   `json:"Ebook"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadGateway})
		return
	}

	id, err := mysql.BooksAddInfo(bc.db, req.Title, req.ISBN, req.Author, req.Total, req.Total, req.Duration, req.Ebook)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"id":     id,
	})
}

func (bc *BooksController) booksInquireByID(ctx *gin.Context) {
	var (
		book *mysql.Book

		req struct {
			ID int64 `json:"id"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadGateway})
		return
	}

	book, err = mysql.BooksInquireByID(bc.db, req.ID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"book":   book,
	})
}

func (bc *BooksController) booksInquireByISBN(ctx *gin.Context) {
	var (
		book *mysql.Book

		req struct {
			ISBN string `json:"ISBN"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadGateway})
		return
	}

	book, err = mysql.BooksInquireByISBN(bc.db, req.ISBN)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"book":   book,
	})
}

func (bc *BooksController) booksInquireByTitle(ctx *gin.Context) {
	var (
		req struct {
			Title string `json:"title"`
		}

		books []*mysql.Book
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	books, err = mysql.BooksInquireByTitle(bc.db, req.Title)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"books":  books,
	})
}

func (bc *BooksController) booksInquireByAuthor(ctx *gin.Context) {
	var (
		req struct {
			Author string `json:"author"`
		}

		books []*mysql.Book
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	books, err = mysql.BooksInquireByAuthor(bc.db, req.Author)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"books":  books,
	})
}

func (bc *BooksController) booksUpdateTotal(ctx *gin.Context) {
	var (
		req struct {
			ISBN  string `json:"ISBN"`
			Value int    `json:"total"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.BooksUpdateTotal(bc.db, req.Value, req.ISBN)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	err = mysql.BooksUpdateInCountByISBN(bc.db, req.Value, req.ISBN)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionRequired, gin.H{"status": http.StatusPreconditionRequired})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (bc *BooksController) booksUpdateCountByISBN(ctx *gin.Context) {
	var (
		req struct {
			Value int    `json:"OutCount"`
			ISBN  string `json:"ISBN"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.BooksUpdateOutCountByISBN(bc.db, req.Value, req.ISBN)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	err = mysql.BooksUpdateInCountReduce(bc.db, req.Value, req.ISBN)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionRequired, gin.H{"status": http.StatusPreconditionRequired})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (bc *BooksController) booksDeleteByISBN(ctx *gin.Context) {
	var (
		req struct {
			ISBN string `json:"ISBN"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.BooksDeleteByISBN(bc.db, req.ISBN)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

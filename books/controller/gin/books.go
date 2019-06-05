package gin

import (
	"database/sql"
	"errors"
	"log"

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

	return nil
}

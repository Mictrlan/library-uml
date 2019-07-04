package model

import (
	"database/sql"
	"errors"
)

const (
	mysqlBooksCreateBooksTable = iota
	mysqlBooksInsertBookInfo
	mysqlBooksInquireByISBN
	mysqlBooksInquireByTitle
	mysqlBooksInquireByAuthor
	mysqlBooksUpdateTotalByISBN
	mysqlBooksModifyInCountPlus
	mysqlBooksModifyInCountReduce
	mysqlBooksModifyOutCountPlus
	mysqlBooksModifyOutCountReduce
	mysqlBooksModifyState
	mysqlBooksDeleteBookByISBN
)

// Book contain book info
type Book struct {
	title   string
	ISBN    string
	author  string
	inCount int
}

var (
	errInvalidInsert = errors.New("Invalid Insert: insert affected 0 rows")
	errInvalidModify = errors.New("Invalid Update: modify affected 0 rows")

	booksSQLString = []string{
		`CREATE TABLE IF NOT EXISTS books(
			id                 INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT, 
			title              VARCHAR(50) NOT NULL,
			ISBN               VARCHAR(50) UNIQUE NOT NULL,
			author             VARCHAR(50) NOT NULL,
			duration           VARCHAR(50) NOT NULL DEFAULT '30',
			state              VARCHAR(50) DEFAULT 'Borrowable',
			startdate          DATETIME DEFAULT '9102-6-15 00:00:00', 
			enddate            DATETIME DEFAULT '9102-6-15 00:00:00', 
			total              INT NOT NULL,
			outCount           INT NOT NULL DEFAULT 0,
			inCount            INT NOT NULL DEFAULT 0,
			PRIMARY KEY (id)
		)ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='books information'`,
		`INSERT INTO books (title,ISBN,author,duration,total,inCount,Ebook)VALUES(?,?,?,?,?,?,?)`,
		`SELECT title,ISBN,author,inCount,Ebook FROM books WHERE ISBN = ? LOCK IN SHARE MODE`,
		`SELECT title,ISBN,author,inCount,Ebook FROM books WHERE title = ? LOCK IN SHARE MODE`,
		`SELECT title,ISBN,author,inCount,Ebook FROM books WHERE author = ? LOCK IN SHARE MODE`,
		`UPDATE books SET total = total + ? WHERE ISBN = ? LIMIT 1 `,
		`UPDATE books SET inCount = inCount + ? WHERE ISBN = ? LIMIT 1 `,
		`UPDATE books SET inCount = inCount - ? WHERE ISBN = ? LIMIT 1 `,
		`UPDATE books SET outCount = outCount + ? WHERE ISBN = ? LIMIT 1 `,
		`UPDATE books SET outCount = outCount - ? WHERE ISBN = ? LIMIT 1 `,
		`UPDATE books SET state = ? WHERE ISBN = ? LIMIT 1`,
		`DELETE FROM books WHERE ISBN = ? LIMIT 1 `,
	}
)

// BooksCreateTable create books table
func BooksCreateTable(db *sql.DB) error {
	_, err := db.Exec(booksSQLString[mysqlBooksCreateBooksTable])
	return err
}

// BooksAddInfo insert new book info
func BooksAddInfo(db *sql.DB, title, ISBN, author string, total, inCount int, duration string, Ebook bool) (int64, error) {
	result, err := db.Exec(booksSQLString[mysqlBooksInsertBookInfo], title, ISBN, author, duration, total, inCount, Ebook)
	if err != nil {
		return 0, err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, errInvalidInsert
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// BooksInquireByISBN query book info by ISBN
func BooksInquireByISBN(db *sql.DB, ISBN string) (*Book, error) {
	var bookInfo Book

	err := db.QueryRow(booksSQLString[mysqlBooksInquireByISBN], ISBN).Scan(&bookInfo.title, &bookInfo.ISBN, &bookInfo.author, &bookInfo.inCount)
	if err != nil {
		return nil, err
	}

	return &bookInfo, nil
}

// BooksInquireByTitle query books info by title
func BooksInquireByTitle(db *sql.DB, title string) ([]*Book, error) {
	var (
		bookInfo  Book
		booksInfo []*Book
	)

	rows, err := db.Query(booksSQLString[mysqlBooksInquireByAuthor], title)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bookInfo.title, &bookInfo.ISBN, &bookInfo.author, &bookInfo.inCount)
		if err != nil {
			return nil, err
		}

		booksInfo = append(booksInfo, &bookInfo)
	}

	return booksInfo, nil
}

//BooksInquireByAuthor query books info by author
func BooksInquireByAuthor(db *sql.DB, author string) ([]*Book, error) {
	var (
		bookInfo  Book
		booksInfo []*Book
	)

	rows, err := db.Query(booksSQLString[mysqlBooksInquireByAuthor], author)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bookInfo.title, &bookInfo.ISBN, &bookInfo.author, &bookInfo.inCount)
		if err != nil {
			return nil, err
		}

		booksInfo = append(booksInfo, &bookInfo)
	}

	return booksInfo, nil
}

// BooksUpdateTotal update total by ISBN
func BooksUpdateTotal(db *sql.DB, value int, ISBN string) error {
	result, err := db.Exec(booksSQLString[mysqlBooksUpdateTotalByISBN], value, ISBN)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// BooksUpdateInCountByISBN update Count by ISBN
func BooksUpdateInCountByISBN(db *sql.DB, value int, ISBN string) error {
	result, err := db.Exec(booksSQLString[mysqlBooksModifyInCountPlus], value, ISBN)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// BooksUpdateInCountReduce update Count by ISBN
func BooksUpdateInCountReduce(db *sql.DB, value int, ISBN string) error {
	result, err := db.Exec(booksSQLString[mysqlBooksModifyInCountReduce], value, ISBN)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// BooksUpdateOutCountByISBN update Count by ISBN
func BooksUpdateOutCountByISBN(db *sql.DB, value int, ISBN string) error {
	result, err := db.Exec(booksSQLString[mysqlBooksModifyOutCountPlus], value, ISBN)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// BooksUpdateOutCountReduce update Count by ISBN
func BooksUpdateOutCountReduce(db *sql.DB, value int, ISBN string) error {
	result, err := db.Exec(booksSQLString[mysqlBooksModifyOutCountReduce], value, ISBN)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// BooksStateUpdateByISBN modify book state
func BooksStateUpdateByISBN(db *sql.DB, value bool, ISBN string) error {
	result, err := db.Exec(booksSQLString[mysqlBooksModifyState], value, ISBN)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// BooksDeleteByISBN delete book by ISBN
func BooksDeleteByISBN(db *sql.DB, ISBN string) error {
	_, err := db.Exec(booksSQLString[mysqlBooksDeleteBookByISBN], ISBN)
	return err
}

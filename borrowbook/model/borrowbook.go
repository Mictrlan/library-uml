package model

import (
	"database/sql"
	"errors"
	"time"
)

const (
	mysqlBorrowerBookCreateTable = iota
	mysqlBorrowerBookaddInfo
	mysqlBorrowerBookQueryByName
	mysqlBorrowerBookdelete
)

// Book contain book info
type Book struct {
	title     string
	ISBN      string
	author    string
	startdate string
}

var (
	errInvalidInsert = errors.New("Invalid Insert: insert affected 0 rows")
	errInvalidModify = errors.New("Invalid Update: modify affected 0 rows")

	borrowerbookSQLString = []string{
		`CREATE TABLE IF NOT EXISTS borrowerbook(
			id                 INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT, 
			name                VARCHAR(50) NOT NULL,
			title              VARCHAR(50) NOT NULL,
			author             VARCHAR(50) NOT NULL,
			startdate          DATETIME DEFAULT '9102-6-15 00:00:00', 
			PRIMARY KEY (id)
		)ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT= borrowerbook information'`,
		`INSERT INTO borrowerbook (name,title,author,startdate)VALUES(?,?,?,?)`,
		`SELECT title,ISBN,author,startdate FROM borrowerbook WHERE name = ? LOCK IN SHARE MODE`,
		`DELETE FROM borrowerbook WHERE name = ? LIMIT 1 `,
	}
)

// BorrowerBookCreateTable -
func BorrowerBookCreateTable(db *sql.DB) error {
	_, err := db.Exec(borrowerbookSQLString[mysqlBorrowerBookCreateTable])
	return err
}

// BorrowerBookAddInfo -
func BorrowerBookAddInfo(db *sql.DB, name, title, ISBN string, startdate time.Time) error {
	result, err := db.Exec(borrowerbookSQLString[mysqlBorrowerBookaddInfo], name, title, ISBN, startdate)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidInsert
	}

	return nil
}

// BorrowerBookQueryByName -
func BorrowerBookQueryByName(db *sql.DB, name string) ([]*Book, error) {
	var (
		bookInfo  Book
		booksInfo []*Book
	)

	rows, err := db.Query(borrowerbookSQLString[mysqlBorrowerBookQueryByName], name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bookInfo.title, &bookInfo.ISBN, &bookInfo.author, &bookInfo.startdate)
		if err != nil {
			return nil, err
		}

		booksInfo = append(booksInfo, &bookInfo)
	}

	return booksInfo, nil

}

// BorrowerBookDelete -
func BorrowerBookDelete(db *sql.DB, name string) error {
	_, err := db.Exec(borrowerbookSQLString[mysqlBorrowerBookdelete], name)
	return err
}

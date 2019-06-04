package model

const (
	mysqlBooksCreateBooksTable = iota
	mysqlBooksInsertBook
	mysqlBooksInquireByID
	mysqlBooksInquireByISBN
	mysqlBooksInquireByTitle
	mysqlBooksInquireByAuthor
	mysqlBooksModifyActive

	mysqlBooksDeleteBook
)

type book struct {
	title    string
	ISBN     string
	author   string
	total    int
	outCount int
	inCount  int
}

var (
	booksSQLString = []string{
		`CREATE TABLE IF NOT EXISTS books(
			id                 INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT, 
			title              VARCHAR(50) NOT NULL,
			ISBN               VARCHAR(50) UNIQUE NOT NULL,
			author             VARCHAR(50) NOT NULL,
			publishingHouse    VARCHAR(50),
			duration           DATETIME NOT NULL,
			startdate          DATETIME DEFAULT '0000-00-00 00:00:00', 
			enddate            DATETIME DEFAULT '0000-00-00 00:00:00', 
			totalCount         INt NOt NULL,
			outCount           INT NOT NULL DEFAULT '0',
			inCount            INT NOT NULL DEFAULT '0',
			PRIMARY KEY (id),
			UNIQUE KEY ISBN(ISBN),
		),ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='books information'`,
		`INSERT INTO books (title,ISBN,publishingHouse,duration,totalCount)VALUES(?,?,?,?,?)`,
		`SELECT * FROM books WHERE id = ? LOCK IN SHARE MODE`,
		`SELECT * FROM books WHERE ISBN = ? LOCK IN SHARE MODE`,
		`SELECT * FROM books WHERE title = ? LOCK IN SHARE MODE`,
		`SELECT * FROM books WHERE author = ? LOCK IN SHARE MODE`,
		`UPDATE books SET outCount = outCount + ? LIMIT 1 `,
	}
)

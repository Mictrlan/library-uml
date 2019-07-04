package model

import (
	"database/sql"
	"errors"
)

const (
	mysqlLibrarainCreateTable = iota
	mysqlLibrarainModify
	mysqlLibrarainAddInfo
	mysqlLibrarainDelete
)

var (
	errInvalidInsert = errors.New("Invalid Insert: insert affected 0 rows")
	errInvalidModify = errors.New("Invalid Update: modify affected 0 rows")

	librarainSQLString = []string{
		`CREATE TABLE IF NOT EXISTS librarain(
			id                 INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT, 
			name                VARCHAR(50) NOT NULL,
			pwd                 VARCHAR(512) NOT NULL ,
			PRIMARY KEY (id)
		)ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT= librarain information'`,
		`INSERT INTO librarain (name,pwd)VALUES(?,?)`,
		`UPDATE pwd SET pwd = ? WHERE name = ? LIMIT 1`,
		`DELETE FROM librarain WHERE name = ? LIMIT 1 `,
	}
)

// LibrarainCreateTable -
func LibrarainCreateTable(db *sql.DB) error {
	_, err := db.Exec(librarainSQLString[mysqlLibrarainCreateTable])
	return err
}

// LibrarainAddInfo -
func LibrarainAddInfo(db *sql.DB, name, pwd string) (int64, error) {
	result, err := db.Exec(librarainSQLString[mysqlLibrarainAddInfo], name, pwd)
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

// LibrarainModifyPassword -
func LibrarainModifyPassword(db *sql.DB, name, pwd string) error {
	_, err := db.Exec(librarainSQLString[mysqlLibrarainModify], pwd, name)
	return err
}

// LibrarainDelete -
func LibrarainDelete(db *sql.DB, id int64) error {
	_, err := db.Exec(librarainSQLString[mysqlLibrarainDelete], id)
	return err
}

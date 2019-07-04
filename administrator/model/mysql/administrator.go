package model

import (
	"database/sql"
	"errors"
)

const (
	mysqlAdminCreateTable = iota
	mysqlAdminModify
	mysqlAdminAddInfo
	mysqlAdminDelete
)

var (
	errInvalidInsert = errors.New("Invalid Insert: insert affected 0 rows")
	errInvalidModify = errors.New("Invalid Update: modify affected 0 rows")

	administratorSQLString = []string{
		`CREATE TABLE IF NOT EXISTS administrator(
			id                 INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT, 
			name                VARCHAR(50) NOT NULL,
			pwd                 VARCHAR(512) NOT NULL DEFAULT '123456' ,
			PRIMARY KEY (id)
		)ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT= administrator information'`,
		`INSERT INTO administrator (name,pwd)VALUES(?,?)`,
		`UPDATE pwd SET pwd = ? WHERE name = ? LIMIT 1`,
		`DELETE FROM administrator WHERE name = ? LIMIT 1 `,
	}
)

// AdministratorCreatetable -
func AdministratorCreatetable(db *sql.DB) error {
	_, err := db.Exec(administratorSQLString[mysqlAdminCreateTable])
	return err
}

// AdministratorAdd -
func AdministratorAdd(db *sql.DB, name, pwd string) error {
	_, err := db.Exec(administratorSQLString[mysqlAdminAddInfo], name, pwd)
	return err
}

// AdministratorModify -
func AdministratorModify(db *sql.DB, name, pwd string) error {
	_, err := db.Exec(administratorSQLString[mysqlAdminModify], pwd, name)
	return err
}

// AdministratorDelete -
func AdministratorDelete(db *sql.DB, name string) error {
	_, err := db.Exec(administratorSQLString[mysqlAdminDelete], name)
	return err
}

package mysql

import (
	"database/sql"
	"errors"

	config "github.com/Mictrlan/library/students/config"
)

const (
	mysqlStudentsCreateTable = iota
	mysqlStudentsInsertInfo
	mysqlStudentsInquireIDByName
	mysqlStudentsModifyMobile
	mysqlStudentInquirePwd
	mysqlStudentsModifyPassword
	mysqlStudentsDeleteByID
)

var (
	errInvalidInsert  = errors.New("Invalid Insert: insert affected 0 rows")
	errInvalidModify  = errors.New("Invalid Update: modify affected 0 rows")
	errInvalidLogin   = errors.New("Invvalid name or password")
	studentsSQLString = []string{
		`CREATE TABLE IF NOT EXISTS students(
			student_id          INT(12) UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT,             
			name                VARCHAR(50) NOT NULL,
			mobile              CHAR(11) UNIQUE NOT NULL,
			pwd                 VARCHAR(512) NOT NULL ,
			department          VARCHAR(50) DEFAULT NULL,
			maximum_borrowing   INT DEFAULT 10,
			current_borrowing   INT NO TNULL DEFAULT 0,
			overdue             BOOLEAN NOT NULL DEFAULT FALSE,
			active              BOOLEAN NOT NULL DEFAULT TRUE
			PRIMARY KEY(student_id)
		)ENGINE=InnoDB AUTO_INCREMENT=201709000001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INTO students(name,pwd,mobile,department) VALUES(?,?,?,?)`,
		`SELECT student_id,pwd FROM students WHERE name = ? LOCK IN SHARE MODE`,
		`UPDATE students SET mobile = ? WHERE student_id = ? LIMIT 1`,
		`SELECT pwd FROM students WHERE student_id = ? LOCK IN SHARE MODE`,
		`UPDATE students SET pwd = ? WHERE student_id = ? LIMIT 1`,
		`DELETE FROM students WHERE student_id = ? LIMIT 1`,
	}
)

// StudentsCreateTable create students table
func StudentsCreateTable(db *sql.DB) error {
	_, err := db.Exec(studentsSQLString[mysqlStudentsCreateTable])
	return err
}

// StudnetsAddInfo insert students info
func StudnetsAddInfo(db *sql.DB, name, pwd, mobile, department string) (int64, error) {
	password := config.Base64ncryptionGenerate(pwd)

	result, err := db.Exec(studentsSQLString[mysqlStudentsCreateTable], name, password, mobile, department)
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

// StudentsLogin return student id
func StudentsLogin(db *sql.DB, name, pwd string) (int64, error) {
	var (
		id       int64
		password string
	)

	err := db.QueryRow(studentsSQLString[mysqlStudentsInquireIDByName], name).Scan(&id, &password)
	if err != nil {
		return 0, err
	}

	if !config.Base64Compare(pwd, password) {
		return 0, errInvalidLogin
	}

	return id, nil
}

// StudentsModifyMobille modify mobile by student id
func StudentsModifyMobille(db *sql.DB, id int64, mobile string) error {
	result, err := db.Exec(studentsSQLString[mysqlStudentsModifyMobile], mobile, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil
}

// StudentsModifyPassword modify students by id
func StudentsModifyPassword(db *sql.DB, id int64, pwd, pwdNew string) error {
	var (
		password string
	)

	err := db.QueryRow(studentsSQLString[mysqlStudentInquirePwd], id).Scan(password)
	if err != nil {
		return err
	}

	if !config.Base64Compare(pwd, password) {
		return errInvalidLogin
	}

	passwordNew := config.Base64ncryptionGenerate(pwdNew)

	result, err := db.Exec(studentsSQLString[mysqlStudentsModifyPassword], passwordNew, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errInvalidModify
	}

	return nil

}

// StudnetsDelete delete the student with the specified ID
func StudnetsDelete(db *sql.DB, id string) error {
	_, err := db.Exec(studentsSQLString[mysqlStudentsDeleteByID], id)
	return err
}
